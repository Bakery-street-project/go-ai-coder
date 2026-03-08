#!/bin/bash
set -e

# 1) Install Ollama CLI (idempotent)
curl -fsSL https://ollama.com/install.sh | sh

# 2) Verify Ollama
ollama -v
ollama list || true

# 3) Ensure server is running for TUIs
systemctl --user enable --now ollama || ollama serve &

# 4) Ask which TUI to install
echo
read -p "Install LazyLlama or parllama? (L/P): " choice

case "$choice" in
  L|l)
    # LazyLlama: lightweight Ollama TUI
    curl -fsSL https://lazyllama.app/install.sh | sh
    echo "Starting LazyLlama..."
    lazyllama
    ;;
  P|p)
    # parllama: multi-provider TUI (Ollama, OpenAI, etc.)
    pip install --user parllama
    export OLLAMA_HOST=127.0.0.1:11434
    echo "Starting parllama..."
    ~/.local/bin/parllama || parllama
    ;;
  *)
    echo "Invalid choice. Exiting."
    exit 1
    ;;
esac
