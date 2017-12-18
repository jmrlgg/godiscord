package main

import (
	"fmt"

	"github.com/jmrlgg/godiscord/bot"
	"github.com/jmrlgg/godiscord/config"
)

func main() {

	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return

}
