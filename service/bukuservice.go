package service

import (
	"aplikasi-buku-go/model"
	"aplikasi-buku-go/repository"
)

type BukuService interface {
	InsertBuku(model.Buku) (model.Buku, error)
	GetBooks(model.Buku) (model.Buku, error)
	DeleteBuku(model.Buku) (model.Buku, error)
	GetAllBooks(model.Buku) (result []model.Buku, err error)
}

type bukuService struct {
	repo repository.BukuRepository
}

func NewBukuService(repo repository.BukuRepository) BukuService {
	return &bukuService{repo: repo}
}

func (c *bukuService) InsertBuku(books model.Buku) (model.Buku, error) {
	return c.repo.InsertBuku(books)
}

func (c *bukuService) GetBooks(books model.Buku) (model.Buku, error) {
	return c.repo.GetBooks(books)
}

func (c *bukuService) GetAllBooks(books model.Buku) (result []model.Buku, err error) {
	return c.repo.GetAllBooks(books)
}

func (c *bukuService) DeleteBuku(books model.Buku) (model.Buku, error) {
	return c.repo.DeleteBuku(books)
}
