package controllers

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/Brun0Nasc/go-pg-api/db/sqlc"
	"github.com/Brun0Nasc/go-pg-api/schemas"
)

type DiretorController struct {
	db *db.Queries
	ctx context.Context
}

func NewDiretorController(db *db.Queries, ctx context.Context) *DiretorController {
	return &DiretorController{db, ctx}
}

func (dc *DiretorController) CreateDiretor(ctx *gin.Context) {
	var payload *schemas.CreateDiretor

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.CreateDiretorParams{
		Nome: payload.Nome,
		Sexo: db.Sexgen(payload.Sexo),
	}
}