# Security Policy

## 🛡️ Supported Versions

| Version | Supported |
|---------|-----------|
| `main` | ✅ Active |
| Tagged releases | ✅ Latest tag |
| Older   | ❌ Not supported |

## 🔒 Reporting a Vulnerability

**Please do NOT report security vulnerabilities via public GitHub Issues.**

Report security vulnerabilities responsibly:

1. **Email:** [security@bakerstreetproject221b.store](mailto:security@bakerstreetproject221b.store)
2. **GitHub Private Advisory:** Use [GitHub's private security advisory feature](https://github.com/Bakery-street-project/go-ai-coder/security/advisories/new)

### What to Include

- Description of the vulnerability
- Steps to reproduce
- Potential impact
- Suggested fixes (optional)

### Response Timeline

- **Acknowledgement:** Within 48 hours
- **Initial Assessment:** Within 5 business days
- **Fix & Coordinated Disclosure:** Coordinated with reporter

## 🚨 Security Architecture

### Local Processing First
- All AI inference runs locally via Ollama — **your code never leaves your machine**
- No telemetry or usage analytics
- Conversation history stored locally only

### Credential Security
- GitHub tokens are read from environment variables — **never hardcoded**
- Tokens are masked in all log output
- Session data is encrypted at rest

### Input Security
- All user-provided URLs validated before scraping
- File paths sanitized to prevent path traversal
- Content length limits enforced

## 🔐 Security Scanning

This repository runs automated security scans on every push:

- **gosec** — Go source code SAST
- **govulncheck** — Known CVE detection in dependencies
- **CodeQL** — Semantic code analysis
- **Trivy** — Filesystem vulnerability scanning
- **SBOM generation** — Software Bill of Materials on every release
- **Dependabot** — Automated dependency updates weekly

## 🔐 Best Practices for Contributors

```bash
# Run security checks locally before pushing
go install github.com/securego/gosec/v2/cmd/gosec@latest
go install golang.org/x/vuln/cmd/govulncheck@latest

gosec ./...
govulncheck ./...
go vet ./...
```

Never commit:
- API keys or tokens
- Passwords or credentials
- Private certificates
- `.env` files with real values

---

*Maintained by [Bakertreet Labs](https://github.com/Bakery-street-project)*
