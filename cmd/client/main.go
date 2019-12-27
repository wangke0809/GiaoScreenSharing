package main

import (
	"github.com/wangke0809/screensharing/client"
)

const srvAddr = "224.0.0.1:9999"

func main() {
	client.Start(srvAddr)
}
