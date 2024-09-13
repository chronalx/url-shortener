package handlers

import (
	"github.com/chronalx/url-shortener/internal/app/storage"
	"github.com/chronalx/url-shortener/internal/app/tools"
	"io"
	"net/http"
)

func HandleRequests(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		shortenURL(w, r)
	} else if r.Method == http.MethodGet {
		redirectURL(w, r)
	} else {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
	}
}

func shortenURL(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
	}
	defer r.Body.Close()

	for {
		newUrl := tools.GenUrl()
		if _, ok := storage.MapURLs[newUrl]; !ok {
			storage.MapURLs[newUrl] = string(body)
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(r.Host + "/" + newUrl))
			break
		}
	}
}

func redirectURL(w http.ResponseWriter, r *http.Request) {
	originUrl, ok := storage.MapURLs[r.URL.Path[1:]]
	if !ok {
		http.Error(w, "Неверный запрос", http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", originUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)
	w.Write([]byte(originUrl))
	//http.Redirect(w, r, originUrl, http.StatusTemporaryRedirect)

}
