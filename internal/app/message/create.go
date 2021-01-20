package message

import (
	"encoding/json"
	"go-simple-svc/constant"
	"go-simple-svc/model"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Create is main handler for create action
func (m *Message) Create(w http.ResponseWriter, r *http.Request) {
	var request model.Model
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warnf("cannot read request, got: %v", err)
		m.handlerCreateResp(w, r, constant.Failed)
		return
	}
	defer func() {
		err = r.Body.Close()
		if err != nil {
			log.Warnf("cannot close request body, got: %v", err)
		}
	}()

	err = json.Unmarshal(reqBody, &request)
	if err != nil {
		log.Warnf("cannot unmarshal request, got: %v", err)
		m.handlerCreateResp(w, r, constant.Failed)
		return
	}

	logFields := log.Fields{
		"request": request,
	}
	log.WithFields(logFields).Info("sending message...")

	err = m.collection.Insert(request)
	if err != nil {
		log.WithFields(logFields).Errorf("error while insert document, got: %v", err)
		m.handlerCreateResp(w, r, constant.Failed)
		return
	}

	log.WithFields(logFields).Info("message sent")
	m.handlerCreateResp(w, r, constant.Success)
	return
}

func (m *Message) handlerCreateResp(w http.ResponseWriter, r *http.Request, status string) {
	var response ResponseStatus
	switch status {
	case constant.Success:
		response.Status = "message sent"
	case constant.Failed:
		response.Status = "cannot sending message"
	default:
		response.Status = "status unknown"
	}
	resByte, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resByte)
	if err != nil {
		log.Warnf("cannot write http response, got: %v", err)
	}

	return
}
