package token

type TokenData struct {
	apiKey string
	cloudIAMURL string
	smRootUrl string
	organizationID string
	solutionID string
	bearerToken string `json:"onboarding_token"`
}

var (
	MainToken TokenData = TokenData{}
	CloudBearer map[string]interface{}

)

func (m *TokenData) SetApiKey(apiKey string) {
	m.apiKey = apiKey
}

func (m *TokenData) GetApiKey() string {
	return m.apiKey
}

func (m *TokenData) SetSMRootURL(smRootUrl string) {
	m.smRootUrl = smRootUrl
}

func (m *TokenData) GetSMRootURL() string {
	return m.smRootUrl
}

func (m *TokenData) SetOrganizationID(organizationID string) {
	m.organizationID = organizationID
}

func (m *TokenData) GetOrganizationID() string {
	return m.organizationID
}

func (m *TokenData) SetSolutionID(solutionID string) {
	m.solutionID = solutionID
}

func (m *TokenData) GetSolutionID() string {
	return m.solutionID
}

func (m *TokenData) SetCloudIAMUrl(cloudIAM string) {
	m.cloudIAMURL = cloudIAM
}

func (m *TokenData) GetCloudIAMUrl() string {
	return m.cloudIAMURL
}
