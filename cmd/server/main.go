package main

import (
	"flag"
	"github.com/wangke0809/screensharing/server"
)

func main() {
	flags := server.FlagStruct{}
	flag.StringVar(&flags.Host, "host", "224.0.0.1", "udp host ip")
	flag.IntVar(&flags.Port, "port", 9999, "udp listen port")
	flag.IntVar(&flags.Display, "display", 0, "screen display index")
	flag.IntVar(&flags.BlockSize, "block", 150, "screen transfer block size")
	flag.IntVar(&flags.Quality, "quality", 75, "jpeg compress quality")
	flag.Parse()
	server.Start(flags)
}
