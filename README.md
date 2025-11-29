# Go Design Patterns

A curated collection of useful design patterns implemented in Go. This repository serves as a boilerplate and reference for common software design patterns, demonstrating idiomatic Go implementations.

## 📚 Available Patterns

### Builder Pattern
Demonstrates different approaches to the Builder pattern in Go:

- **[builder/](builder/)** - Traditional Builder Pattern with Fluent Interface
  - Implements sub-builders for logical grouping
  - Method chaining for readable object construction
  - Specialized builders for address and job information

- **[builder-funcional/](builder-funcional/)** - Functional Builder Pattern with Lazy Evaluation
  - Uses functions as first-class citizens
  - Deferred execution with modifier functions
  - Lazy evaluation for flexible object construction

## 🚀 Getting Started

Each pattern is contained in its own directory with:
- Complete working implementation (`main.go`)
- Detailed documentation (`README.md`)
- Runnable examples

### Running Examples

Navigate to any pattern directory and run:

```bash
cd <pattern-directory>
go run main.go
```

For example:
```bash
cd builder
go run main.go
```

## 📖 Learning Path

1. Start with the traditional **builder** pattern to understand the basics of fluent interfaces
2. Explore **builder-funcional** to see how functional programming concepts can enhance flexibility
3. Compare both approaches to understand trade-offs between eager and lazy evaluation

## 🎯 Use Cases

These patterns are particularly useful for:
- Building complex objects with many optional parameters
- Creating readable and maintainable configuration code
- Implementing flexible APIs with method chaining
- Separating object construction from representation

## 🤝 Contributing

Feel free to add new patterns or improve existing implementations. Each pattern should include:
- Working code example
- Comprehensive README explaining the pattern
- Clear use cases and benefits
- Running examples

## 📝 Pattern Documentation

Each pattern directory contains its own detailed README covering:
- Pattern overview and structure
- How it works with code examples
- Key benefits and design principles
- When to use the pattern
- Comparison with alternative approaches

## 🛠️ Requirements

- Go 1.16 or higher

## 📄 License

This repository is intended for educational purposes and as a reference implementation of design patterns in Go.
