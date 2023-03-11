# gRPC-In-Go-Lang

- [gRPC-In-Go-Lang](#grpc-in-go-lang)
  - [What is gRPC?](#what-is-grpc)
    - [RPC](#rpc)
    - [gRPC](#grpc)
    - [Why it's getting popular?](#why-its-getting-popular)
    - [Wait, what is HTTP/2 protocol?](#wait-what-is-http2-protocol)
    - [I heard people are using it instead of REST API.](#i-heard-people-are-using-it-instead-of-rest-api)
    - [Why not keep using REST?](#why-not-keep-using-rest)
- [Basic gRPC call In Go Lang](#basic-grpc-call-in-go-lang)
  - [Create proto file](#create-proto-file)
  - [Write gRPC server code](#write-grpc-server-code)
  - [How to test our grpc end point](#how-to-test-our-grpc-end-point)

<small><i><a href='http://ecotrust-canada.github.io/markdown-toc/'>Table of contents generated with markdown-toc</a></i></small>

## What is gRPC?
It's google RPC. 	:sweat_smile: </br>
Hey, I know about google. But what is RPC?
### RPC
- It stands for Remote Procedural Call.
- It uses a form of function call instead of the popular HTTP calls.
- It uses IDL(Interface Definition Lang.) as a form of contract.

### gRPC
- It's a google RPC.
- It uses HTTP/2 protocol.
- It uses Protocol Buffer (ProtoBuf) format for sending/receiving messages.

### Why it's getting popular?
- Easy - It's a simple function call.
- Support - It is supported in a lot of languages.
- Fast - It uses HTTP/2 protocol. 

### Wait, what is HTTP/2 protocol?
- Well it's HTTP version 2. :facepalm:
- It's faster than HTTP 1 as it comes with the following features:
  - Enable request and response multiplexing - Servers can send multiple messages in a single request.
  - Header compression
  - binary protocol - Send message in 0/1 format instead of text-based format
- It uses Protocol Buffer (ProtoBuf) as the IDL.

### I heard people are using it instead of REST API.
- Well REST is an architectural style.
- It defines protocols to talk between client and server.
- It uses HTTP/1 protocol.
- It uses JSON format for sending/receiving messages.

### Why not keep using REST?
- Well it's simple, easy to use and there are many tools to test it easily.
- Issues: 
  - Sometimes the payload becomes huge and it decreases the performance.
  - It's unary, that is we can send one request at a time.
  - REST uses HTTP/1 protocol which does three-way handshaking for the first message, making it slower.

Note: It's best to use REST when communicating between browser and back-end. gRPC is best for inter-microservice communication.


# Basic gRPC call In Go Lang
## Create proto file

1. Create go.mod file
   go mod init github.com/ughosh/grcp-learn

2. Create a file `hello.proto`

3. In add  `hello.proto`
   1. Add the version of proto file
        ```
            syntax = "proto3"; 
        ```
   2. Add folder name where we want to save generated files
        ```
          option go_package = "github.com/ughosh/grcp-learn/invoicer";
        ```
   3.  Define service
    ```
    service Invoicer {
        rpc Hello(HelloRequest) returns (HelloResponse);
    } 
    ```
    4. Define request and response as message
    ```
    message HelloRequest{
        string name = 1;
    }

    message HelloResponse {
    string msg = 1;
    }
    ```
 4. Create a folder name `hello`
 5. Run the command to generate proto files based on the description given in our `hello.proto` file
     ```
     protoc \
    --go_out=hello \
    --go_opt=paths=source_relative \
    --go-grpc_out=hello \
    --go-grpc_opt=paths=source_relative \
    hello.proto
    ```
 6. If error occur, try running
    ```
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@lates
    ```
 7. The generate file will be showing some error, in case the lib is not present. Do `go mod tidy` after this to clean-up `go.mod` file
    ```
     go get -u google.golang.org/grpc
    ```
 8. If we make any changes in our `hello.proto` file, we have to run command #5. To save our time, we can save the entire command in Makefile. Create a `Makefile` and add the command
   ```
    generate_grpc_code_hello:
	protoc \
    --go_out=hello \
    --go_opt=paths=source_relative \
    --go-grpc_out=hello \
    --go-grpc_opt=paths=source_relative \
    hello.proto
   ```
   Now, if we want to do any updates, run `make generate_grpc_code_hello`

## Write gRPC server code
Add main.go file and run that file using `go run main.go`
This will start our server which is listening  to port 8080

## How to test our grpc end point
  To test it, we will be using grpcurl (https://github.com/fullstorydev/grpcurl)
  Run the command in another terminal
  ```
  grpcurl -plaintext -d '{"name": "Test"}' localhost:8080 Hello.Hello
  ```

  You will output like 
  ```
    {
        "msg": "Hello Test"
    }
  ```

Congratulations! You created your first gRPC call in Go lang :tada: 