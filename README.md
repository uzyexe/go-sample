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
| s | run http server on localhost:8080 (using env:GO_SAMPLE_HTTP_PORT if you want to change the port) |

Run on Docker
--------------

[ushios/go-sample@docker](https://hub.docker.com/r/ushios/go-sample/)

```bash
$ docker run -p 80:8080 ushios/go-sample
```

Checking `localhost`

```bash
$ curl localhost
Hello, World
```

##### local

```bash
$ docker build -t go-sample .
$ docker run -it --rm --name run-go-sample -p 80:8080 go-sample
```
