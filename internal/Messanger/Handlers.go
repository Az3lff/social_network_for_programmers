package Messanger

import "github.com/gin-gonic/gin"

func MessangerApp() {
	api := gin.RouterGroup{}
	api.Handlers = []gin.HandlerFunc{getMessangerPage}
}

func getMessangerPage(c *gin.Context) {

}
