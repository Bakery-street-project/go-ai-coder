# go-ai-coder

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)
[![Bakerstreet Labs](https://img.shields.io/badge/Bakerstreet-Labs-black.svg)](https://github.com/Bakery-street-project)

> Production-grade autonomous AI coding agent built in Go.

**Watermark:** `BAKERSTREET-LABS-2025`

---

## Overview

go-ai-coder is the code synthesis engine of the Bakerstreet Labs ecosystem. It receives code generation tasks via `repository_dispatch`, produces high-quality Go/Python/TypeScript code, and submits PRs back to the triggering repository.

```
[Task Request] → go-ai-coder → [Generated Code] → [PR]
  (dispatch)      (LLM + Go)      (tested)
```

---

## Quick Start

```bash
git clone https://github.com/Bakery-street-project/go-ai-coder.git
cd go-ai-coder
go mod download
cp .env.example .env  # configure your API keys
go run ./cmd/chat
```

---

## Project Structure

```
go-ai-coder/
├── cmd/               # Entry points (chat, bash_tool, edit_tool, etc.)
│   ├── chat/          # Interactive AI coding session
│   ├── bash_tool/     # Bash command execution tool
│   ├── edit_tool/     # Code editing tool
│   ├── github_ai_agent/  # GitHub-integrated AI agent
│   └── ...
├── internal/          # Internal packages
│   ├── config/        # Configuration
│   ├── license/       # License validation
│   └── security/      # Security utilities
├── scripts/           # Helper scripts
└── .github/workflows/ # CI/CD pipelines
```

---

## Tools

| Command | Description |
|---------|-------------|
| `chat` | Interactive AI coding session with LLM backend |
| `bash_tool` | Execute bash commands with AI supervision |
| `edit_tool` | AI-powered code editing and refactoring |
| `github_ai_agent` | GitHub event-driven code generation agent |
| `cloud-ai` | Cloud-based AI inference interface |
| `code_search_tool` | Semantic code search |
| `list_files` | File tree listing utility |
| `read` | File content reader |
| `read_simple` | Minimal file reader |

---

## Integration

Receives `repository_dispatch` events from Bakerstreet Labs ecosystem agents and returns generated code as PRs or dispatch responses.

---

## License

MIT — see [LICENSE](LICENSE)

For commercial/enterprise use, contact [kiliaanv2@gmail.com](mailto:kiliaanv2@gmail.com)

---

**© 2025-2026 Bakerstreet Labs · Bakery Street Project · ALL RIGHTS RESERVED**  
Watermark: `BAKERSTREET-LABS-2025`
