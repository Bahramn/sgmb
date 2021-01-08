package app

import (
	"encoding/json"
	"net/http"
)

type httpApi struct {
	s *Server
}

func newHttpApi(s *Server) *httpApi {
	return &httpApi{
		s,
	}
}

func (api *httpApi) runApi(s *Server, addr string) error {
	http.HandleFunc("/", api.appIndex)
	http.HandleFunc("/clients", api.getClients)

	return http.ListenAndServe(addr, nil)
}

func (api *httpApi) appIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Simple Golang Message Broker v1"))
}

/*
	@TODO: Add devices and mobiles info in response
*/
type clientsRes struct {
	Total int
}

func (api *httpApi) getClients(w http.ResponseWriter, r *http.Request) {
	cRes := clientsRes{
		Total: api.s.NumberOfClients(),
	}

	cResJson, err := json.Marshal(cRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cResJson)
}
