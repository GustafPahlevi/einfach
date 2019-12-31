package messaged

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func (g *Message) Get(w http.ResponseWriter, r *http.Request) {
	log.Info("Get Message")
}
