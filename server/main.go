/*
 * @Author: Gaston Siffert
 * @Date: 2017-05-29 23:44:26
 * @Last Modified by: Gaston Siffert
 * @Last Modified time: 2017-05-30 23:50:56
 */
package main

import (
	"log"
	"net"
	"os"
	"time"

	"github.com/Vorian-Atreides/morse"
	pb "github.com/Vorian-Atreides/morse_server_raspberrypi/pb"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) Translate(ctx context.Context, in *pb.Body) (*pb.Empty, error) {
	// Init the gpio
	r := raspi.NewAdaptor()
	led := gpio.NewLedDriver(r, os.Args[1])

	// init the morse encoder
	m := morse.New(time.Second / 2)
	// translate the received data
	log.Println(in.Data)
	m.Translate(in.Data, func() {
		led.On()
	}, func() {
		led.Off()
	})

	return &pb.Empty{}, nil
}

const (
	port = ":50051"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("./server [gpio]")
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMorseServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
