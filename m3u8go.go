package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/grafov/m3u8"
)

func main() {
	var m3u8FilePath string
	flag.StringVar(&m3u8FilePath, "i", "", "m3u8 master playlist to filepath")
	flag.Parse()

	f, err := os.Open(m3u8FilePath)
	if err != nil {
		log.Fatalln(err)
	}

	p, t, err := m3u8.DecodeFrom(f, true)
	if err != nil {
		log.Fatalln(err)
	}

	if t != m3u8.MASTER {
		log.Fatalf("not support file type [%d]", t)
	}

	for _, v := range p.(*m3u8.MasterPlaylist).Variants {
		fmt.Println(v.URI)
	}
}
