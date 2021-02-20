package login

import (
	"login_system/unitycom"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	backcom := unitycom.GetSendClientComDefault()

	c.JSON(http.StatusOK, backcom)
}
