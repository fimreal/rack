package email

type Letter struct {
	Mailto  string `json:"mailto" form:"mailto" validate:"required"`
	Subject string `json:"subject" form:"subject" validate:"required"`
	Body    string `json:"body" form:"body" validate:"required"`
}
