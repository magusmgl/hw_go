package valid

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"net/textproto"
	"validation/api/configs"

	"github.com/jordan-wright/email"
)

type SendEmailRequest struct {
	To string `json:"to"`
}

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
		var req SendEmailRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if req.To == "" {
			http.Error(w, "Recipient email ('to') is required", http.StatusBadRequest)
			return
		}

		e := &email.Email{
			To:      []string{req.To},
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

		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Email request accepted and is being processed.")
	}
}

func (h *EmailHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		log.Printf("Verifying hash: %s", hash)
		fmt.Fprintf(w, "Hash %s verified", hash)
	}
}
