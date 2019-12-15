package main

import "encoding/json"

func SendDB(HTML string) {
	//Send the DB over the websocket for the browser to update
	Message, err := json.Marshal(Send_Request{
		Action: "updateDB",
		DB:     HTML,
	})
	if err != nil {
		return
	}
	Websocket_Broadcast(string(Message))
}

func SendSmallDB(HTML string) {
	//Send the DB over the websocket for the browser to update
	Message, err := json.Marshal(Send_Request{
		Action: "smallupdateDB",
		DB:     HTML,
	})
	if err != nil {
		return
	}
	Websocket_Broadcast(string(Message))
}
