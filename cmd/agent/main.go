// BAKERSTREET-LABS-2025 — Bakerstreet Labs
// Claude-like agent TUI with tool use, streaming, and file operations
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

// ── Styles ────────────────────────────────────────────────────────────────────

var (
	styleUser      = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	styleAssistant = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	styleTool      = lipgloss.NewStyle().Foreground(lipgloss.Color("11")).Bold(true)
	styleError     = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
	styleDim       = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	styleBorder    = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("5")).
			Padding(0, 1)
	styleInput = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, false, true).
			BorderForeground(lipgloss.Color("5")).
			Padding(0, 1)
)

// ── Tool definitions ──────────────────────────────────────────────────────────

var tools = []anthropic.ToolUnionParam{
	{OfTool: &anthropic.ToolParam{
		Name:        "read_file",
		Description: anthropic.String("Read the contents of a file"),
		InputSchema: anthropic.ToolInputSchemaParam{
			Properties: map[string]interface{}{
				"path": map[string]string{"type": "string", "description": "File path to read"},
			},
			Required: []string{"path"},
		},
	}},
	{OfTool: &anthropic.ToolParam{
		Name:        "write_file",
		Description: anthropic.String("Write content to a file"),
		InputSchema: anthropic.ToolInputSchemaParam{
			Properties: map[string]interface{}{
				"path":    map[string]string{"type": "string", "description": "File path to write"},
				"content": map[string]string{"type": "string", "description": "Content to write"},
			},
			Required: []string{"path", "content"},
		},
	}},
	{OfTool: &anthropic.ToolParam{
		Name:        "list_files",
		Description: anthropic.String("List files in a directory"),
		InputSchema: anthropic.ToolInputSchemaParam{
			Properties: map[string]interface{}{
				"path": map[string]string{"type": "string", "description": "Directory path (default: .)"},
			},
		},
	}},
	{OfTool: &anthropic.ToolParam{
		Name:        "bash",
		Description: anthropic.String("Execute a bash command"),
		InputSchema: anthropic.ToolInputSchemaParam{
			Properties: map[string]interface{}{
				"command": map[string]string{"type": "string", "description": "Bash command to execute"},
			},
			Required: []string{"command"},
		},
	}},
	{OfTool: &anthropic.ToolParam{
		Name:        "search_code",
		Description: anthropic.String("Search for a pattern in code files using grep"),
		InputSchema: anthropic.ToolInputSchemaParam{
			Properties: map[string]interface{}{
				"pattern": map[string]string{"type": "string", "description": "Search pattern"},
				"path":    map[string]string{"type": "string", "description": "Directory to search (default: .)"},
			},
			Required: []string{"pattern"},
		},
	}},
}

// ── Tool execution ────────────────────────────────────────────────────────────

func executeTool(name string, input map[string]interface{}) string {
	switch name {
	case "read_file":
		path, _ := input["path"].(string)
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		return string(data)

	case "write_file":
		path, _ := input["path"].(string)
		content, _ := input["content"].(string)
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return fmt.Sprintf("error creating dirs: %v", err)
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		return fmt.Sprintf("wrote %d bytes to %s", len(content), path)

	case "list_files":
		path, _ := input["path"].(string)
		if path == "" {
			path = "."
		}
		entries, err := os.ReadDir(path)
		if err != nil {
			return fmt.Sprintf("error: %v", err)
		}
		var lines []string
		for _, e := range entries {
			if e.IsDir() {
				lines = append(lines, e.Name()+"/")
			} else {
				lines = append(lines, e.Name())
			}
		}
		return strings.Join(lines, "\n")

	case "bash":
		command, _ := input["command"].(string)
		out, err := exec.Command("bash", "-c", command).CombinedOutput()
		if err != nil {
			return fmt.Sprintf("exit error: %v\n%s", err, out)
		}
		return string(out)

	case "search_code":
		pattern, _ := input["pattern"].(string)
		path, _ := input["path"].(string)
		if path == "" {
			path = "."
		}
		out, err := exec.Command("grep", "-r", "--include=*.go", "-n", pattern, path).CombinedOutput()
		if err != nil && len(out) == 0 {
			return "no matches found"
		}
		return string(out)
	}
	return "unknown tool"
}

// ── Message types ─────────────────────────────────────────────────────────────

type ChatEntry struct {
	Role    string // "user" | "assistant" | "tool"
	Content string
	Tool    string // tool name if role == "tool"
}

// ── Bubbletea model ───────────────────────────────────────────────────────────

type model struct {
	client   *anthropic.Client
	history  []anthropic.MessageParam
	chat     []ChatEntry
	input    string
	cursor   int
	width    int
	height   int
	thinking bool
	err      string
}

type responseMsg struct {
	entries []ChatEntry
	history []anthropic.MessageParam
	err     error
}

func initialModel(client *anthropic.Client) model {
	return model{client: client}
}

func (m model) Init() tea.Cmd { return nil }

// ── Update ────────────────────────────────────────────────────────────────────

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.thinking || strings.TrimSpace(m.input) == "" {
				return m, nil
			}
			userText := m.input
			m.input = ""
			m.cursor = 0
			m.thinking = true
			m.err = ""
			m.chat = append(m.chat, ChatEntry{Role: "user", Content: userText})
			m.history = append(m.history, anthropic.NewUserMessage(anthropic.NewTextBlock(userText)))
			return m, m.runAgent()
		case tea.KeyBackspace:
			if m.cursor > 0 {
				runes := []rune(m.input)
				m.input = string(runes[:m.cursor-1]) + string(runes[m.cursor:])
				m.cursor--
			}
		case tea.KeyLeft:
			if m.cursor > 0 {
				m.cursor--
			}
		case tea.KeyRight:
			if m.cursor < len([]rune(m.input)) {
				m.cursor++
			}
		case tea.KeyRunes:
			runes := []rune(m.input)
			m.input = string(runes[:m.cursor]) + string(msg.Runes) + string(runes[m.cursor:])
			m.cursor += len(msg.Runes)
		}

	case responseMsg:
		m.thinking = false
		if msg.err != nil {
			m.err = msg.err.Error()
		} else {
			m.chat = append(m.chat, msg.entries...)
			m.history = msg.history
		}
	}
	return m, nil
}

// ── Agent loop ────────────────────────────────────────────────────────────────

func (m model) runAgent() tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		history := make([]anthropic.MessageParam, len(m.history))
		copy(history, m.history)

		var newEntries []ChatEntry

		for {
			resp, err := m.client.Messages.New(ctx, anthropic.MessageNewParams{
				Model:     anthropic.ModelClaudeSonnet4_0,
				MaxTokens: 4096,
				System: []anthropic.TextBlockParam{
					{Text: "You are an expert coding agent. You have tools to read/write files, run bash commands, and search code. Be concise and precise. When asked to implement something, do it directly using tools."},
				},
				Tools:    tools,
				Messages: history,
			})
			if err != nil {
				return responseMsg{err: err}
			}

			// Collect assistant content blocks
			var assistantText strings.Builder
			var toolUses []anthropic.ToolUseBlock
			var assistantBlocks []anthropic.ContentBlockParamUnion

			for _, block := range resp.Content {
				switch v := block.AsAny().(type) {
				case anthropic.TextBlock:
					assistantText.WriteString(v.Text)
					assistantBlocks = append(assistantBlocks, anthropic.NewTextBlock(v.Text))
				case anthropic.ToolUseBlock:
					toolUses = append(toolUses, v)
					assistantBlocks = append(assistantBlocks, anthropic.NewToolUseBlock(v.ID, v.Input, v.Name))
				}
			}

			if assistantText.Len() > 0 {
				newEntries = append(newEntries, ChatEntry{Role: "assistant", Content: assistantText.String()})
			}

			history = append(history, anthropic.NewAssistantMessage(assistantBlocks...))

			if resp.StopReason == "end_turn" || len(toolUses) == 0 {
				break
			}

			// Execute tools
			var toolResults []anthropic.ContentBlockParamUnion
			for _, tu := range toolUses {
				var input map[string]interface{}
				json.Unmarshal(tu.Input, &input)

				result := executeTool(tu.Name, input)
				newEntries = append(newEntries, ChatEntry{
					Role:    "tool",
					Tool:    tu.Name,
					Content: result,
				})
				toolResults = append(toolResults, anthropic.NewToolResultBlock(tu.ID, result, false))
			}
			history = append(history, anthropic.NewUserMessage(toolResults...))
		}

		return responseMsg{entries: newEntries, history: history}
	}
}

// ── View ──────────────────────────────────────────────────────────────────────

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sb strings.Builder

	header := styleBorder.Width(m.width - 4).Render(
		styleAssistant.Render("⚡ go-ai-coder") + "  " + styleDim.Render("claude-sonnet-4 · Ctrl+C to quit"),
	)
	sb.WriteString(header + "\n")

	chatHeight := m.height - 8
	var chatLines []string
	for _, entry := range m.chat {
		switch entry.Role {
		case "user":
			chatLines = append(chatLines, styleUser.Render("▶ You")+"\n"+entry.Content)
		case "assistant":
			chatLines = append(chatLines, styleAssistant.Render("◆ Claude")+"\n"+entry.Content)
		case "tool":
			preview := entry.Content
			if len(preview) > 200 {
				preview = preview[:200] + "…"
			}
			chatLines = append(chatLines, styleTool.Render("⚙ "+entry.Tool)+"\n"+styleDim.Render(preview))
		}
		chatLines = append(chatLines, "")
	}

	if m.thinking {
		chatLines = append(chatLines, styleDim.Render("◆ Claude is thinking…"))
	}
	if m.err != "" {
		chatLines = append(chatLines, styleError.Render("✗ "+m.err))
	}

	allLines := strings.Join(chatLines, "\n")
	lines := strings.Split(allLines, "\n")
	if len(lines) > chatHeight {
		lines = lines[len(lines)-chatHeight:]
	}
	chatArea := lipgloss.NewStyle().Height(chatHeight).Width(m.width - 2).Render(strings.Join(lines, "\n"))
	sb.WriteString(chatArea + "\n")

	// Input with cursor
	runes := []rune(m.input)
	var inputDisplay string
	if m.cursor >= len(runes) {
		inputDisplay = m.input + lipgloss.NewStyle().Reverse(true).Render(" ")
	} else {
		cur := string(runes[m.cursor])
		inputDisplay = string(runes[:m.cursor]) + lipgloss.NewStyle().Reverse(true).Render(cur) + string(runes[m.cursor+1:])
	}
	inputBox := styleInput.Width(m.width - 4).Render(styleUser.Render("▶ ") + inputDisplay)
	sb.WriteString(inputBox)

	return sb.String()
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	godotenv.Load()

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "ANTHROPIC_API_KEY not set")
		os.Exit(1)
	}

	client := anthropic.NewClient(option.WithAPIKey(apiKey))

	p := tea.NewProgram(
		initialModel(&client),
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
