package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"
	"github.com/knabben/tradelens/token"
)

var cfgFile string

func init() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName(".hodl")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func main() {
	token.MainToken.SetApiKey(viper.GetString("apiKey"))
	token.MainToken.SetSMRootURL(viper.GetString("smRootUrl"))
	token.MainToken.SetOrganizationID(viper.GetString("organizationId"))
	token.MainToken.SetSolutionID(viper.GetString("solutionId"))
	token.MainToken.SetCloudIAMUrl(viper.GetString("cloudIAMUrl"))
	token.Init()
}
