# 🤖 Go AI Coder

[![Go](https://img.shields.io/badge/Go-1.21+-00ADD8.svg)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Bakerstreet Labs](https://img.shields.io/badge/Bakerstreet-Labs-black.svg)](https://github.com/Bakery-street-project)

> **Production-grade autonomous AI coding agent built in Go.**

The code generation engine of the [Bakerstreet Labs](https://github.com/Bakery-street-project) ecosystem.

---

## ⚡ What It Does

Go AI Coder is a production-ready autonomous coding agent that receives code generation tasks via `repository_dispatch`, produces high-quality Go/Python/TypeScript code, and can submit PRs back to the triggering repository.

```
[Task Request] → go-ai-coder → [Generated Code] → [PR / Dispatch Back]
  (dispatch)      (LLM + Go)      (tested)          beeai-hive-999
                                                     Baker-Street-Lab
```

---

## 🚀 Quick Start

```bash
git clone https://github.com/Bakery-street-project/go-ai-coder.git
cd go-ai-coder

go mod download
cp .env.example .env  # Add your NVIDIA_API_KEY
go run main.go
```

---

## 🔗 Ecosystem Integration

Go AI Coder is the **code synthesis engine** of Bakerstreet Labs:

```yaml
# Receives repository_dispatch from:
# - beeai-hive-999: agent code generation tasks
# - Baker-Street-Laboratory-1: code analysis requests
# - fusion360-agent: CNC program generation

# Sends back:
# - Completed code as PR
# - repository_dispatch: code_ready event
```

---

## 🏗️ Architecture

```
go-ai-coder/
├── agent/           # Core AI coding agent
├── llm/             # LLM interface (NVIDIA NIM compatible)
├── codegen/         # Code generation & testing
├── dispatch/        # GitHub event handler
└── main.go          # Entry point
```

---

## 🛡️ Security

- Secret scanning enabled (org-level)
- NVIDIA API key via org secrets only
- No credentials in code

---

## 📄 License

MIT · [Bakerstreet Labs](https://github.com/Bakery-street-project)
