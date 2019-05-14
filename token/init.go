package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Init() {
	getBearerToken()
}

func getBearerToken() (string) {
	if MainToken.bearerToken == "" {
		loginToSM()
	}

	err := checkSMToken()
	if err != nil {
		loginToSM()
		return getBearerToken()
	}
	return MainToken.bearerToken
}

func loginToSM() {
	responseIAM, err := loginIntoCloudIAM()
	defer responseIAM.Body.Close()

	if err != nil {
		fmt.Println(err)
	}

	smRootURL := MainToken.GetSMRootURL()
	solutionID := MainToken.GetSolutionID()
	organizationID := MainToken.GetOrganizationID()

	exchangeArgs := fmt.Sprintf(
		"%s/%s/organization/%s",
		"/onboarding/v1/iam/exchange_token/solution/", solutionID, organizationID)

	client := &http.Client{}
	req, err := http.NewRequest("POST", smRootURL + exchangeArgs, responseIAM.Body)
	req.Header.Add("Content-Type", "application/json")

	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	err = json.Unmarshal([]byte(body), &MainToken)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("MAIN TOKEN ", MainToken)
}

func checkSMToken() error {
	return nil
}

func loginIntoCloudIAM() (*http.Response, error) {
	cloudIAMUrl := MainToken.GetCloudIAMUrl()
	apiKey := MainToken.GetApiKey()

	values := url.Values{}
	values.Add("grant_type", "urn:ibm:params:oauth:grant-type:apikey")
	values.Add("apikey", apiKey)

	client := &http.Client{}
	response, err := client.PostForm(cloudIAMUrl, values)
	if err != nil {
		fmt.Println(err)
	}
	return response, err
}