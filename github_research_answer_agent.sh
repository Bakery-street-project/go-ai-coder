#!/usr/bin/env bash
set -euo pipefail

ORG="Bakery-street-project"
REPO="go-ai-coder"
REPO_DIR="${HOME}/github-org/${REPO}"
QUESTIONS_FILE="research/questions.md"

QID="${1:-}"

if [ -z "${QID}" ]; then
  echo "Usage: $0 Q20260217180236" >&2
  exit 1
fi

EDITOR_BIN="${EDITOR:-nano}"

if [ ! -d "${REPO_DIR}" ]; then
  echo "[*] Cloning ${ORG}/${REPO}..."
  mkdir -p "${HOME}/github-org"
  gh repo clone "${ORG}/${REPO}" "${REPO_DIR}"
fi

cd "${REPO_DIR}"

if [ ! -f "${QUESTIONS_FILE}" ]; then
  echo "[!] ${QUESTIONS_FILE} not found." >&2
  exit 1
fi

echo "[*] Opening question ${QID} in ${QUESTIONS_FILE} with ${EDITOR_BIN}"
${EDITOR_BIN} "${QUESTIONS_FILE}"

# After manual edit (you update Status / Notes / Final Answer), commit.
if git status --porcelain | grep -q "${QUESTIONS_FILE}"; then
  git add "${QUESTIONS_FILE}"
  git commit -m "research: update ${QID}"
  git push -u origin main
fi

echo "[✓] Updated research question ${QID}."
