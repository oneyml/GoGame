package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"time"
)

var (
	Args struct {
		Rid   uint64 `short:"i" long:"rid" description:"roleid" default:"1"`
		Rname string `short:"n" long:"rname" description:"role name" default:"yml"`
	}
)

func InitArgs() error {
	_, err := flags.Parse(&Args)
	return err
}

func initArgv() {
	InitArgs()
	log.Printf("rid:%d, rname:%s\n", Args.Rid, Args.Rname)
}

func main() {
	initArgv()
	ClientStart()
	CommandStart()
	for true {
		time.Sleep(time.Second * 10)
	}
}
