# gRPC Hello World

## Prerequisites

```
$ sudo apt install protobuf-compiler
$ protoc --version
```


```
$ vim /etc/profile.d/custom.sh

export GO_HOME=/usr/local/go/bin
export O=/d/odrive
export PATH=$PATH:$GO_HOME
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
$ go get google.golang.org/grpc
```

## Setup and Run

Follow these setup to run the [quick start][] example:

 1. Get the code:

    ```console
    $ go get google.golang.org/grpc/examples/helloworld/greeter_client
    $ go get google.golang.org/grpc/examples/helloworld/greeter_server
    ```

 2. [Generate gRPC code](https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code]):

    ```
    protoc --go_out=. --go_opt=paths=source_relative \
      --go-grpc_out=. --go-grpc_opt=paths=source_relative \
      helloworld/helloworld.proto
    ```

 3. Run the server:

    ```console
    $ $(go env GOPATH)/bin/greeter_server &
    ```

 4. Run the client:

    ```console
    $ $(go env GOPATH)/bin/greeter_client
    Greeting: Hello world
    ```

For more details (including instructions for making a small change to the
example code) or if you're having trouble running this example, see [Quick
Start][].

[quick start]: https://grpc.io/docs/languages/go/quickstart
