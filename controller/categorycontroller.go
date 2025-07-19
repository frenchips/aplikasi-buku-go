package controller

import (
	"aplikasi-buku-go/common"
	"aplikasi-buku-go/config"
	"aplikasi-buku-go/model"
	"aplikasi-buku-go/repository"
	"aplikasi-buku-go/service"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InsertCategory(ctx *gin.Context) {
	var input model.Category

	// Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	// Buat repo dan service
	repo := repository.NewCategoryRepo(config.DB)
	service := service.NewCategoryService(repo)

	categori, err := service.InsertCategory(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully saved category", categori)
}

func UpdateCategory(ctx *gin.Context) {
	var input model.Category

	// // Bind JSON request body ke struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		common.GenerateErrorResponse(ctx, "Invalid input")
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))

	input.Id = id

	// Buat repo dan service
	repo := repository.NewCategoryRepo(config.DB)
	service := service.NewCategoryService(repo)

	fmt.Println("Input JSON Name : ", input.Name)

	categori, err := service.UpdateCategory(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully saved category", categori)
}

func DeleteCategory(ctx *gin.Context) {
	var input model.Category

	id, _ := strconv.Atoi(ctx.Param("id"))

	input.Id = id

	// Buat repo dan service
	repo := repository.NewCategoryRepo(config.DB)
	service := service.NewCategoryService(repo)

	fmt.Println("Input JSON Name : ", input.Name)

	categori, err := service.DeleteCategory(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully delete category", categori)
}

func GetAllCategory(ctx *gin.Context) {
	var input model.Category

	// Buat repo dan service
	repo := repository.NewCategoryRepo(config.DB)
	service := service.NewCategoryService(repo)

	fmt.Println("Input JSON Name : ", input.Name)

	categori, err := service.GetAllCategory(input)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully get all category", categori)
}

func GetBooksByCategory(ctx *gin.Context) {

	id, _ := strconv.Atoi(ctx.Param("id"))

	// Buat repo dan service
	repo := repository.NewCategoryRepo(config.DB)
	service := service.NewCategoryService(repo)

	categori, err := service.GetBooksByCategory(id)
	if err != nil {
		common.GenerateErrorResponse(ctx, err.Error())
		return
	}

	common.GenerateSuccessResponseWithData(ctx, "Successfully get books by category", categori)
}
