# go-yatr - Yandex translate
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FFlameInTheDark%2Fgo-yatr.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2FFlameInTheDark%2Fgo-yatr?ref=badge_shield)

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


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2FFlameInTheDark%2Fgo-yatr.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2FFlameInTheDark%2Fgo-yatr?ref=badge_large)