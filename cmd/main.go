package main

import (
	"SetCronJob/service"
	"flag"
	"fmt"
)

var token, command, ids *string

func init() {
	token = flag.String("token", "", "API Token")
	command = flag.String("command", "", "list-all/list-enabled/disable/enable/run")
	ids = flag.String("ids", "", "list of , seperated cron ids")
}

const (
	ListAll     = "list-all"
	ListEnabled = "list-enabled"
	Enable      = "enable"
	Disable     = "disable"
	Run         = "run"
)

func main() {
	flag.Parse()

	if *token == "" {
		fmt.Println("token missing")
		return
	}

	err := processCommand()

	if err != nil {
		fmt.Println(err)
	}
}

func processCommand() error {
	var err error

	switch *command {
	case ListAll:
		err = service.ListAllCrons(*token)
	case ListEnabled:
		err = service.ListEnabledCrons(*token)
	case Enable:
		err = service.EnableCronList(*token, *ids)
	case Disable:
		err = service.DisableCronList(*token, *ids)
	case Run:
		err = service.RunCronList(*token, *ids)
	default:
		err = fmt.Errorf("invalid command: %s", *command)
	}
	return err
}
