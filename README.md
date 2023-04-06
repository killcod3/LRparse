# LRparse

LRparse is a Go package that provides two functionalities for parsing strings between specified left and right delimiters. The package is lightweight and easy to use, offering a simple API for extracting matches from a given string.

## Installation

To install LRparse, simply run the following command in your terminal:

```bash
go get github.com/killcod3/LRparse
```

This command will fetch and install the LRparse package in your Go workspace.

## Usage

```Go
import "github.com/killcod3/LRparse"
```

Then, you can call the GetFirstMatch or MatchRecursive functions with the required parameters:

```Go
s := "This is a (sample) string with one or (multiple) matches."

firstMatch, err := matchresults.GetFirstMatch(s, "(", ")") // results : sample

allMatches, err := matchresults.MatchRecursive(s, "(", ")") // results : sample , multiple
```

## License

This project is licensed under the MIT License - see the LICENSE file

