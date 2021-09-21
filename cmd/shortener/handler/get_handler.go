package handler

import (
	"net/http"

	"GoPracticum/cmd/shortener/repository"

	"github.com/gorilla/mux"
)

func GetHandler(s *repository.Service) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			return
		}
		vars := mux.Vars(r)
		link, err := s.Repository.Get(vars["key"])
		if err != nil {
			http.Error(w, "Cannot find link", http.StatusInternalServerError)
		}

		w.Header().Set("Location", link)
		w.WriteHeader(http.StatusTemporaryRedirect)
	}
}
