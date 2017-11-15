package main

import (
	"log"
	"os"

	"crypto/tls"
	"crypto/x509"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"

	"io/ioutil"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {

	certificate, err := tls.LoadX509KeyPair(
		"out/127.0.0.1.crt",
		"out/127.0.0.1.key",
	)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile("out/My_Root_CA.crt")
	if err != nil {
		log.Fatalf("failed to read ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		log.Fatal("failed to append certs")
	}

	transportCreds := credentials.NewTLS(&tls.Config{
		ServerName:   "davesdomain.com",
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(transportCreds))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
