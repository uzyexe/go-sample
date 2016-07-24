ushios/go-sample
===============
golang sample app

Installation
-------------

```bash
$ go get github.com/ushios/go-sample
```


Samples
--------

```bash
$ go-sample -o
Hello, World
```

Options
--------

| option |  |
| ------ | --- |
| o | output `Hello, World` |
| p | make panic |
| s | run http server on localhost:8080 |

Run on Docker
--------------

```bash
$ docker build -t go-sample .
$ docker run -it --rm --name run-go-sample -p 80:8080 go-sample
$ curl localhost:80
Hello, World
```
