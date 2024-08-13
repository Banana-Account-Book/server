package oauth

type OauthInfo struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	ProfileImage string `json:"profileImage"`
	Provider     string `json:"provider"`
}
