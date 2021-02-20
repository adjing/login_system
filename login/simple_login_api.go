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

var m_count = 0

//add +1
func GetCount() int {
	m_count = m_count + 1
	return m_count
}
