# ☁️ Go AI Coder — Enterprise AI Coding Assistant

> **Local-first, privacy-preserving AI coding assistant with GitHub integration, web scraping & multi-cloud deployment**
> Built with Go · Powered by Ollama · Part of [Bakertreet Labs](https://github.com/Bakery-street-project)

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-22c55e?style=for-the-badge)](LICENSE)
[![CI](https://github.com/Bakery-street-project/go-ai-coder/actions/workflows/ci.yml/badge.svg)](https://github.com/Bakery-street-project/go-ai-coder/actions/workflows/ci.yml)
[![Security](https://github.com/Bakery-street-project/go-ai-coder/actions/workflows/security.yml/badge.svg)](https://github.com/Bakery-street-project/go-ai-coder/actions/workflows/security.yml)
[![Homepage](https://img.shields.io/badge/Website-bakerstreetproject221b.store-6366f1?style=for-the-badge)](https://bakerstreetproject221b.store)

---

## 🎯 What It Does

**Go AI Coder** is an enterprise-grade AI coding assistant that combines **local AI inference** (via Ollama) with **GitHub repository analysis** and **autonomous web research** — all from your terminal. Your code never leaves your machine.

```
┌──────────────────────────────────────────────────────────┐
│                    go-ai-coder CLI                       │
│                                                          │
│  ┌────────────────┐  ┌──────────────┐  ┌─────────────┐  │
│  │  Local AI      │  │   GitHub     │  │  Web        │  │
│  │  (Ollama)      │  │  Integration │  │  Research   │  │
│  │                │  │              │  │             │  │
│  │  llama3.2      │  │  Repo        │  │  Scrape     │  │
│  │  codellama     │  │  Analysis    │  │  Learn      │  │
│  │  qwen2.5       │  │  Issues/PRs  │  │  Curate     │  │
│  └───────┬────────┘  └──────┬───────┘  └──────┬──────┘  │
│          └──────────────────┴──────────────────┘         │
│                             │                            │
│                  ┌──────────▼──────────┐                 │
│                  │  AI Synthesis Layer │                 │
│                  │  (Context-aware     │                 │
│                  │   conversations)    │                 │
│                  └─────────────────────┘                 │
└──────────────────────────────────────────────────────────┘
```

---

## ✨ Features

### 🤖 AI-Powered Intelligence
- **Hybrid AI Architecture** — Local Ollama + Cloud AI with smart fallback
- **Go-Specific Models** — Fine-tuned models optimized for Go development
- **Context-Aware Conversations** — Remembers conversation history
- **Learning Mode** — Continuously learns from the Go ecosystem

### 🔗 GitHub Integration
- **Repository Analysis** — Analyze any GitHub repo for code quality, issues, architecture
- **Smart Search** — Find relevant repositories and code patterns
- **Issue & PR Analysis** — Get insights on project health
- **Rate Limit Handling** — Graceful fallbacks for API limits

### 🌐 Web Scraping & Research
- **Autonomous Learning** — Scrapes and learns from Go resources
- **Topic Research** — Deep dive into specific technologies
- **Curated Resources** — Access to best Go learning materials

### ☁️ Enterprise Features
- **Multi-Cloud Deployment** — AWS, GCP, Azure support
- **Custom Model Training** — Train Go-specific AI models
- **Secure Configuration** — Encrypted settings and session management
- **Conversation Persistence** — Auto-save and history management
- **Plugin-Ready Architecture** — Extensible design

---

## 🚀 Quick Start

### Prerequisites

1. **Go 1.21+** — [Install Go](https://golang.org/doc/install)
2. **Ollama** — [Install Ollama](https://ollama.ai/download)
3. **GitHub Token** (optional) — [Create Personal Access Token](https://github.com/settings/tokens)

### Installation

```bash
# Clone the repository
git clone https://github.com/Bakery-street-project/go-ai-coder.git
cd go-ai-coder

# Install dependencies
go mod tidy

# Build the application
go build -o go-ai-coder cmd/main.go

# Install to PATH (optional)
sudo cp go-ai-coder /usr/local/bin/
```

### Pull a model with Ollama first

```bash
ollama pull llama3.2:3b
# or for code-focused work:
ollama pull qwen2.5-coder:7b
```

### Basic Usage

```bash
# Start with default settings
go-ai-coder

# Use a custom model
go-ai-coder --model codellama:13b --tokens 4000

# Use cloud AI with fallback
go-ai-coder --cloud --cloud-url "https://your-service.com" --fallback

# Enable verbose mode
go-ai-coder --verbose

# Show help
go-ai-coder --help
```

---

## 🔧 Configuration

### Environment Variables

```bash
export AI_MODEL="llama3.2:3b"
export AI_MAX_TOKENS="2000"
export AI_TEMPERATURE="0.7"
export OLLAMA_URL="http://localhost:11434/v1"
export GITHUB_TOKEN="your_token_here"         # Never hardcode — use env vars!
export LEARNING_DIR="ai_learning"
export CACHE_ENABLED="true"
export AUTO_SAVE="true"
export VERBOSE="false"
```

### Command Line Flags

```bash
go-ai-coder \
  --model llama3.2:3b \
  --tokens 2000 \
  --temp 0.7 \
  --ollama http://localhost:11434/v1 \
  --learning ai_learning \
  --cache \
  --autosave \
  --verbose
```

---

## 💬 Commands

### Core
- `help` — Show help message
- `config` — Display current configuration
- `quit` / `exit` — Exit the application

### File Operations
- `read <file_or_folder>` — Read and analyze file content with AI
- `list <directory>` — List directory contents with AI analysis

### GitHub Integration
- `github repos` — List and analyze your repositories
- `github search <query>` — Search GitHub repositories
- `github issues <repo>` — Analyze repository issues
- `github prs <repo>` — Analyze pull requests
- `github clone <repo>` — Clone a repository

### AI Learning
- `ai learn` — Comprehensive Go ecosystem research
- `ai research <topic>` — Research specific topics
- `ai scrape <url>` — Learn from web content
- `go resources` — Show curated Go learning resources

### Examples

```bash
# Analyze your code
You: read main.go

# Search for Go web frameworks
You: github search golang web framework

# Research machine learning in Go
You: ai research machine learning

# Learn from official Go documentation
You: ai scrape https://golang.org/doc/tutorial/

# Get curated Go resources
You: go resources
```

---

## ☁️ Cloud Deployment

```bash
# Set up cloud environment
export CLOUD_PROVIDER="aws"  # or gcp/azure
export GITHUB_TOKEN="your_token"
./setup-cloud-environment.sh

# Deploy cloud AI service
./deploy-cloud-ai.sh

# Train custom Go model
source go-ai-env/bin/activate
python go-ai-model-trainer.py
```

---

## 📁 Project Structure

```
go-ai-coder/
├── cmd/                         # Main application entry
├── internal/                    # Internal packages
│   ├── config/                  # Configuration management
│   └── security/                # Security utilities
├── configs/                     # Configuration files
├── docs/                        # Documentation
├── scripts/                     # Build and utility scripts
├── assets/                      # Static assets
├── examples/                    # Example configurations
└── .github/
    ├── workflows/
    │   ├── ci.yml               # Build · lint · test
    │   ├── security.yml         # gosec · govulncheck · CodeQL
    │   └── release.yml          # Automated releases via goreleaser
    └── dependabot.yml
```

---

## 🔒 Security

### Privacy First
- **Local Processing** — All AI inference happens locally via Ollama
- **No Data Collection** — No personal data sent to external services
- **Secure Storage** — Encrypted configuration and session data
- **Input Sanitization** — All user input is sanitized and validated

### Security Features
- **Token Masking** — Sensitive tokens are never logged
- **URL Validation** — Web scraping URLs validated for safety
- **Content Limits** — File size and content length limits
- **Secure Filenames** — Generated filenames prevent path traversal

### Best Practices
1. **Use environment variables** for all sensitive configuration
2. **Regular token rotation** — Rotate GitHub tokens regularly
3. **Local network** — Run Ollama on localhost only
4. **Keep updated** — Pull latest for security patches

See [SECURITY.md](SECURITY.md) for our full security policy.

---

## 💰 Monetization & Licensing

Go AI Coder uses an **Open-Core model**:

| Tier | Features |
|------|---------|
| **Free / Open Source** | Core CLI, local Ollama inference, GitHub integration, web research |
| **Pro** | Advanced models, cloud sync, priority support, custom integrations |
| **Enterprise** | RBAC, SSO, audit logs, dedicated support, custom model training |

---

## 🤝 Contributing

We welcome contributions! See our [Contributing Guide](CONTRIBUTING.md) for details.

```bash
# Fork and clone
git clone https://github.com/your-username/go-ai-coder.git
cd go-ai-coder

# Install development dependencies
go mod download

# Run tests
go test ./...

# Build for development
go build -o go-ai-coder cmd/main.go
```

---

## 📄 License

MIT — see [LICENSE](LICENSE) for details.

---

<div align="center">

**[Bakertreet Labs](https://github.com/Bakery-street-project)** · [Website](https://bakerstreetproject221b.store) · Building the future, one agent at a time 🧪

*Local AI · Enterprise Grade · Open Source*

</div>
