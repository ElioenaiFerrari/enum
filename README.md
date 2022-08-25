# Enum type for you GO project

## Install
`go get github.com/ElioenaiFerrari/enum`

## Examples

```go
e := enum.New("Gandalf", "Frodo", "Sauron")

// true
e.IsValid("Gandalf")
// true
e.IsValid("Frodo")
// true
e.IsValid("Sauron")
// false
e.IsValid("Aragorm")
```
