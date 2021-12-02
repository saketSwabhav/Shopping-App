package general

import uuid "github.com/satori/go.uuid"

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
	TokenUUID    uuid.UUID
	RefreshUUID  uuid.UUID
}
