package main

/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements. See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership. The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License. You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

import (
	"crypto/tls"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	// "strconv"
	// "time"
	"github.com/dcelasun/misc/thrift/THRIFT-4188/tutorial"
)

func handleClient(client *tutorial.CalculatorClient, i int32) (err error, s int32) {
	client.Ping()
	// sum, _ := client.Add(i, 1)
	return nil, int32(0)
}
func handleClient1(client *tutorial.CalculatorClient) (err error) {
	client.Ping()
	fmt.Println("ping()")

	sum, _ := client.Add(1, 1)
	fmt.Print("1+1=", sum, "\n")

	work := tutorial.NewWork()
	work.Op = tutorial.Operation_DIVIDE
	work.Num1 = 1
	work.Num2 = 0
	quotient, err := client.Calculate(1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operation:", err)
		}
		return err
	} else {
		fmt.Println("Whoa we can divide by 0 with new value:", quotient)
	}

	work.Op = tutorial.Operation_SUBTRACT
	work.Num1 = 15
	work.Num2 = 10
	diff, err := client.Calculate(1, work)
	if err != nil {
		switch v := err.(type) {
		case *tutorial.InvalidOperation:
			fmt.Println("Invalid operation:", v)
		default:
			fmt.Println("Error during operation:", err)
		}
		return err
	} else {
		fmt.Print("15-10=", diff, "\n")
	}

	log, err := client.GetStruct(1)
	if err != nil {
		fmt.Println("Unable to get struct:", err)
		return err
	} else {
		fmt.Println("Check log:", log.Value)
	}
	return err
}

func runClient(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string, secure bool, i int32) (e error, s int32) {
	var transport thrift.TTransport
	var err error
	if secure {
		cfg := new(tls.Config)
		cfg.InsecureSkipVerify = true
		transport, err = thrift.NewTSSLSocket(addr, cfg)
	} else {
		transport, err = thrift.NewTSocket(addr)
	}
	if err != nil {
		fmt.Println("Error opening socket:", err)
		return err, int32(0)
	}
	transport = transportFactory.GetTransport(transport)
	defer transport.Close()
	if err := transport.Open(); err != nil {
		return err, int32(0)
	}
	client := tutorial.NewCalculatorClientFactory(transport, protocolFactory)
	_, c := handleClient(client, i)

	return nil, c

}
