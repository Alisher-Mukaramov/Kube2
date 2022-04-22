package models

import (
	"encoding/json"
)

type Response struct {
	Status string `json:"status,omitempty"`
}

func (r Response) Marshal() (bt []byte, err error) {
	if bt, err = json.Marshal(r); err != nil {
		return
	}
	return
}
