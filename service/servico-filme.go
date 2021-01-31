package service

import "github.com/mizek1/apiWithGo/entity"

type ServicoFilme interface {
	SalvarFilme(entity.Filme) entity.Filme
	ListarFilmes() []entity.Filme
}

type servicoFilme struct {
	filmes []entity.Filme
}

func NovoFilme() *servicoFilme {
	return &servicoFilme{}
}

func (s *servicoFilme) SalvarFilme(filme entity.Filme) entity.Filme {
	s.filmes = append(s.filmes, filme)
	return filme
}

func (s *servicoFilme) ListarFilmes() []entity.Filme {
	return s.filmes
}
