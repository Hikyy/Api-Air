package handlers

import (
	"net/http"
	"os"
)

type CustomRouter struct {
	Router http.Handler
	C      chan os.Signal
}

func (cr CustomRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cr.Router.ServeHTTP(w, r)
	select {
	case <-cr.C:
		// Signal d'interruption reçu, effectuez les actions de nettoyage et de sortie appropriées
		// ...
	default:
		// Aucun signal reçu, continuez à traiter les requêtes normalement
	}
}
