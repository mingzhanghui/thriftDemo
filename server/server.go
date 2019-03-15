package main

import (
	"ThriftDemo/example"
	"strings"
	"git.apache.org/thrift.git/lib/go/thrift"
	"fmt"
	"log"
)

type FormatDataImpl struct {}

func (fdi *FormatDataImpl) DoFormat(data *example.Data) (r *example.Data, err error){
	var rData example.Data
	rData.Text = strings.ToUpper(data.Text)

	return &rData, nil
}

//func (fdi *FormatDataImpl)Process(ctx context.Context, in, out thrift.TProtocol) (bool, thrift.TException) {
//	return true, nil
//}

const (
	HOST = "localhost"
	PORT = "8080"
)

func main() {

	handler := &FormatDataImpl{}
	processor := example.NewFormatDataProcessor(handler)
	serverTransport, err := thrift.NewTServerSocket(HOST + ":" + PORT)
	if err != nil {
		log.Fatalln("Error:", err)
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	fmt.Println("Running at:", HOST + ":" + PORT)
	server.Serve()
}