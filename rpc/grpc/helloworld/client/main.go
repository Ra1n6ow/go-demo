package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/ra1n6ow/go-demo/rpc/grpc/helloworld/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

// unaryCallWithMetadata 普通RPC调用客户端metadata操作
func unaryCallWithMetadata(c pb.GreeterClient, name string) {
	// fmt.Println("--- UnarySayHello client---")
	// 创建metadata
	md := metadata.Pairs(
		"token", "iamtoken",
		"request_id", "123456789",
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 将 metadata 添加到 context 中
	ctxWithMd := metadata.NewOutgoingContext(ctx, md)

	// 调用 RPC 方法
	var header, trailer metadata.MD
	resp, err := c.UnarySayHello(
		ctxWithMd,
		&pb.HelloRequest{Name: name},
		grpc.Header(&header),   // 获取响应头
		grpc.Trailer(&trailer), // 获取响应尾
	)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	fmt.Println("header: ", header)

	// 从 header 中获取值
	if t, ok := header["location"]; ok {
		fmt.Println("location from header:")
		for i, e := range t {
			fmt.Printf("%d. %s\n", i, e)
		}
	} else {
		log.Printf("location not found in header")
	}

	// 获取响应结果
	fmt.Println("get response: ", resp.Reply)

	// 从 trailer 中获取值
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}

// bidirectionalWithMetadata 流式RPC调用客户端metadata操作
func bidirectionalWithMetadata(c pb.GreeterClient, name string) {
	// 创建metadata和context.
	md := metadata.Pairs("token", "iamtoken")
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// 使用带有metadata的context执行RPC调用.
	stream, err := c.BidiHello(ctx)
	if err != nil {
		log.Fatalf("failed to call BidiHello: %v\n", err)
	}

	go func() {
		// 当header到达时读取header.
		header, err := stream.Header()
		if err != nil {
			log.Fatalf("failed to get header from stream: %v", err)
		}
		// 从返回响应的header中读取数据.
		if l, ok := header["location"]; ok {
			fmt.Printf("location from header:\n")
			for i, e := range l {
				fmt.Printf(" %d. %s\n", i, e)
			}
		} else {
			log.Println("location expected but doesn't exist in header")
			return
		}

		// 发送所有的请求数据到server.
		for i := 0; i < 5; i++ {
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("failed to send streaming: %v\n", err)
			}
		}
		stream.CloseSend()
	}()

	// 读取所有的响应.
	var rpcStatus error
	fmt.Printf("got response:\n")
	for {
		r, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %s\n", r.Reply)
	}
	if rpcStatus != io.EOF {
		log.Printf("failed to finish server streaming: %v", rpcStatus)
		return
	}

	// 当RPC结束时读取trailer
	trailer := stream.Trailer()
	// 从返回响应的trailer中读取metadata.
	if t, ok := trailer["timestamp"]; ok {
		fmt.Printf("timestamp from trailer:\n")
		for i, e := range t {
			fmt.Printf(" %d. %s\n", i, e)
		}
	} else {
		log.Printf("timestamp expected but doesn't exist in trailer")
	}
}

func main() {
	conn, err := grpc.NewClient("localhost:7200", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// unaryCallWithMetadata(client, "Du")
	bidirectionalWithMetadata(client, "Du")
}
