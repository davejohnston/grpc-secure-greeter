# GRPC Secure Greeter

This is an example that extends the grpc greeter demo to run over TLS.
It uses mutual auth to negotiate the handshake.

# Prerequisites

The demo comes with pre-generated certs for the client and server signed
by a shared key.  You can use these out of the box, or generate your own keys using the following steps:

## Get Certstrap
https://github.com/square/certstrap.git

## Create Certificate Authortity
    certstrap init --common-name "My Root CA"

## Create and Sign Server Certs

### Create a cert for the server
    certstrap request-cert --domain mytestdomain.com

### Sign the server cert
    certstrap sign --CA "My Root CA" mytestdomain.com

## Create and Sign Client Certs

### Create a cert for the client
    certstrap request-cert --ip 127.0.0.1

### Sign the client cert
    certstrap sign --CA "My Root CA" 127.0.0.1


The certs will be created in a directory called out.
They should be moved to the relevant directories under client/server i.e

    mv out/127.0.0.1.crt out/127.0.0.1.key client/out
    mv out/mytestdomain.com.crt  out/mytestdomain.com.key server/out

Both the server and client need to access the CA cert
    
    cp out/My_Root_CA.crt server/out/My_Root_CA.crt
    cp out/My_Root_CA.crt client/out/My_Root_CA.crt
