package main

import (
	"els/internal/app/service"
	"els/internal/repository/elastic"
	"els/internal/repository/pg"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	closeDB := pg.Init()
	defer closeDB()

	elastic.Init()

	r := service.NewHTTPHandler()

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		errs <- r.Run(":3000")
	}()

	log.Println("exit", <-errs)
}
