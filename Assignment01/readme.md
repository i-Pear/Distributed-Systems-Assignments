Origin Docs: https://m4q0j7d2yi.feishu.cn/docx/EZqfdEmrGo6JpWxu1ndc9yPvnhg

# Analysis and design

## Requirements

Design a distributed application that generates and displays current time in several machines: One server object/procedure/node generates the current time and displays it on local machines. Other two/more client objects/procedures/nodes get time to be displayed from the server, and display it in digital mode/per second or analogue mode/per minute. Requirements: Use RPC as the communication mechanism. Only authorized clients could get time from the server.

## Design

### Roles

Server: set up a rpc server and listen to a specified tcp port, return local time while receiving rpc calls

Clients: send rpc requests every second, and display the time returned by the server

## Developing environments

I prefer to use GO for programming language, and use gRPC for rpc server implementation.

gRPC is a modern, open source, high-performance RPC framework that can run anywhere. It enables client and server applications to communicate transparently, and simplifies the building of connected systems.

gRPC uses protobuf for encoding messages. protobuf is one of the most popular libraries in production environments.

## Message Formats

Message formats must be defined for server-client communications. Luckily, protobuf provides a user-friendly way to define it, and our design is also simple:

![img](https://m4q0j7d2yi.feishu.cn/space/api/box/stream/download/asynccode/?code=MmFjYjk5OGY4OTZmZmM3YzZlZWRkMzQ1ZWFhZmM0MzFfOFFud2tlTkVmekhMSmdhVUJCcEk3NEU0RlNicmJBM3lfVG9rZW46Ym94Y254cDhuaVZXaXpDS1Bwd2JxVExtNjZnXzE2NjU5MzM5MDg6MTY2NTkzNzUwOF9WNA)

We declared a service named TimeService. In this model, the client simply sends a TimeRequest with a token, which will be used to authorize its permission. The server will return a TimeReply structure after receiving such a TimeRequest, where the "success" field indicates whether the request is valid or not, and the "server_time" field offers a time result in Unix time format (count of seconds passed from 1970-01-01).

# Implement

## Generate protobuf codes

We use the following command to generate code for populating, serializing, and retrieving TimeRequest and TimeReply message types:

```Go
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protocol/protocol.proto
```

## Token authorization

A token is used to verify the identities of clients. To make it easier, we simply agreed on a prefix "client_". For further security, it should be replaced by some encryption algorithms, such as AES.

The isValidToken() function is used to verify tokens:

```Go
func isValidToken(token string) bool {
   return strings.HasPrefix(token, "client_")
}
```

And the getToken() function is for generating tokens for clients:

```Go
func getToken() string {
   return "client_" + strconv.Itoa(rand.Int())
}
```

## Server implementation

Firstly, we try to allocate the port specified by program arguments.

```Go
listen, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
```

Then create a server and register the rpc function onto it:

```Go
timeServer := grpc.NewServer()
pb.RegisterTimeServiceServer(timeServer, &server{})
```

## Client implementation

Firstly, we set up a connection to the server:

```Go
conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
```

Then we initialize a client from the above connection:

```Go
client := pb.NewTimeServiceClient(conn)
ctx, cancel := context.WithCancel(context.Background())
```

Finally, simply call "client.GetTime()" as a local method, and print the formatted time on the screen:

```Go
response, err := c(ctx, &pb.TimeRequest{
   Token: getToken(),
})

timestr := time.Unix(response.GetServerTime(), 0).String()
fmt.Printf("Time from server: %s\n", timestr)
```

## Deploying

We provide a Makefile script for quickly deploying:

```Go
.PHONY: client server

all:
   echo "Usage:"
   echo "Run client: make client"
   echo "Run server: make server"
client:
   go run client/client.go

server:
   go run server/server.go

proto:
   protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protocol/protocol.proto
```

To run a server, you can just run:

```Go
make server
```

For clients, you can also simply run:

```Go
make client
```

For futher development, run the following command to update protobuf codes:

```Go
make proto
```

# Testing

Start up a server on computer A (172.24.88.176):

![img](https://m4q0j7d2yi.feishu.cn/space/api/box/stream/download/asynccode/?code=MzY3MDVmMDE4NTQzYmIzOTQwZTg1YzJiYzEzMzY5NGZfYmpuR0tWWE94N3VUZDIzSm1JcjNUeVZUWW9RUGFoNkdfVG9rZW46Ym94Y25HYmtJWE9DNUh3bXc1eGlYb3Rya1NmXzE2NjU5MzM5MDg6MTY2NTkzNzUwOF9WNA)

Run client with the following command:

```Go
go run client/client.go 172.24.88.176:11451
```

After the client initialized its connection to the server, it prints time data received from the server:

![img](https://m4q0j7d2yi.feishu.cn/space/api/box/stream/download/asynccode/?code=MjRmMGZmNjdjMzkwZTcyOGRhMGVmNTk0MDM5YWU0N2RfMXNsZWRBQ3Y5MTdicVVZTWZrV243R0l6b01GODRHNHNfVG9rZW46Ym94Y25EYTN1ZmpQVU44R2VrTDZwdlkzVUZiXzE2NjU5MzM5MDg6MTY2NTkzNzUwOF9WNA)

Starting another client on another computer will display the same.
And the server is also printing the times it sends to the client, with the clients' tokens:

![img](https://m4q0j7d2yi.feishu.cn/space/api/box/stream/download/asynccode/?code=ZTQ2ZDcwMTllNDlhNTI5NjMzMTYyN2Q4OTZhOTc5N2JfYmdWZGRybWQxVmdwOTZNaGV3a2RTUmZaNlpobm96MVBfVG9rZW46Ym94Y25tTUpnbGV6c3RJMDZobEx3UndFOG5lXzE2NjU5MzM5MDg6MTY2NTkzNzUwOF9WNA)

# Summary

RPC technology provides encapsulation of server-client connections, and calling remote functions is as easy as locally, which greatly facilitates the development of distributed programs.

# Attachments

Opensource project: https://github.com/i-Pear/Distributed-Systems-Assignments