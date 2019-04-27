goem - EM algorithm implementation for golang
===
[![GoDoc](https://godoc.org/github.com/6br/goem/goem?status.svg)](https://godoc.org/github.com/6br/goem/goem)
[![Build Status](https://drone.io/github.com/6br/goem/status.png)](https://drone.io/github.com/6br/goem/latest)
[![Coverage Status](https://coveralls.io/repos/github/6br/goem/badge.svg?branch=master)](https://coveralls.io/github/6br/goem?branch=master)

goem requires Go version 1.4.2 or greater

## Description
This package is an implementation of golang to use EM algorithm

EM-algorithm(expectation maximization) is a method for finding maximum likelihood estimates of hidden statistical parameters.

<img src="https://cloud.githubusercontent.com/assets/12047794/14041984/296de90a-f2b9-11e5-852d-a85dc021d6b1.gif" width="700">

## Usage of commandline
```sh
go run main.go -m=1.0 -d=pic/ < space_separated.txt
```

Two-dimentional example data of `space_separated.txt` is below.

```space_separated.txt
0.471726116612005	0.266595928855752
0.308522155121713	0.0859873278429996
0.444219121551699	0.92693805785837
0.232566914860858	0.314525852686857
0.559125332185874	0.867327637203401
0.399080742257205	0.817437666085725
0.536394027063934	0.474805220342716
```

Important options are below.

* verbose(v bool): if it is true, graphs might be output in pic/ and show the result implicit.
* meanshift(m float64): you have to try to search the suitable parameters to avoid getting localized solution.
* directory(d string): where to save plotted images of EM-algorithm iteration.

Use the following command if you want to know more details of flags.

```sh
go run main.go --help
```

### Convert Plotted PNG to Animation GIF
Use ImageMagick.
```sh
convert -delay 50 *.png animation.gif
```

## Install
Please git clone. If you want to use it as a package, use go get.

```sh
go get github.com/6br/goem
```

## LICENSE
Please see LICENSE. 

Using library packages below.

* github.com/gonum/matrix 
* github.com/gonum/plot

(Copyright (c)2013 The gonum Authors. All rights reserved.)
