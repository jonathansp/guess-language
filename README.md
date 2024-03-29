# guess-language

[![Build Status](https://travis-ci.org/jonathansp/guess-language.svg?branch=master)](https://travis-ci.org/jonathansp/guess-language) [![GoDoc](https://godoc.org/github.com/jonathansp/guess-language?status.svg)](http://godoc.org/github.com/jonathansp/guess-language)

Guess the natural language of a text (idiom).

Based on: https://websvn.kde.org/branches/work/sonnet-refactoring/common/nlp/guesslanguage.cpp?view=markup


## Install

Download and install it:

```sh
go get github.com/jonathansp/guess-language
```

Import it in your code:

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
