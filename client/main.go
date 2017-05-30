/*
 * @Author: Gaston Siffert
 * @Date: 2017-05-30 22:26:48
 * @Last Modified by: Gaston Siffert
 * @Last Modified time: 2017-05-31 00:13:27
 */
package main

import (
	"context"
	"log"

	"os"

	pb "github.com/Vorian-Atreides/morse_server_raspberrypi/pb"
	"google.golang.org/grpc"
)

const (
	address = "192.168.0.47:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMorseClient(conn)

	if len(os.Args) != 2 {
		log.Fatalln("./client [message]")
	}
	body := pb.Body{Data: os.Args[1]}
	if _, err := c.Translate(context.Background(), &body); err != nil {
		log.Println(err)
	}
}
