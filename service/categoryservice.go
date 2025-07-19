package service

import (
	"aplikasi-buku-go/model"
	"aplikasi-buku-go/repository"
	"aplikasi-buku-go/response"
)

type CategoryService interface {
	InsertCategory(model.Category) (model.Category, error)
	UpdateCategory(model.Category) (model.Category, error)
	DeleteCategory(model.Category) (model.Category, error)
	GetAllCategory(model.Category) (result []model.Category, err error)
	GetBooksByCategory(categoryId int) (response.CategoryResponse, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

func (c *categoryService) InsertCategory(category model.Category) (model.Category, error) {
	return c.repo.InsertCategory(category)
}

func (c *categoryService) UpdateCategory(category model.Category) (model.Category, error) {
	return c.repo.UpdateCategory(category)
}

func (c *categoryService) DeleteCategory(category model.Category) (model.Category, error) {
	return c.repo.DeleteCategory(category)
}

func (c *categoryService) GetAllCategory(category model.Category) (result []model.Category, err error) {
	return c.repo.GetAllCategory(category)
}

func (c *categoryService) GetBooksByCategory(categoryId int) (response.CategoryResponse, error) {
	return c.repo.GetBooksByCategory(categoryId)
}
