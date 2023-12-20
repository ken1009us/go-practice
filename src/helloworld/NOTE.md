# Hello World

## Basic Syntax

Package fmt implements formatted I/O with functions analogous to C's printf and scanf.

```go
// Returns the resulting string
fmt.Sprintf("Hello, %s!", name)
```

os.Args provides access to raw command-line arguments.

```go
if len(os.Args) > 1 {
		fmt.Println(SayHello(os.Args[1]))
	}
```

## Unit Test

If we want to write a unit test for out method, we have to import "testing".

```go
import (
	"testing"
)
```

## Project Structure

Another way to deal with the Go module if we want to reuse the hello lib.

```bash
/project-root
  /src
    /helloworld
      main.go
  /pkg
    /hello
      hello.go
  go.mod

```