// BAKERSTREET-LABS-2025 — Bakerstreet Labs
// omarchy desktop TUI — AI coding agent launcher for the omarchy desktop environment
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/joho/godotenv"
)

// ── Styles ────────────────────────────────────────────────────────────────────

var (
	styleTitle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("5")).Padding(0, 1)
	styleSelected = lipgloss.NewStyle().Background(lipgloss.Color("5")).Foreground(lipgloss.Color("0")).Padding(0, 2)
	styleNormal   = lipgloss.NewStyle().Padding(0, 2)
	styleDesc     = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Padding(0, 4)
	styleStatus   = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Padding(0, 1)
	styleBox      = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("5")).
			Padding(1, 2)
	styleKey = lipgloss.NewStyle().Foreground(lipgloss.Color("11"))
)

// ── App entries ───────────────────────────────────────────────────────────────

type App struct {
	Name    string
	Desc    string
	Icon    string
	Command []string
	Env     map[string]string
}

var apps = []App{
	{
		Name: "AI Coder Agent",
		Desc: "Claude-powered coding agent with tool use (read/write/bash/search)",
		Icon: "⚡",
		Command: []string{"go-ai-coder-agent"},
	},
	{
		Name: "GitHub AI Agent",
		Desc: "GitHub integration with Ollama — repos, issues, PRs, code analysis",
		Icon: "🐙",
		Command: []string{"go-ai-coder"},
	},
	{
		Name: "Cloud AI Agent",
		Desc: "Cloud-based AI agent with enterprise features",
		Icon: "☁️",
		Command: []string{"go-ai-coder-cloud"},
	},
	{
		Name: "New Project",
		Desc: "Scaffold a new Go project with AI assistance",
		Icon: "🚀",
		Command: []string{"bash", "-c", "go-ai-coder-agent"},
	},
}

// ── Model ─────────────────────────────────────────────────────────────────────

type model struct {
	cursor  int
	width   int
	height  int
	status  string
	apiKey  bool
}

func initialModel() model {
	godotenv.Load()
	return model{
		apiKey: os.Getenv("ANTHROPIC_API_KEY") != "",
	}
}

func (m model) Init() tea.Cmd { return nil }

// ── Update ────────────────────────────────────────────────────────────────────

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(apps)-1 {
				m.cursor++
			}
		case "enter", " ":
			return m, m.launchApp(apps[m.cursor])
		}
	}
	return m, nil
}

func (m model) launchApp(app App) tea.Cmd {
	return tea.ExecProcess(exec.Command(app.Command[0], app.Command[1:]...), func(err error) tea.Msg {
		if err != nil {
			return statusMsg("failed to launch " + app.Name + ": " + err.Error())
		}
		return statusMsg("")
	})
}

type statusMsg string

// ── View ──────────────────────────────────────────────────────────────────────

func (m model) View() string {
	if m.width == 0 {
		return "Loading..."
	}

	var sb strings.Builder

	// Title
	title := styleTitle.Render("omarchy desktop") + "  " +
		lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render("go-ai-coder apps")
	sb.WriteString(title + "\n\n")

	// API key warning
	if !m.apiKey {
		sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Render(
			"  ⚠  ANTHROPIC_API_KEY not set — AI Coder Agent requires it",
		) + "\n\n")
	}

	// App list
	var rows []string
	for i, app := range apps {
		icon := app.Icon + " "
		if i == m.cursor {
			rows = append(rows, styleSelected.Render(icon+app.Name))
			rows = append(rows, styleDesc.Render(app.Desc))
		} else {
			rows = append(rows, styleNormal.Render(icon+app.Name))
		}
	}

	list := styleBox.Width(m.width - 4).Render(strings.Join(rows, "\n"))
	sb.WriteString(list + "\n\n")

	// Status
	if m.status != "" {
		sb.WriteString(styleStatus.Render(m.status) + "\n")
	}

	// Keys
	keys := styleKey.Render("↑↓") + " navigate  " +
		styleKey.Render("Enter") + " launch  " +
		styleKey.Render("q") + " quit"
	sb.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render(keys))

	return sb.String()
}

// ── Main ──────────────────────────────────────────────────────────────────────

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
