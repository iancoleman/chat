package chat

import (
	"harken/base/groups"
	"harken/base/http"
)

var rooms = map[string]string{}

func Subscribe(session string, data string) *http.OutgoingMsg {
	c := http.Connections[session]
	groups.Subscribe(c, data)
	rooms[session] = data
	response := http.OutgoingMsg{Ext:"chat",Method:"Subscribe",Data:data}
	return &response
}

func Send(session string, data string) *http.OutgoingMsg {
	msg := http.OutgoingMsg{Ext: "chat", Method: "Send", Data: data}
	room := rooms[session]
	groups.Broadcast(room, msg)
	return nil
}
