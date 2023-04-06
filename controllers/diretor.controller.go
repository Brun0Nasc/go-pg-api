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

	diretor, err := dc.db.CreateDiretor(ctx, *args)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "diretor": diretor})
}

func (dc *DiretorController) UpdateDiretor(ctx *gin.Context) {
	var payload *schemas.UpdateDiretor
	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	args := &db.UpdateDiretorParams{
		ID: uuid.MustParse(id),
		Nome: sql.NullString{String: payload.Nome, Valid: payload.Nome != ""},
		Sexo: db.NullSexgen{Sexgen: db.Sexgen(payload.Sexo), Valid: payload.Sexo != ""},
	}

	diretor, err := dc.db.UpdateDiretor(ctx, *args)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No director with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "diretor": diretor})
}

func (dc *DiretorController) GetDiretorById(ctx *gin.Context) {
	id := ctx.Param("id")

	diretor, err := dc.db.GetDiretorById(ctx, uuid.MustParse(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No director with taht ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "diretor": diretor})
}

func (dc *DiretorController) ListDiretores(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	offset := (intPage - 1) * intLimit

	args := &db.ListDiretoresParams{
		Limit: int32(intLimit),
		Offset: int32(offset),
	}

	diretores, err := dc.db.ListDiretores(ctx, *args)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if diretores == nil {
		diretores = []db.Diretor{}
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(diretores), "data": diretores})
}

func (dc *DiretorController) DeleteDiretor(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := dc.db.GetDiretorById(ctx, uuid.MustParse(id))
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No director with that ID exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	err = dc.db.DeleteDiretor(ctx, uuid.MustParse(id))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}