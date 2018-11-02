package main

import (
	"context"
	"log"

	pb "github.com/syhlion/2019-10th-ithome30day-grpc-example/echospec"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("連線失敗：%v", err)
	}
	defer conn.Close()

	c := pb.NewEchoClient(conn)

	r, err := c.Echo(context.Background(), &pb.EchoRequest{Msg: "HI HI HI HI"})
	if err != nil {
		log.Fatalf("無法執行 Plus 函式：%v", err)
	}
	log.Printf("回傳結果：%s , 時間:%d", r.Msg, r.Unixtime)

}
