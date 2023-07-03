package main

import "net/http"

func (app *Config) sendMail(w http.ResponseWriter, r *http.Request) {
	type mailMessage struct {
		FromAddress string `json:"from_address"`
		ToAddress   string `json:"to_address"`
		Subject     string `json:"subject"`
		Message     string `json:"message"`
	}

	var requestPayload mailMessage

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	message := &Message{
		FromAddress: requestPayload.FromAddress,
		ToAddress:   requestPayload.ToAddress,
		Subject:     requestPayload.Subject,
		Data:        requestPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(message)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	payload := jsonResponse{
		Error:   false,
		Message: "Sent to " + requestPayload.ToAddress,
	}

	app.writeJSON(w, http.StatusAccepted, payload)
}
