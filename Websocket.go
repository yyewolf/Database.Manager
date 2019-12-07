package main

import (
	"net/http"
)

// Structs of all the request : how they'll look

type Send_Request struct {
	Action string `json:"action"`
	DB     string `json:"db"`
}

type Column_Type struct {
	ID               int    `json:"id"`
	Eleve            string `json:"eleve"`
	Responsable      string `json:"responsable"`
	Telephone        string `json:"tel"`
	Ecole            string `json:"ecole"`
	Instrument       string `json:"instrument"`
	Niveauinstrument string `json:"niveauinstrument"`
	Solfege          string `json:"solfege"`
	Niveausolfege    string `json:"niveausolfege"`
	Location         string `json:"location"`
	Harmonie         string `json:"harmonie"`
	Email            string `json:"mail"`
	Adresse          string `json:"adresse"`
}

type Receive_Request struct {
	//When User send a message
	Action string `json:"action"`

	//When a user sends a request to DB
	Request string `json:"req"`

	//When a user deletes an entry from the DB
	ID string `json:"id"`

	//When a user adds to DB
	Add Column_Type `json:"data"`

	//When a user edits a DB Column
	Column  string `json:"column"`
	NewData string `json:"newdata"`
}

func Websocket_Broadcast(msg string) {
	for i := range Connections {
		err := Connections[i].WriteMessage(1, []byte(msg))
		if err != nil {
			Connections = append(Connections[:i], Connections[i+1:]...)
			//Removes the connection if it's unable to send a message.
		}
	}
}

func Websocket_Connection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity
	if err != nil {
		return
	}
	conn.SetCloseHandler(func(code int, text string) error {
		for i := range Connections {
			if Connections[i] == conn {
				Connections = append(Connections[:i], Connections[i+1:]...)
			}
		}
		return nil
	})
	Connections = append(Connections, conn)

	for {
		// Read message from browser
		rec := Receive_Request{}
		err := conn.ReadJSON(&rec)
		if err != nil {
			return
		}
		if _, ok := Websocket_Receive_Functions[rec.Action]; !ok {
			return
		}
		Websocket_Receive_Functions[rec.Action](rec)
	}
}
