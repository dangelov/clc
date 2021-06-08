# CaLCulator

`clc` is a simple calulator that reads an input file and prints the result.

## Example

```bash
$ cat input.txt
# UNIT: USD
  1000 # Salary
  -320 # Rent
# 3 other expenses (comment)
 -1.99 # Netflix
 -0.99 # Apple Arcade
  -200 # Groceries
$ clc input.txt
TOTAL: 477.02 USD
```

## Building

`go build main.go`

To cross-build for Windows:
`env GOOS=windows GOARCH=amd64 go build main.go`

## Todo

- [x] Section labels
- [ ] Use case examples
- [ ] More operations
