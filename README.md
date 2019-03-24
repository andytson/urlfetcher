urlfetcher
==========

A experiment of Go to fetch a list of URLs concurrently and return totals based on them.

Install
-------

Requirements:
* Go
* Dep (Though Go has newer dependency management, this used instead to try earlier usage)

```
dep ensure
go build cmd/urlfetcher
```

Usage
-----

```
urlfetcher <file>
```