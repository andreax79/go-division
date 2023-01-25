# go-division

Column Division in Go

## Run with Docker

```console
$ docker run -it --rm andreax79/division 3275 25
 3275 | 25
 25   |----
  ----| 131
  77  |
  75  |
   ---|
   25 |
   25 |
    --|
    0 |

3275 : 25 = 131 (0)
```

## Run with Go

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
