package entity

type Filme struct {
	Titulo  string `json:"titulo"`
	Genero  string `json:"genero"`
	Duracao int64  `json:"duracao"`
	// Atores  *[]Ator
}

// type Ator struct {
// 	Nome  string `json: "nome"`
// 	Idade int64  `json:"idade"`
// }
