package main

import (
	"kang-by-xoverse/events-server/tcp"
	"kang-by-xoverse/events-server/utils"
)

func main() {
	utils.LoadDotEnv()
	rdb, close := utils.GetRedisClient()

	go tcp.CreateEventsServer(rdb)

	close()
}
