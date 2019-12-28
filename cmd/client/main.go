package main

import (
	"flag"
	"github.com/wangke0809/screensharing/client"
)

func main() {
	flags := client.FlagStruct{}
	flag.StringVar(&flags.Host, "host", "224.0.0.1", "udp host ip")
	flag.IntVar(&flags.Port, "port", 9999, "udp listen port")
	flag.Parse()
	client.Start(flags)
}
