package controller

import (
	"aplikasi-buku-go/common"
	"aplikasi-buku-go/config"
	"aplikasi-buku-go/model"
	"aplikasi-buku-go/repository"
	"aplikasi-buku-go/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertBuku(ctx *gin.Context) {
	var input model.Buku

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	// Buat repo dan service
	repo := repository.NewBukuRepo(config.DB)
	service := service.NewBukuService(repo)

	categori, err := service.InsertBuku(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully saved books", categori)
}

func GetBuku(ctx *gin.Context) {
	var input model.Buku

	id, _ := strconv.Atoi(ctx.Param("id"))

	input.Id = id

	// Buat repo dan service
	repo := repository.NewBukuRepo(config.DB)
	service := service.NewBukuService(repo)

	categori, err := service.GetBooks(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully saved books", categori)
}

func DeleteBuku(ctx *gin.Context) {
	var input model.Buku

	id, _ := strconv.Atoi(ctx.Param("id"))

	input.Id = id

	// Buat repo dan service
	repo := repository.NewBukuRepo(config.DB)
	service := service.NewBukuService(repo)

	_, err := service.DeleteBuku(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully delete books", nil)
}

func GetAllBooks(ctx *gin.Context) {
	var input model.Buku

	// Buat repo dan service
	repo := repository.NewBukuRepo(config.DB)
	service := service.NewBukuService(repo)

	categori, err := service.GetAllBooks(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully get all books", categori)
}
