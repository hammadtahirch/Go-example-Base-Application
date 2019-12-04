package main

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
)

func main() {

	var cnf = config.Config{
		Broker:        "redis://redis-container:6379",
		ResultBackend: "redis://redis-container:6379",
	}

	server, err := machinery.NewServer(&cnf)
	if err != nil {
		//todo:handle errors
	}

	server.RegisterTask("Say", Say)

	worker := server.NewWorker("worker-1", 10)
	err = worker.Launch()
	if err != nil {
		//todo:handle errors
	}

}

func Say(name string) (string, error) {
	return "Hello " + name + "!", nil
}
