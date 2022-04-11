# Go Cheat Sheet

## Build-in Types

- `bool`, `string`
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- `rune`, `byte`
- `float32`, `float64`
- `complex64`, `complex128`

## Variables

```go
var name = "Fabian"
var level, age = 9, 35
var isSkilled bool
language := "go"
```

## Constants

```go
const name string = "Fabian"
const powerLevel = 9001

const opLevel = 3e20
// a numeric constant has no type
// until it's given one
fmt.Printf("%T\n", opLevel)
```

## Loops

```go
isSkilled := true
for isSkilled {
    fmt.Println("Ready for mission!")
    isSkilled = false
}

for level := 7; level < 9; level++ {
    fmt.Println(level)
    fmt.Println("Leveling up!")
}

for {
    fmt.Println("I’m a Golang Ninja")
    break
}
```