package ksclient

import (
	"github.com/go-resty/resty/v2"
)

type KApisClient struct {
	RestClient   resty.Client
	Token        string
	Expire       int64
	RefreshToken string
}
