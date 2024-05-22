# goresult

Rust's Result&lt;T> in Go

## Usage

- Using Ok[T] value

```go
import (
    "fmt"

    "github.com/mnaufalhilmym/goresult"
)

data := goresult.Ok("This is a sample Ok value")

if data.IsOk() { // true
    d := data.Unwrap()
    fmt.Println(d) // print: This is a sample Ok value
}
fmt.Println(data.IsErr()) // print: false
// data.UnwrapErr() will panic
```

- Using Err[T] value

```go
import (
    "fmt"

    "github.com/mnaufalhilmym/goresult"
)

data := goresult.Err[any](errors.New("This is a sample Err value"))

if data.IsErr() { // true
    err := data.UnwrapErr()
    fmt.Println(err) // print: This is a sample Err value
}
fmt.Println(data.IsOk()) // print: false
// data.Unwrap() will panic
```
