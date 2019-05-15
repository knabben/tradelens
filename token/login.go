package token

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func loginToSM() {
	responseIAM, err := loginIntoCloudIAM()
	defer responseIAM.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	smRootURL := MainToken.GetSMRootURL()
	solutionID := MainToken.GetSolutionID()
	organizationID := MainToken.GetOrganizationID()

	exchangeArgs := fmt.Sprintf(
		"%s/%s/organization/%s",
		"/onboarding/v1/iam/exchange_token/solution/", solutionID, organizationID)

	client := &http.Client{}
	req, err := http.NewRequest("POST", smRootURL+exchangeArgs, responseIAM.Body)
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	var s map[string]interface{}
	if err := json.Unmarshal(body, &s); err != nil {
		log.Fatal(err)
	}

	MainToken.bearerToken = s["onboarding_token"].(string)
}

func checkSMToken() error {
	// TODO - implement check and retry
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
		log.Fatal(err)
	}
	return response, err
}
