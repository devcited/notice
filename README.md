# Notice
Notice is a Go package for creating and manipulating text messages with various styles and colors. It allows you to format messages using ANSI escape sequences and print them with ease.

## Features
- Chainable methods for text styles: bold, dim, italic, underline, inverse, hidden, and strikethrough
- Color manipulations: RGB, Hex, Hue, and HSL
- Output-related methods: Print and Panic

## Installation
```
go get -u github.com/devcited/notice
```

## Usage
Here are some examples of how to use the notice package:

```go
notice.
    Message("Hello World").Blue().Dim().
    Print()
```

```go
now := time.Now()

notice.
    Message("Shutdown sequence completed in ").Dim().
    Message(time.Now().Sub(now)).Bold().BlueBright().
    Print()
```

```go
sig := "SIGTERM"

notice.
    Message("Signal: ").Dim().
    Message(sig).Bold().Underline().Dim().
    Message(". Shutting down.").Dim().
    Print()
```

```go
for j := 0; j < 3; j++ {
    msg := notice.MessageChain{}

    for i := j * 20; i < 360+j*20; i += 2 {
        msg.Message(" ").BG().Hue(float64(i))
    }

    msg.Print()
}
```

```go
for j := 0; j < 3; j++ {
    msg := notice.MessageChain{}

    for i := j * 20; i < 360+j*20; i += 2 {
        msg.Message(" ").BG().HSL(i, 10, 30)
    }

    msg.Print()
}
```

## License
This project is licensed under the MIT License.