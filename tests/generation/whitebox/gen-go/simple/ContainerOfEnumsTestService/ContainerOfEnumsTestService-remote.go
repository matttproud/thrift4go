/* Autogenerated by Thrift Compiler (0.9.0-dev)
 *
 * DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
 */
package main

import (
	"flag"
	"fmt"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"thrift"
	"thriftlib/simple"
)

func Usage() {
	fmt.Fprint(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:\n")
	flag.PrintDefaults()
	fmt.Fprint(os.Stderr, "Functions:\n")
	fmt.Fprint(os.Stderr, "  echo(message *ContainerOfEnums) (retval32 *ContainerOfEnums, err error)\n")
	fmt.Fprint(os.Stderr, "\n")
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var help bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.BoolVar(&help, "help", false, "See usage string")
	flag.Parse()
	if help || flag.NArg() == 0 {
		flag.Usage()
	}

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprint(os.Stderr, "Error parsing URL: ", err.Error(), "\n")
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprint(os.Stderr, "Error parsing URL: ", err.Error(), "\n")
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		addr, err := net.ResolveTCPAddr("tcp", fmt.Sprint(host, ":", port))
		if err != nil {
			fmt.Fprint(os.Stderr, "Error resolving address", err.Error())
			os.Exit(1)
		}
		trans, err = thrift.NewTNonblockingSocketAddr(addr)
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprint(os.Stderr, "Error creating transport", err.Error())
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprint(os.Stderr, "Invalid protocol specified: ", protocol, "\n")
		Usage()
		os.Exit(1)
	}
	client := simple.NewContainerOfEnumsTestServiceClientFactory(trans, protocolFactory)
	if err = trans.Open(); err != nil {
		fmt.Fprint(os.Stderr, "Error opening socket to ", host, ":", port, " ", err.Error())
		os.Exit(1)
	}

	switch cmd {
	case "echo":
		if flag.NArg()-1 != 1 {
			fmt.Fprint(os.Stderr, "Echo requires 1 args\n")
			flag.Usage()
		}
		arg33 := flag.Arg(1)
		mbTrans34 := thrift.NewTMemoryBufferLen(len(arg33))
		defer mbTrans34.Close()
		_, err35 := mbTrans34.WriteString(arg33)
		if err35 != nil {
			Usage()
			return
		}
		factory36 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt37 := factory36.GetProtocol(mbTrans34)
		argvalue0 := simple.NewContainerOfEnums()
		err38 := argvalue0.Read(jsProt37)
		if err38 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Echo(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprint(os.Stderr, "Invalid function ", cmd, "\n")
	}
}
