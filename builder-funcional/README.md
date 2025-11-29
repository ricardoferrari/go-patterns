# Builder Pattern (Functional Lazy)

This example demonstrates the **Functional Builder Pattern** with **Lazy Evaluation** in Go, using functions as first-class citizens to defer object construction until explicitly requested.

## Overview

The Functional Lazy Builder Pattern collects construction operations as functions and executes them only when the object is finally built. This approach combines functional programming concepts with the builder pattern for maximum flexibility and deferred execution.

## Pattern Structure

### Core Components

1. **Person** - The object being built
2. **PersonModifier** - A function type that modifies a Person: `func(*Person)`
3. **PersonBuilder** - Accumulates modifier functions in a slice
4. **Build()** - Executes all accumulated functions to construct the Person

## How It Works

### 1. Function Accumulation
Instead of immediately setting values, each builder method appends a **modifier function** to an `actions` slice:

```go
func (pb *PersonBuilder) Called(name string) *PersonBuilder {
    pb.actions = append(pb.actions, func(p *Person) {
        p.Name = name
    })
    return pb
}
```

### 2. Lazy Evaluation
The Person object isn't created or modified until `Build()` is called:

```go
personBuilder := NewPersonBuilder().
    Called("John Doe").
    LivesAt("123 Main St", "Anytown")
// At this point, NO Person object exists yet!
```

### 3. Deferred Execution
When `Build()` is called, it:
1. Creates an empty Person object
2. Executes each accumulated function in order
3. Returns the fully constructed Person

```go
func (pb *PersonBuilder) Build() *Person {
    p := &Person{}
    for _, action := range pb.actions {
        action(p)
    }
    return p
}
```

## Example Usage

```go
// Phase 1: Accumulate operations (no Person created yet)
personBuilder := NewPersonBuilder().
    Called("John Doe").
    LivesAt("123 Main St", "Anytown")

// Phase 2: Execute all operations and build Person
person := personBuilder.Build()

fmt.Println(*person)
```

**Output:**
```
Building person lazily because of deferred execution of actions...
Adding name...
Adding address and city...
Person built:
{John Doe 123 Main St Anytown}
```

## Key Differences from Traditional Builder

| Traditional Builder | Functional Lazy Builder |
|---------------------|-------------------------|
| Modifies object immediately | Stores functions for later execution |
| Object exists during building | Object created only at Build() |
| Direct field assignment | Functional composition |
| Eager evaluation | Lazy evaluation |

## Key Benefits

1. **Lazy Evaluation** - Object construction is deferred until needed
2. **Functional Composition** - Operations are composable functions
3. **Flexibility** - Can conditionally add operations or reuse builder configurations
4. **Immutability of Builder** - Builder itself doesn't hold state, only functions
5. **Testability** - Can inspect or manipulate the actions array
6. **Reusability** - Can create templates and apply them multiple times

## Advanced Usage Scenarios

### Conditional Building
```go
builder := NewPersonBuilder().Called("Jane")

if needAddress {
    builder = builder.LivesAt("456 Oak Ave", "Springfield")
}

person := builder.Build()
```

### Builder Templates
```go
baseBuilder := NewPersonBuilder().Called("Default Name")

person1 := baseBuilder.LivesAt("Location 1", "City 1").Build()
person2 := baseBuilder.LivesAt("Location 2", "City 2").Build()
```

### Custom Modifiers
```go
func (pb *PersonBuilder) AddCustomModifier(modifier PersonModifier) *PersonBuilder {
    pb.actions = append(pb.actions, modifier)
    return pb
}

// Usage
builder.AddCustomModifier(func(p *Person) {
    p.Name = strings.ToUpper(p.Name)
})
```

## Design Principles

- **First-Class Functions**: Functions are values that can be stored and passed around
- **Closure**: Each modifier function captures its parameters in a closure
- **Deferred Execution**: Computation is delayed until explicitly invoked
- **Functional Pipeline**: Building becomes a pipeline of function applications

## When to Use

- When object construction is expensive and should be delayed
- When you want to separate construction logic from execution
- When you need to conditionally apply modifications
- When you want to create reusable builder templates
- When you need to inspect or modify the construction pipeline before execution

## Advantages Over Traditional Builder

1. **More Flexible**: Can add, remove, or reorder operations dynamically
2. **Better Composition**: Functions compose naturally
3. **Easier Testing**: Can verify accumulated actions without building
4. **Memory Efficient**: Object allocated only when needed

## Performance Considerations

- Small overhead from function calls during Build()
- Closures capture variables, which uses memory
- Trade-off: flexibility vs. immediate execution

## Running the Example

```bash
go run main.go
```
