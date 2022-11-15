package randomstring

type RandomCode struct {
	Length int `json:"length" validate:"required"`
	// charpool string `json:"charpool"`
	Chartype string `json:"chartype"`
}
