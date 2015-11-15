goem -- EM algorithm implementation for golang
===
This package is an implementation of golang to use EM algorithm

# Description
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

# Install
Please git clone.
If you want to use as package, use go-get.
```sh
go get github.com/6br/goem
```
