package routes

import (
	"github.com/Brun0Nasc/go-pg-api/controllers"
	"github.com/gin-gonic/gin"
)

type DiretorRoutes struct {
	diretorController controllers.DiretorController
}

func NewRouteDiretor(diretorController controllers.DiretorController) DiretorRoutes {
	return DiretorRoutes{diretorController}
}

func (dc *DiretorRoutes) DiretorRoute(rg *gin.RouterGroup) {

	router := rg.Group("diretores")
	router.POST("/", dc.diretorController.CreateDiretor)
	router.GET("/", dc.diretorController.ListDiretores)
	router.PATCH("/:id", dc.diretorController.UpdateDiretor)
	router.GET("/:id", dc.diretorController.GetDiretorById)
	router.DELETE("/:id", dc.diretorController.DeleteDiretor)
	
}