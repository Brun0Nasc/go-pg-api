package schemas

type CreateDiretor struct {
	Nome string `json:"nome" binding:"required"`
	Sexo string `json:"sexo" binding:"required"`
}

type UpdateDiretor struct {
	Nome string `json:"nome" binding:"required"`
	Sexo string `json:"sexo" binding:"required"`
}