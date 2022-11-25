package tcp

import (
	"fmt"
	"net"
	"os"

	"kang-by-xoverse/events-server/utils"

	"github.com/go-redis/redis/v8"
)

func CreateEventsServer(rdb *redis.Client) {
	host := utils.GetEnv("TCP_HOST", "localhost")
	port := utils.GetEnv("TCP_PORT", "5003")

	server, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()

	fmt.Println("Events Server started Running on tcp:" + port)

	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("client connected through tcp")
		go handleConnectedClient(rdb, connection)
	}

}

func handleConnectedClient(rdb *redis.Client, connection net.Conn) {
	utils.Listen(rdb, func(data string) {
		connection.Write([]byte(data))
	})
}
