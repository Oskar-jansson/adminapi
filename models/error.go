package models

// Response on defined errors from api
type RejectItem struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
	Code  int    `json:"code,omitempty"`
}

type ErrorMessage struct {
	Errorcode        int           `json:"errorcode,omitempty"`
	Friendlymessage  string        `json:"friendlymessage,omitempty"`
	Developermessage string        `json:"developermessage,omitempty"`
	Moreinfo         string        `json:"moreinfo,omitempty"`
	Reject           *[]RejectItem `json:"reject,omitempty"`
}
