package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// se implementa este tipo de funcion por estar usando gin

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
