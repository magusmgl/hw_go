package valid

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"net/textproto"
	"validation/api/configs"
	"validation/api/pkg/req"
	"validation/api/pkg/res"

	"github.com/jordan-wright/email"
)

type EmailHandlerDeps struct {
	Config *configs.Config
}
type EmailHandler struct {
	Config *configs.Config
}

func NewEmailHandler(router *http.ServeMux, deps EmailHandlerDeps) {
	handler := &EmailHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())

}

func (h *EmailHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := req.HandleBody[SendEmailRequest](&w, r)
		if err != nil {
			return
		}

		e := &email.Email{
			To:      []string{body.To},
			From:    fmt.Sprintf("My App <%s>", h.Config.Email),
			Subject: "Awesome Subject",
			Text:    []byte("Text Body is, of course, supported!"),
			HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
			Headers: textproto.MIMEHeader{},
		}

		smtpAddr := fmt.Sprintf("%s:%s", h.Config.Address, h.Config.Port)
		auth := smtp.PlainAuth("", h.Config.Email, h.Config.Password, h.Config.Address)

		go func() {
			log.Printf("Starting to send email to %s in background...", e.To[0])
			err := e.Send(smtpAddr, auth)
			if err != nil {
				log.Printf("Failed to 	send email in background: %v", err)
			} else {
				log.Printf("Email successfully sent in background to %s", e.To[0])
			}
		}()

		res.JSON(w, map[string]string{"info": "Email request accepted and is being processed."}, http.StatusAccepted)
	}
}

func (h *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		log.Printf("Verifying hash: %s", hash)
		fmt.Fprintf(w, "Hash %s verified", hash)
	}
}
