# Basic Types

## Basic Syntax

Types related to strings:
- byte: a synonym for uint8
- rune: a synonym for int32 for characters
- string: an immutable sequence of characters
    - physically a sequence of bytes (UTF-8 encoding)
    - logically a sequence of runes (unicode)

The length of the string is the number of bytes required to represent the unicode characters, not the number of the unicode characters.

```go
s := "élite"
fmt.Printf("%8T %[1]v %d\n", s, len(s))
fmt.Printf("%8T %[1]v\n", []rune(s))

b := []byte(s)
fmt.Printf("%8T %[1]v %d\n", b, len(b))
```

```bash
string élite 6
[]int32 [233 108 105 116 101]
[]uint8 [195 169 108 105 116 101] 6
```

Fprintf: writes formatted text to the output stream you specify.

```go
fmt.Fprintf(w io.Writer, format string, a ...any)
```

Printf is equivalent to writing Fprintf() and writes formatted text to wherever the standard output stream is currently pointing.

Sprintf writes formatted text to an array of char, as opposed to a stream.

## Search and replace program

When you use < filename, the shell redirects the contents of filename as the standard input (stdin) to the program.

```go
old, new := os.Args[1], os.Args[2]
scan := bufio.NewScanner(os.Stdin)
for scan.Scan() {
	s := strings.Split(scan.Text(), old)
	fmt.Println(s)
	t := strings.Join(s, new)
	fmt.Println(t)
}
```

```bash
go run main.go Bob Matt < test.txt
```

Original text file:

```bash
Bob went to the store.
Hi Bob, how are you?
Alan went to Japan.
```

After run the program:

```bash
Matt went to the store.
Hi Matt, how are you?
Alan went to Japan.
```