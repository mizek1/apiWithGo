package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mizek1/apiWithGo/controller"
	"github.com/mizek1/apiWithGo/service"
)

var (
	servicoFilme    service.ServicoFilme       = service.NovoFilme()
	controllerFilme controller.FilmeController = controller.Novo(servicoFilme)
)

func main() {
	server := gin.Default()

	server.GET("/filmes", func(c *gin.Context) {
		c.JSON(200,
			controllerFilme.ListarFilmes())
	})

	server.POST("/filme", func(c *gin.Context) {
		c.JSON(201,
			controllerFilme.SalvarFilme(c))
	})

	server.Run(":8080")
}
