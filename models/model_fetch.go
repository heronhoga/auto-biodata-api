package models

type ApiResult struct {
	Service string
	Body []byte
	Err error
}