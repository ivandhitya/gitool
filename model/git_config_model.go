package model

import (
	resty "github.com/go-resty/resty/v2"
)

type GitConfig struct {
	URL    string
	Token  string
	Client *resty.Client
}
