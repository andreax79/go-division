# go-division

[![Go Report Card](https://goreportcard.com/badge/github.com/andreax79/go-division)](https://goreportcard.com/report/github.com/andreax79/go-division)

Column Division in Go

### Use Docker

```console
$ docker run -it --rm andreax79/division 3279 25
 3279 | 25
 25   |----
  ----| 131
  77  |
  75  |
   ---|
   29 |
   25 |
    --|
    4 |

3279 : 25 = 131 (4)
```

### Run from source

```console
$ go run division.go 1024 8
 1024 | 8
  8   |----
  ----| 128
  22  |
  16  |
   ---|
   64 |
   64 |
    --|
    0 |

1024 : 8 = 128 (0)
```
