package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mangohow/easygin"
	"github.com/mangohow/gin-template/internal/service"
	"github.com/mangohow/gin-template/pkg/logger"
	"github.com/sirupsen/logrus"
)

type HelloController struct {
	service *service.HelloService
	logger  *logrus.Logger
}

func NewHelloController() *HelloController {
	return &HelloController{
		service: &service.HelloService{},
		logger:  logger.Logger(),
	}
}

// Hello handler
// @get /api/hello?id=xx&username=xx
func (c *HelloController) Hello(ctx *gin.Context, id int, username string) *easygin.Response {
	c.logger.Debugf("id:%d, username: %s", id, username)
	ok, err := c.service.Verify(id, username)
	if err != nil {
		c.logger.Error("error:", err)
		if easygin.IsRespError(err) {
			return easygin.Fail(easygin.AsRespError(err))
		}

		return easygin.Fail(easygin.NewFromError(err))
	}

	c.logger.Info("hello controller")
	return easygin.OkData(ok)
}
