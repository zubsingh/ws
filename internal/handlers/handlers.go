package handlers

import (
	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode())

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Home(rw http.ResponseWriter, r *http.Request) {
	err := renderPage(rw, "home.jet.html", nil)
	if err != nil {
		log.Println(err)
		return
	}
}

//	WsJsonResponse defines the response send back from websocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"Message"`
	MessageType string `json:"MessageType"`
}

// WsEndPoint upgrades connection to websocket
func WsEndPoint(rw http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(rw, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Print("Client connected to input")

	var response WsJsonResponse
	response.Message = `<em><small>Connected to server</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
		return
	}
}

func renderPage(rw http.ResponseWriter, tmpl string, data jet.VarMap) error {
	views, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = views.Execute(rw, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
