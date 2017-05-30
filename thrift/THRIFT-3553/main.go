package main

import (
	"fmt"
	"github.com/dcelasun/misc/thrift/THRIFT-3553/bug"
)

var (
	_ = bug.NewTestsResponse()
)

func main() {
	fmt.Println("ok")
}

