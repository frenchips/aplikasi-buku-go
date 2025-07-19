package repository

import (
	"aplikasi-buku-go/model"
	"database/sql"
	"time"
)

type BukuRepository interface {
	InsertBuku(c model.Buku) (model.Buku, error)
	GetBooks(c model.Buku) (model.Buku, error)
	DeleteBuku(c model.Buku) (model.Buku, error)
	GetAllBooks(c model.Buku) (result []model.Buku, err error)
}

type bukuRepository struct {
	db *sql.DB
}

func NewBukuRepo(db *sql.DB) BukuRepository {
	return &bukuRepository{db: db}
}

func (c *bukuRepository) InsertBuku(buku model.Buku) (model.Buku, error) {

	now := time.Now()
	createBy := "Admin"

	sql := "INSERT INTO Buku(title, category_id, description, image_url, release_year, price, total_page, thickness,  created_at, created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	errs := c.db.QueryRow(sql, &buku.Title, &buku.Category.Id, &buku.Description, &buku.ImageUrl, &buku.ReleaseYear, &buku.Price, &buku.TotalPage, &buku.Thickness, &now, &createBy).Scan(&buku.Id)
	if errs != nil {
		panic(errs)
	}

	buku.CreatedAt = now
	buku.CreatedBy = createBy

	var categoryName string
	errs = c.db.QueryRow("SELECT name FROM category WHERE id = $1", buku.Category.Id).Scan(&categoryName)
	if errs != nil {
		return buku, errs
	}

	buku.Category.Name = categoryName

	return buku, nil
}

func (c *bukuRepository) GetBooks(buku model.Buku) (model.Buku, error) {
	sql := "SELECT b.title, c.id as category, c.name, b.description, b.image_url, b.release_year, b.price, b.total_page, b.thickness FROM buku b JOIN category c on b.category_id = c.id WHERE b.id = $1"
	errs := c.db.QueryRow(sql, buku.Id).Scan(&buku.Title, &buku.Category.Id, &buku.Category.Name, &buku.Description, &buku.ImageUrl, &buku.ReleaseYear, &buku.Price, &buku.TotalPage, &buku.Thickness)
	if errs != nil {
		panic(errs)
	}

	return buku, nil

}

func (c *bukuRepository) GetAllBooks(buku model.Buku) (result []model.Buku, err error) {
	sql := "SELECT b.id, b.title, c.id as category, c.name, b.description, b.image_url, b.release_year, b.price, b.total_page, b.thickness FROM buku b JOIN category c on b.category_id = c.id ORDER BY b.id ASC"
	rows, err := c.db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var buku model.Buku

		err = rows.Scan(&buku.Id, &buku.Title, &buku.Category.Id, &buku.Category.Name, &buku.Description, &buku.ImageUrl, &buku.ReleaseYear, &buku.Price, &buku.TotalPage, &buku.Thickness)
		if err != nil {
			return
		}

		result = append(result, buku)
	}
	return
}

func (c *bukuRepository) DeleteBuku(buku model.Buku) (model.Buku, error) {

	sql := "DELETE FROM buku WHERE id = $1 RETURNING id"
	errs := c.db.QueryRow(sql, &buku.Id).Scan(&buku.Id)
	if errs != nil {
		panic(errs)
	}
	return buku, nil
}
