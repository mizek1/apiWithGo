package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/mizek1/apiWithGo/entity"
	"github.com/mizek1/apiWithGo/service"
)

type FilmeController interface {
	SalvarFilme(ctx *gin.Context) entity.Filme
	ListarFilmes() []entity.Filme
}

type controller struct {
	servico service.ServicoFilme
}

func Novo(servico service.ServicoFilme) FilmeController {
	return &controller{
		servico: servico,
	}
}

func (c *controller) SalvarFilme(ctx *gin.Context) entity.Filme {
	var filme entity.Filme
	ctx.BindJSON(&filme)
	c.servico.SalvarFilme(filme)
	return filme
}

func (c *controller) ListarFilmes() []entity.Filme {
	return c.servico.ListarFilmes()
}
