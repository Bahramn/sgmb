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
	Total   int                 `json:"total"`
	Mobiles []TransformedClient `json:"mobiles"`
	Devices []TransformedClient `json:"devices"`
}

type TransformedClient struct {
	TYPE          string `json:"type"`
	Id            string `json:"id"`
	ConnType      string `json:"connectionType"`
	LastCheckedAt string `json:"lastCheckedAt"`
}

func (api *httpApi) getClients(w http.ResponseWriter, r *http.Request) {
	cRes := clientsRes{
		Total:   api.s.NumberOfClients(),
		Mobiles: transformClients(api.s.GetMobileClients()),
		Devices: transformClients(api.s.GetDeviceClients()),
	}

	cResJson, err := json.Marshal(cRes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cResJson)
}

func transformClients(clients []*Client) []TransformedClient {
	tClients := make([]TransformedClient, 0)
	for _, c := range clients {
		tClients = append(tClients, TransformedClient{
			TYPE:          c.originType,
			Id:            c.id,
			ConnType:      c.connType,
			LastCheckedAt: c.lastCheckedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return tClients
}
