# go-ai-coder

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Bakerstreet Labs](https://img.shields.io/badge/Bakerstreet-Labs-black.svg)](https://github.com/Bakery-street-project)

> CLI toolkit for AI-assisted coding — local LLM chat, GitHub repo analysis, and code manipulation.

**Watermark:** `BAKERSTREET-LABS-2025`

---

## Overview

go-ai-coder is a collection of standalone Go CLI tools that connect to local (Ollama) or cloud-based LLMs for coding assistance. Each tool is a single-file program you can run independently.

---

## Tools

| Tool | Run | What it does |
|------|-----|-------------|
| `chat` | `go run chat.go` | Interactive chat with a local Ollama model |
| `github_ai_agent` | `go run github_ai_agent.go` | GitHub-aware AI agent — list repos, search issues/PRs, clone repos, scrape URLs, analyze code |
| `bash_tool` | `go run bash_tool.go` | Execute bash commands and capture output |
| `edit_tool` | `go run edit_tool.go` | Read and write files on disk |
| `code_search_tool` | `go run code_search_tool.go` | Search codebases using ripgrep |
| `cloud-ai` | `go run cmd/cloud-ai/main.go` | Cloud AI inference via NVIDIA NIM-compatible API |
| `list_files` | `go run list_files.go` | List directory contents |
| `read` | `go run read.go` | Read file contents |
| `read_simple` | `go run read_simple.go` | Minimal file reader |

---

## Quick Start

```bash
git clone https://github.com/Bakery-street-project/go-ai-coder.git
cd go-ai-coder
go mod download
cp .env.example .env   # configure your API keys

# Start a local AI chat (requires Ollama running)
go run chat.go

# Or launch the GitHub AI agent
go run github_ai_agent.go
```

Requires [Ollama](https://ollama.com) running locally with a model pulled (default: `llama3.2:3b`).

---

## github_ai_agent Commands

Once running, the GitHub AI agent supports these commands:

| Command | Description |
|---------|-------------|
| `/repos <org>` | List repositories for a GitHub organization |
| `/search <query>` | Search GitHub repos |
| `/issues <owner/repo>` | List open issues for a repo |
| `/prs <owner/repo>` | List open pull requests |
| `/clone <owner/repo>` | Clone a repo locally |
| `/research <topic>` | AI-powered research on a topic |
| `/scrape <url>` | Scrape and summarize a web page |
| `/learn` | Browse AI-assisted Go learning resources |
| `/read <path>` | Read a file from disk |
| `/write <path>` | Write content to a file |
| `/bash <command>` | Execute a bash command |

---

## License

MIT — see [LICENSE](LICENSE)

For commercial/enterprise use, contact [kiliaanv2@gmail.com](mailto:kiliaanv2@gmail.com)

---

**© 2025-2026 Bakerstreet Labs · Bakery Street Project · ALL RIGHTS RESERVED**  
Watermark: `BAKERSTREET-LABS-2025`
