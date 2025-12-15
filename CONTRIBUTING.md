# Contributing to Yandex Disk Go SDK

Thank you for your interest in contributing to the Yandex Disk Go SDK! This document provides guidelines and instructions for contributing.

## Getting Started

1. Fork the repository
2. Clone your fork: `git clone https://github.com/YOUR_USERNAME/yandex-disk-go.git`
3. Create a new branch: `git checkout -b feature/your-feature-name`
4. Make your changes
5. Run tests: `make test`
6. Commit your changes: `git commit -am 'Add some feature'`
7. Push to the branch: `git push origin feature/your-feature-name`
8. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.21 or higher
- Make (optional, but recommended)

### Installing Dependencies

```bash
go mod download
```

Or using Make:

```bash
make install-deps
```

## Code Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Run `go vet` to check for common mistakes
- Use meaningful variable and function names
- Add comments for exported functions and types

### Formatting

```bash
make fmt
```

### Vetting

```bash
make vet
```

### Linting

```bash
make lint
```

## Testing

All new features and bug fixes should include tests.

### Running Tests

```bash
make test
```

### Coverage Report

```bash
make coverage
```

This will generate an HTML coverage report in `coverage.html`.

### Writing Tests

- Place tests in `*_test.go` files
- Use the `testing` package
- Use `github.com/stretchr/testify/assert` for assertions
- Aim for high test coverage
- Test both success and error cases

Example:

```go
func TestNewFeature(t *testing.T) {
    // Arrange
    client := NewClient("test-token")
    
    // Act
    result, err := client.NewFeature()
    
    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
}
```

## Pull Request Guidelines

1. **Keep PRs focused**: Each PR should address a single concern
2. **Write clear commit messages**: Use descriptive commit messages
3. **Update documentation**: Update README.md if you add/change features
4. **Add tests**: Ensure your changes are covered by tests
5. **Pass CI checks**: Make sure all tests pass before submitting
6. **Follow code style**: Ensure your code follows Go conventions

## Commit Message Format

Use clear and descriptive commit messages:

```
Add feature X

- Implement Y
- Update Z
- Fix issue with W
```

Or for bug fixes:

```
Fix issue with file upload

The file upload was failing when...
This commit fixes it by...
```

## Reporting Bugs

When reporting bugs, please include:

1. Go version
2. Operating system
3. Steps to reproduce
4. Expected behavior
5. Actual behavior
6. Error messages (if any)

## Feature Requests

We welcome feature requests! Please:

1. Check if the feature already exists
2. Check if it's already requested in issues
3. Provide a clear description of the feature
4. Explain why it would be useful
5. Provide examples of how it would be used

## Code of Conduct

- Be respectful and inclusive
- Welcome newcomers
- Focus on constructive feedback
- Help others learn and grow

## Questions?

If you have questions, feel free to:

- Open an issue
- Start a discussion
- Contact the maintainers

## License

By contributing, you agree that your contributions will be licensed under the MIT License.

---

Thank you for contributing to Yandex Disk Go SDK! 🎉
