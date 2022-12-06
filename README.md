
# CLI Colors

## The ultimate simplest color api set for your terminal.
This is very simple almost toy based solution for coloring your terminal, I've 
created this just to use in my toy lang: ``cat-lang`` check it out [here](https://github.com/Grubba27/cat-lang)

## How to use ?
```go
import (
    paint "github/painter"
)

func func main() {
    fmt.Printf(paint.InPurple("Some Text"))
}

```
### Important notes

It returns a ready to use in terminal string, if you use ``%s`` notation it will not produce a nice result. 
to make it look good I recommend concatenating strings using the + notation.
Also it will not work in Windows because of how bash works
