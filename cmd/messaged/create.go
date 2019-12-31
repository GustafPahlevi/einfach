package messaged

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (g *Message) Create(w http.ResponseWriter, r *http.Request) {
	log.Info("Create Message")
}
