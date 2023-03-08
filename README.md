# gRPC-In-Go-Lang
Basic gRPC call In Go Lang

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