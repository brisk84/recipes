package tools

import uuid "github.com/satori/go.uuid"

func GetGuid() string {
	u1 := uuid.Must(uuid.NewV1(), nil)
	return u1.String()
}
