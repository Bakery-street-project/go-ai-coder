# Contributing to go-ai-coder

Thank you for your interest in contributing to go-ai-coder!

## Getting Started

1. Fork the repository
2. Clone your fork locally: `git clone https://github.com/bakery-street-project/go-ai-coder.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Commit your changes: `git commit -m 'Add some feature'`
6. Push to the branch: `git push origin feature/your-feature-name`
7. Open a Pull Request

## Development Setup

```bash
# Format code
gofmt -w .

# Run vetting
go vet ./...
```

## Code Style

This project follows standard Go conventions.

- **Style Guide**: https://standardjs.com
- **Linting**: Run `gofmt` before committing

## Testing

Run tests with: `go test ./...`

## Pull Request Process

1. Ensure your code passes all CI/CD checks
2. Update documentation as needed
3. Add tests for new features
4. Keep PRs focused and concise

## Questions?

Feel free to open an issue for questions or discussion.
