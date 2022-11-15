package email

type Letter struct {
	Mailto  []string `json:"mailto" validate:"required"`
	Subject string   `json:"subject" validate:"required"`
	Body    string   `json:"body" validate:"required"`
}
