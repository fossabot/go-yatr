# go-yatr - Yandex translate
Yandex Translate API on Golang

## Usage

`go get github.com/FlameInTheDark/go-yatr`

```Golang
package main

import (
    "fmt"
    "github.com/FlameInTheDark/go-yatr"
    )

func main() {
    c := yatr.NewContext("yandex-translation-api-key")
    r, err := c.Translate("en", "ru", "Hello, World!")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println(r) // Привет, Мир!
}
```
