package token

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
