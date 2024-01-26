package broker_controller

import "github.com/gin-gonic/gin"

type controller struct {
}

func BrokerController(router *gin.RouterGroup) {
	c := &controller{}

	router.POST("/", c.Broker)
	router.POST("/handle", c.HandleSubmission)
}
