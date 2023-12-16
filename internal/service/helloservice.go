package service

import (
	"errors"

	"github.com/mangohow/gin-template/internal/codes"
)

type HelloService struct {
}

func (s HelloService) Verify(id int, username string) (bool, error) {
	if id != 1 {
		return false, errors.New("interal error")
	}

	if username != "root" {
		return false, codes.ErrUserInvalid
	}

	return true, nil
}
