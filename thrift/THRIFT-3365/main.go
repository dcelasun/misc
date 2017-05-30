package main

import (
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/dcelasun/misc/thrift/THRIFT-3365/union"
)

func main() {
	t := thrift.NewTMemoryBufferLen(1024)
	//p := thrift.NewTBinaryProtocolFactoryDefault().GetProtocol(t)
	p := thrift.NewTJSONProtocolFactory().GetProtocol(t)

	z := &thrift.TSerializer{
		Transport: t,
		Protocol:  p,
	}

	s := &union.Select{}
	s.Bar = &union.Bar{"bar"}

	if out, err := z.Write(s); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}

