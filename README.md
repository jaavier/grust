# grust

`grust` is a Go library inspired by Rust's Result type, providing ergonomic error handling and functional programming constructs.

## Features

### Result

The `Result` type represents the result of an operation that may fail. It can hold either a successful value (`Ok`) or an error (`Err`).

#### Example Usage

```go
package main

import (
    "fmt"
    "github.com/jaavier/grust"
)

func divide(a, b int) grust.Result {
    if b == 0 {
        return grust.Result{
            Err: func() interface{} { return fmt.Errorf("division by zero") },
            IsErr: true,
        }
    }
    return grust.Result{
        Ok: func() interface{} { return a / b },
        IsOk: true,
    }
}

func main() {
    result := divide(10, 2)
    if result.IsErr {
        fmt.Println("Error:", result.Err())
    } else {
        fmt.Println("Result:", result.Ok())
    }
}
```

### Map

The `Map` function applies a function to the value contained in a `Result`, returning a new `Result` with the transformed value.

#### Example Usage

```go
result := divide(10, 2).Map(func(val interface{}) interface{} {
    return val.(int) * 2
})
fmt.Println("Mapped Result:", result.Ok()) // Output: Mapped Result: 10
```

### AndThen

The `AndThen` function chains a function that returns a `Result`, allowing sequential operations with error handling.

#### Example Usage

```go
result := divide(10, 2).AndThen(func(val interface{}) *grust.Result {
    return divide(val.(int), 2)
})
fmt.Println("Chained Result:", result.Ok()) // Output: Chained Result: 5
```

### OrElse

The `OrElse` function returns the `Ok` value if the `Result` is successful, otherwise it returns a default value.

#### Example Usage

```go
result := divide(10, 2).OrElse(0)
fmt.Println("Result:", result) // Output: Result: 5
```

### Unwrap

The `Unwrap` function extracts the value from a successful `Result` or panics with the contained error if it's an error.

#### Example Usage

```go
result := divide(10, 2)
fmt.Println("Unwrapped Result:", result.Unwrap()) // Output: Unwrapped Result: 5
```

### UnwrapOr

The `UnwrapOr` function extracts the value from a successful `Result` or returns a default value if it's an error.

#### Example Usage

```go
result := divide(10, 0)
fmt.Println("Unwrapped Result:", result.UnwrapOr(0)) // Output: Unwrapped Result: 0
```

### UnwrapErr

The `UnwrapErr` function extracts the error from a failed `Result` or panics with the error if it's successful.

#### Example Usage

```go
result := divide(10, 0)
fmt.Println("Unwrapped Error:", result.UnwrapErr()) // Output: Unwrapped Error: division by zero
```

## Installation

To use `grust` in your Go project, simply import it:

```go
import "github.com/jaavier/grust"
```

## License

This library is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---