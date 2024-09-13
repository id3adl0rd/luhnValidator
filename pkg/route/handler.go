package route

import (
	"cardValidator/pkg/responder"
	"cardValidator/pkg/validator"
	"encoding/json"
	"log"
	"net/http"
)

const (
	invalidPayload = "Invalid Payload"
)

type cardPayload struct {
	CardNumber int `json:"cardNumber,string,omitempty"`
}

type cardResponder struct {
	Value int `json:"cardNumber,omitempty"`
}

func cardHandler(w http.ResponseWriter, r *http.Request) {
	var cPayload cardPayload
	err := json.NewDecoder(r.Body).Decode(&cPayload)
	if err != nil {
		log.Println(err)
		responder.RespondWithError(w, http.StatusBadRequest, invalidPayload)
		return
	}

	defer r.Body.Close()

	luhn := validator.CalculateLuhn(cPayload.CardNumber)

	responder.RespondWithJSON(w, http.StatusOK, &cardResponder{
		Value: luhn,
	})
}
