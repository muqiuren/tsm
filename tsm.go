package main

import (
	"github.com/muqiuren/tsm/cmd"
	"github.com/muqiuren/tsm/utils"
	"log"
)

func main() {
	defer func() {
		err := utils.Db.Close()
		if err != nil {
			log.Fatalf("db close panic: %s\n", err)
		}

		if err := recover(); err != nil {
			log.Fatalf("catch panic: %s\n", err)
		}
	}()
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	cmd.Execute()
}
