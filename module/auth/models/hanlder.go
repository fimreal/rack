package models

type AddUserData struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname" `
	PhoneNumber string `json:"phone_number"`
}

type UpdateUserData struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	PhoneNumber string `json:"phone_number"`
}

type LoginData struct {
	UserIdentifier string `json:"user_identifier" binding:"required"` // username or  email
	Password       string `json:"password" binding:"required"`
	IdentifiedCode string `json:"identified_code"`
}

type LoginResult struct {
	ID           uint   `json:"id" binding:"required"`
	Username     string `json:"username" binding:"required"`
	Email        string `json:"email" binding:"required"`
	PhoneNumber  string `json:"phone_number"`
	Nickname     string `json:"nickname"`
	Role         int    `json:"role" binding:"required"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
