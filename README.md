# guess-language [![Build Status](https://travis-ci.org/jonathansp/guess-language.svg?branch=master)](https://travis-ci.org/jonathansp/guess-language)
Guess the natural language (idiom) of a text.


## Install
1. Download and install it:

```sh
go get github.com/jonathansp/guess-language
```
2. Import it in your code:

```go
import "github.com/jonathansp/guess-language"
```

## Usage
```go
package main

import (
    "fmt"
    "github.com/jonathansp/guess-language"
)

func main () {
        lang, _ := guesslanguage.Parse("We know what we are, but know not what we may be.")
        fmt.Print(lang)
}
```

## Authors

Jonathan Simon Prates (@jonathansp)
