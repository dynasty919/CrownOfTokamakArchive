package main

import (
	"CrownOfTokamak/util"
	"context"
	"log"
	"net"
	"time"

	"CrownOfTokamak/server/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAnsServiceServer
	InfoChan chan util.AnsInfo
}

var cnt = 1

func (s *server) ProcessAnsList(ctx context.Context, ansList *pb.AnsList) (*pb.Ans, error) {
	// 在这里处理接收到的 AnsList 数据，
	for _, ans := range ansList.Arr {
		var newAns util.AnsInfo
		newAns.Author = ans.Author
		newAns.Title = ans.Title
		newAns.Content = ans.Content
		newAns.PostTime = time.Now()
		newAns.Counter = cnt
		cnt++
		newAns.Id = util.ContentSha1(ans.Content)

		log.Printf("ans received %v", newAns.Title)
		s.InfoChan <- newAns
	}

	// 在这里返回响应，这里只是简单地返回接收到的第一个消息
	if len(ansList.Arr) > 0 {
		return ansList.Arr[0], nil
	}

	// 如果没有收到消息，可以返回一个空的 Ans
	return &pb.Ans{}, nil
}

func GrpcServer() {
	listen, err := net.Listen("tcp", ":1111")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	} else {
		log.Printf("server listening tcp:1111")
	}

	ch := make(chan util.AnsInfo)
	s := &server{
		InfoChan: ch,
	}

	grpcServer := grpc.NewServer()

	log.Printf("Grpc Server created")

	pb.RegisterAnsServiceServer(grpcServer, s)

	log.Printf("Grpc Server registed, servicing %v", grpcServer.GetServiceInfo())

	go httpServer(ch)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
