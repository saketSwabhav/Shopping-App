package user

import uuid "github.com/satori/go.uuid"

type User struct {
	ID    uuid.UUID `json:"id"`
	Token string    `json:"token"`
}
