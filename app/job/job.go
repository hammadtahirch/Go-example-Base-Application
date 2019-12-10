package main

import (
	machinery "github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/hammadtahirch/golang_basic_app/app/services"
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
	server.RegisterTask("SendEmailJob", SendEmailJob)

	worker := server.NewWorker("worker-1", 10)
	err = worker.Launch()
	if err != nil {
		//todo:handle errors
	}

}

//SendEmailJob ...
func SendEmailJob(email string) (string, error) {
	sender := services.NewSender("hammad.tahir.ch@gmail.com", "Honey@03227058541")

	//The receiver needs to be in slice as the receive supports multiple receiver
	Receiver := []string{"hammad.tahir.ch+testing@gmail.com", "hammad.tahir.ch+testing2@gmail.com", "hammad.tahir.ch+testing3@gmail.com"}

	Subject := "Testing HTLML Email from golang"
	message := `
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>This is the body<br>
	<div class="moz-signature"><i><br>
	<br>
	Regards<br>
	Alex<br>
	<i></div>
	</body>
	</html>
	`
	bodyMessage := sender.WriteHTMLEmail(Receiver, Subject, message)

	sender.SendMail(Receiver, Subject, bodyMessage)
	return "Hello " + email + "!", nil
}
