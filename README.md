goem - EM algorithm implementation for golang
===
[![GoDoc](https://godoc.org/github.com/6br/goem/goem?status.svg)](https://godoc.org/github.com/6br/goem/goem)

goem requires Go version 1.4.2 or greater

# Description
This package is an implementation of golang to use EM algorithm

EM-algorithm(expectation maximization) is a method for finding maximum likelihood estimates of hidden statistical parameters.

# Usage
```sh
go run main.go -m=1.0 < space_separated.txt
```

Important options are below.

* verbose(v bool): if it is true, graphs might be output in pic/ and show the result implicit.
* meanshift(m float64): you have to try to search the suitable parameters to avoid getting localized solution.

Please use below if you want to know more.

```sh
go run main.go --help
```

## Convert Plotted PNG to Animation GIF
Use ImageMagick.
```sh
convert -delay 50 *.png animation.gif
```

# Install
Please git clone. If you want to use as package, use go get.

```sh
go get github.com/6br/goem
```

# LICENSE
Please see LICENSE.

Using library packages below.

* github.com/gonum/matrix 
* github.com/gonum/plot

(Copyright (c)2013 The gonum Authors. All rights reserved.)
