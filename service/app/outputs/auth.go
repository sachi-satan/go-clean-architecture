package outputs

type AuthSignUpOutputData struct {
	AccessToken string `json:"accessToken"`
}

func NewAuthSignUpOutputData(accessToken string) *AuthSignUpOutputData {
	return &AuthSignUpOutputData{
		AccessToken: accessToken,
	}
}

type AuthSignInOutputData struct {
	AccessToken string `json:"accessToken"`
}

func NewAuthSignInOutputData(accessToken string) *AuthSignInOutputData {
	return &AuthSignInOutputData{
		AccessToken: accessToken,
	}
}
