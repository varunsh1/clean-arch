package main

import (
	"fmt"
	"os"

	config "github.com/spf13/viper"
)

func initEnv() {
	env = os.Getenv("ENV") // Clean Architecture Environment
	if env != production && env != staging {
		env = development
	}
}

func initConfigFile() {
	var err error
	fileDir := "files/etc/"
	fileName := "config.%s.toml"
	fileName = fmt.Sprintf(fileName, env)

	// viper configuration setting
	config.AddConfigPath(fileDir)  // add the directory containing the config file
	config.SetConfigName(fileName) // set the base name of the config file (without directory)
	config.SetConfigType("toml")   // REQUIRED if the config file does not have the extension in the name

	err = config.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func main() {
	// init configs
	initEnv()
	initConfigFile()
	fmt.Println("Db connection", config.GetString("dbconnection.user.DBName"))
}
