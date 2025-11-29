# Builder Pattern (Fluent Interface)

This example demonstrates the **Builder Pattern** with a **Fluent Interface** in Go, allowing for readable and flexible object construction.

## Overview

The Builder Pattern separates the construction of a complex object from its representation, allowing the same construction process to create different representations. The fluent interface adds method chaining for improved readability.

## Pattern Structure

### Core Components

1. **Person** - The complex object being built
2. **PersonBuilder** - Main builder that orchestrates construction
3. **PersonAddressBuilder** - Specialized builder for address-related fields
4. **PersonJobBuilder** - Specialized builder for job-related fields

## How It Works

### 1. Initialization
```go
person := NewPersonBuilder()
```
Creates a new builder with an empty `Person` object.

### 2. Method Chaining (Fluent Interface)
Each builder method returns a builder instance (either `*PersonBuilder` or a specialized builder), enabling method chaining:

```go
NewPersonBuilder().
    Called("John Doe").
    Lives().
    At("123 Main St").
    // ... more methods
    Build()
```

### 3. Sub-Builders for Logical Grouping

**Address Building:**
- `Lives()` - Switches to `PersonAddressBuilder`
- `At(address)` - Sets address
- `WithPostcode(postcode)` - Sets postcode
- `In(city)` - Sets city and returns to `PersonBuilder`

**Job Building:**
- `Works()` - Switches to `PersonJobBuilder`
- `AsA(position)` - Sets position
- `Earning(income)` - Sets income and returns to `PersonBuilder`

### 4. Final Construction
```go
Build()
```
Returns the fully constructed `Person` object.

## Example Usage

```go
person := NewPersonBuilder().
    Called("John Doe").
    Lives().
        At("123 Main St").
        WithPostcode("12345").
        In("Anytown").
    Works().
        AsA("Software Engineer").
        Earning(75000).
    Build()

fmt.Println(person)
```

**Output:**
```
&{John Doe 123 Main St 12345 Anytown Software Engineer 75000}
```

## Key Benefits

1. **Readability** - Code reads like natural language
2. **Flexibility** - Build objects with only the fields you need
3. **Immutability** - Can enforce immutability by keeping the struct fields private
4. **Separation of Concerns** - Related fields are grouped in sub-builders
5. **Type Safety** - Compile-time checking of method calls

## Design Principles

- **Method Chaining**: Each method returns a builder instance
- **Fluent API**: Methods are named to read naturally (e.g., `Lives().At().In()`)
- **Sub-Builders**: Logical grouping of related properties
- **Return Navigation**: Sub-builders return to the main builder when their section is complete

## When to Use

- Constructing objects with many optional parameters
- When you want to make object creation more readable
- When you need to enforce a specific construction sequence
- When constructing complex objects with interdependent fields

## Running the Example

```bash
go run main.go
```
