package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"github.com/spf13/viper"
	"github.com/mitchellh/go-homedir"

	"github.com/knabben/tradelens/token"
	"github.com/knabben/tradelens/request"
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

	var bearerToken string = token.MainToken.GetBearerToken()

	// Get Consignments with V1
	response := request.Get("/api/v1/consignments", bearerToken)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
}
