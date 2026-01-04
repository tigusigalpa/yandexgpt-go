# Contributing to YandexGPT Go SDK

Thank you for your interest in contributing to the YandexGPT Go SDK! This document provides guidelines and instructions for contributing.

## 🚀 Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/yandexgpt-go.git
   cd yandexgpt-go
   ```
3. **Create a new branch** for your feature or bugfix:
   ```bash
   git checkout -b feature/your-feature-name
   ```

## 📋 Development Setup

### Prerequisites

- Go 1.21 or higher
- Git
- A Yandex Cloud account with YandexGPT API access (for integration tests)

### Install Dependencies

```bash
go mod download
```

### Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test -run TestName ./...
```

### Code Formatting

Before submitting your changes, ensure your code is properly formatted:

```bash
# Format code
go fmt ./...

# Run linter
go vet ./...

# Install and run golangci-lint (recommended)
golangci-lint run
```

## 🔧 Making Changes

### Code Style

- Follow standard Go conventions and idioms
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions focused and concise
- Use error wrapping for better error context

### Commit Messages

Write clear and descriptive commit messages:

```
feat: add support for streaming responses
fix: correct token expiry calculation
docs: update README with new examples
test: add tests for image generation
refactor: simplify error handling
```

Prefix types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `test`: Adding or updating tests
- `refactor`: Code refactoring
- `style`: Code style changes (formatting, etc.)
- `chore`: Maintenance tasks

### Testing

- Write tests for new features
- Ensure all tests pass before submitting
- Aim for good test coverage
- Include both unit tests and integration tests where appropriate

Example test structure:

```go
func TestFeatureName(t *testing.T) {
    tests := []struct {
        name     string
        input    interface{}
        expected interface{}
        wantErr  bool
    }{
        {
            name:     "Valid input",
            input:    "test",
            expected: "result",
            wantErr:  false,
        },
        // Add more test cases
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            // Test implementation
        })
    }
}
```

## 📝 Documentation

- Update README.md if you add new features
- Add examples for new functionality
- Document all exported functions and types
- Keep documentation clear and concise

## 🐛 Reporting Bugs

When reporting bugs, please include:

1. **Description**: Clear description of the bug
2. **Steps to reproduce**: Detailed steps to reproduce the issue
3. **Expected behavior**: What you expected to happen
4. **Actual behavior**: What actually happened
5. **Environment**: Go version, OS, etc.
6. **Code sample**: Minimal code to reproduce the issue

## 💡 Suggesting Features

When suggesting features:

1. **Use case**: Describe the problem you're trying to solve
2. **Proposed solution**: How you think it should work
3. **Alternatives**: Other solutions you've considered
4. **Additional context**: Any other relevant information

## 🔄 Pull Request Process

1. **Update documentation** as needed
2. **Add tests** for new functionality
3. **Ensure all tests pass**
4. **Update CHANGELOG.md** (if applicable)
5. **Submit the pull request** with a clear description

### Pull Request Checklist

- [ ] Code follows project style guidelines
- [ ] Tests added/updated and passing
- [ ] Documentation updated
- [ ] Commit messages are clear
- [ ] No merge conflicts
- [ ] All checks pass

## 📜 Code of Conduct

### Our Standards

- Be respectful and inclusive
- Welcome newcomers
- Accept constructive criticism
- Focus on what's best for the community
- Show empathy towards others

### Unacceptable Behavior

- Harassment or discriminatory language
- Trolling or insulting comments
- Personal or political attacks
- Publishing others' private information
- Other unprofessional conduct

## 📧 Contact

If you have questions or need help:

- **Email**: sovletig@gmail.com
- **GitHub Issues**: [Create an issue](https://github.com/tigusigalpa/yandexgpt-go/issues)

## 📄 License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to YandexGPT Go SDK! 🎉
