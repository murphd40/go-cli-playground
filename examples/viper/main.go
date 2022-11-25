package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

func main() {
	config := flag.String("config", "", "The config file to use")
	flag.Parse()

	viper.SetDefault("frequency", "10s")
	if *config != "" {
		viper.SetConfigFile(*config)
		if err := viper.ReadInConfig(); err != nil {
			log.Fatal(err.Error())
		}
	}

	fmt.Println(viper.GetString("frequency"))

	time.ParseDuration(viper.GetString("frequency"))
}
