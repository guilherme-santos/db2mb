// +build !no_ldflags

package main

// #cgo LDFLAGS: -L./mysql/mysql-binary-log-events/lib
import "C"
