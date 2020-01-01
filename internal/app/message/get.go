package message

import (
	"einfach-msg/model"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (m *Message) Get(w http.ResponseWriter, r *http.Request) {
	var response []model.Model
	log.Info("get request for retrieving all message")

	messages, err := m.collection.Get()
	if err != nil {
		log.Errorf("error while get message, got: %v ", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(`{"status":"oops, cannot get message"}`))
		if err != nil {
			log.Warnf("cannot write http response, got: %v", err)
		}

		return
	}

	for _, message := range messages {
		response = append(response, model.Model{
			ID:         message.ID,
			Subject:    message.Subject,
			Message:    message.Message,
			ReceiverID: message.ReceiverID,
			SenderID:   message.SenderID,
			Status:     message.Status,
			Time:       message.Time,
		})
	}

	logFields := log.Fields{
		"messages": messages,
	}
	log.WithFields(logFields).Info("successfully get all message")

	resByte, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resByte)
	if err != nil {
		log.WithFields(logFields).Warnf("cannot write http response, got: %v", err)
	}

	return
}
