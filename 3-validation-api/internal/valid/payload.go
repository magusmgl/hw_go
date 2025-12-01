package valid

type SendEmailRequest struct {
	To string `json:"to" validation:"required"`
}
