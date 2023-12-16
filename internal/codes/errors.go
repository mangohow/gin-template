package codes

import "github.com/mangohow/easygin"

type RespErrorImpl = easygin.RespErrorImpl

var (
	ErrUserInvalid = &RespErrorImpl{
		Codee:    1,
		Messagee: "user invalid",
	}
)
