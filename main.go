package main

import (
	"flag"
	"log"

	"github.com/dtchanpura/svelte-tailwind-golang/server"
)

var (
	addr  = flag.String("addr", ":8001", "Address to listen on.")
	debug = flag.Bool("debug", false, "Enables debug")
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	flag.Parse()
	s := server.NewServer(*debug)
	log.Fatal(s.RunServer(*addr))
}
