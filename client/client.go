package main

import (
	"ThriftDemo/example"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"log"
	"net"
)

const (
	HOST = "192.168.4.157"
	PORT = "8030"
)

func main() {
	tSocket, err := thrift.NewTSocket(net.JoinHostPort(HOST, PORT))
	if err != nil {
		log.Fatalln("tSocket error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	transport, err := transportFactory.GetTransport(tSocket)
	if err != nil {
		log.Fatalln(err.Error())
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	client := example.NewFormatDataClientFactory(transport, protocolFactory)

	if err := transport.Open(); err != nil {
		log.Fatalln("Error opening:", HOST+":"+PORT)
	}
	defer transport.Close()

	data := example.Data{Text: "hello,world!"}
	d, err := client.DoFormat(&data)
	fmt.Println(d.Text)
}
