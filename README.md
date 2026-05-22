# go-ai-coder

[![Go](https://img.shields.io/badge/Go-1.24+-00ADD8.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Bakerstreet Labs](https://img.shields.io/badge/Bakerstreet-Lab-black.svg)](https://github.com/Bakery-street-project)

> Claude-powered autonomous coding agent with a TUI — read/write files, run bash, search code, all from your terminal.

**Watermark:** `BAKERSTREET-LABS-2025`

---

## Overview

go-ai-coder is a collection of Go CLI tools for AI-assisted coding. The flagship is `cmd/agent` — a Claude Code-style TUI agent that uses Anthropic's tool-use API to autonomously read/write files, run commands, and search codebases. It also ships a GitHub-aware Ollama agent and an omarchy desktop launcher.

---

## Agent TUI (Claude Code-style)

```
┌─────────────────────────────────────────────────────┐
│ ⚡ go-ai-coder  claude-sonnet-4 · Ctrl+C to quit    │
├─────────────────────────────────────────────────────┤
│ ◆ Claude                                            │
│   I'll read the file and fix the bug.               │
│                                                     │
│ ⚙ read_file                                         │
│   package main...                                   │
│                                                     │
│ ◆ Claude                                            │
│   Fixed. The issue was a nil pointer dereference.   │
├─────────────────────────────────────────────────────┤
│ ▶ fix the bug in main.go█                           │
└─────────────────────────────────────────────────────┘
```

### Tools available to the agent

| Tool | Description |
|------|-------------|
| `read_file` | Read any file from disk |
| `write_file` | Write/create files |
| `list_files` | List directory contents |
| `bash` | Execute shell commands |
| `search_code` | Grep across Go source files |

---

## Quick Start

```bash
git clone https://github.com/Bakery-street-project/go-ai-coder.git
cd go-ai-coder
go mod download
cp .env.example .env
# Add ANTHROPIC_API_KEY to .env

# Launch the Claude agent TUI
go run ./cmd/agent/

# Or the GitHub/Ollama agent (requires Ollama running)
go run github_ai_agent.go
```

### Build binaries

```bash
go build -o go-ai-coder-agent ./cmd/agent/
go build -o go-ai-coder-omarchy ./cmd/omarchy/
go build -o go-ai-coder ./
```

---

## GitHub Codespaces

Open instantly in a pre-configured dev environment:

[![Open in GitHub Codespaces](https://github.com/codespaces/badge.svg)](https://codespaces.new/Bakery-street-project/go-ai-coder)

```bash
# Inside the codespace
go run ./cmd/agent/
```

---

## omarchy Desktop App

go-ai-coder ships as an **omarchy desktop app** — a TUI launcher that integrates with the [omarchy](https://github.com/basecamp/omarchy) desktop environment.

### Launch

```bash
go run ./cmd/omarchy/
# or
go-ai-coder-omarchy
```

### omarchy Desktop App Entry

```toml
# ~/.config/omarchy/apps/go-ai-coder.toml
name = "go-ai-coder"
description = "Claude-powered AI coding agent"
icon = "⚡"
command = "go-ai-coder-omarchy"
category = "development"
```

### App listing

| App | Command | Description |
|-----|---------|-------------|
| ⚡ AI Coder Agent | `go-ai-coder-agent` | Claude-powered coding agent with tool use |
| 🐙 GitHub AI Agent | `go-ai-coder` | GitHub integration with Ollama |
| ☁️ Cloud AI Agent | `go-ai-coder-cloud` | Cloud-based AI agent |

---

## All Tools

| Tool | Run | What it does |
|------|-----|-------------|
| `cmd/agent` | `go run ./cmd/agent/` | **Claude TUI agent** — tool use, file ops, bash |
| `cmd/omarchy` | `go run ./cmd/omarchy/` | omarchy desktop launcher |
| `github_ai_agent` | `go run github_ai_agent.go` | GitHub-aware Ollama agent |
| `bash_tool` | `go run bash_tool.go` | Execute bash commands |
| `edit_tool` | `go run edit_tool.go` | Read/write files |
| `code_search_tool` | `go run code_search_tool.go` | Search codebases with ripgrep |
| `cmd/cloud-ai` | `go run cmd/cloud-ai/main.go` | Cloud AI via NVIDIA NIM API |

---

## Configuration

Copy `.env.example` to `.env` and set:

```env
ANTHROPIC_API_KEY=sk-ant-...   # for cmd/agent (Claude)
GITHUB_TOKEN=ghp_...           # for GitHub features
# AIMLAPI_API_KEY=...          # optional cloud AI
```

---

## Architecture

```
go-ai-coder/
├── cmd/
│   ├── agent/main.go      # Claude TUI agent (bubbletea + anthropic-sdk-go)
│   ├── omarchy/main.go    # omarchy desktop launcher
│   └── cloud-ai/main.go   # cloud AI agent
├── github_ai_agent.go     # GitHub + Ollama agent (main package)
├── chat.go                # basic Ollama chat
├── bash_tool.go           # bash execution tool
├── edit_tool.go           # file edit tool
├── code_search_tool.go    # code search tool
└── .devcontainer/         # GitHub Codespaces config
```

---

## License

MIT — see [LICENSE](LICENSE)

For commercial/enterprise use, contact [kiliaanv2@gmail.com](mailto:kiliaanv2@gmail.com)

---

**© 2025-2026 Bakerstreet Labs · Bakery Street Project · ALL RIGHTS RESERVED**  
Watermark: `BAKERSTREET-LABS-2025`
