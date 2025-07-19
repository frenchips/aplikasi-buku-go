package repository

import (
	"aplikasi-buku-go/model"
	"aplikasi-buku-go/response"
	"database/sql"
	"fmt"
	"time"
)

type CategoryRepository interface {
	InsertCategory(c model.Category) (model.Category, error)
	UpdateCategory(c model.Category) (model.Category, error)
	DeleteCategory(c model.Category) (model.Category, error)
	GetAllCategory(c model.Category) (result []model.Category, err error)
	GetBooksByCategory(categoryId int) (response.CategoryResponse, error)
}

type categoryRepository struct {
	db *sql.DB
}

func NewCategoryRepo(db *sql.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (c *categoryRepository) InsertCategory(category model.Category) (model.Category, error) {

	now := time.Now()
	createBy := "Admin"

	sql := "INSERT INTO category(name, created_at, created_by) VALUES ($1, $2, $3) RETURNING id"
	errs := c.db.QueryRow(sql, &category.Name, &now, &createBy).Scan(&category.Id)
	if errs != nil {
		panic(errs)
	}
	category.CreatedAt = now
	category.CreatedBy = createBy
	return category, nil
}

func (c *categoryRepository) UpdateCategory(category model.Category) (model.Category, error) {

	now := time.Now()
	updateBy := "Admin"

	sql := "UPDATE category SET name = $1, modified_at = $2, modified_by = $3 WHERE id = $4 RETURNING id"
	errs := c.db.QueryRow(sql, category.Name, now, updateBy, category.Id).Scan(&category.Id)
	if errs != nil {
		panic(errs)
	}
	fmt.Println("Category name before update:", category.Name)
	category.ModifiedAt = now
	category.ModifiedBy = updateBy
	return category, nil
}

func (c *categoryRepository) DeleteCategory(category model.Category) (model.Category, error) {

	sql := "DELETE FROM category WHERE id = $1 RETURNING id"
	errs := c.db.QueryRow(sql, &category.Id).Scan(&category.Id)
	if errs != nil {
		panic(errs)
	}
	return category, nil
}

func (c *categoryRepository) GetAllCategory(category model.Category) (result []model.Category, err error) {
	sql := "SELECT id, name FROM category ORDER BY id ASC"
	rows, err := c.db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var categori model.Category

		err = rows.Scan(&categori.Id, &categori.Name)
		if err != nil {
			return
		}

		result = append(result, categori)
	}

	return
}

func (c *categoryRepository) GetBooksCategory(category model.Category) (result []model.Category, err error) {
	sql := "SELECT id, name FROM category ORDER BY id ASC"
	rows, err := c.db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var categori model.Category

		err = rows.Scan(&categori.Id, &categori.Name)
		if err != nil {
			return
		}

		result = append(result, categori)
	}

	return
}

func (c *categoryRepository) GetBooksByCategory(categoryId int) (response.CategoryResponse, error) {
	sql := "select distinct c.name, b.title, b.description, b.release_year, b.price, b.total_page, b.thickness from category c join buku b on c.id = b.category_id where c.id = $1"
	rows, err := c.db.Query(sql, categoryId)

	defer rows.Close()

	var name string
	var books []response.BukuResponse

	for rows.Next() {

		var buku response.BukuResponse

		err = rows.Scan(&name, &buku.Title, &buku.Description, &buku.ReleaseYear, &buku.Price, &buku.TotalPage, &buku.Thickness)
		if err != nil {
			return response.CategoryResponse{}, err
		}

		books = append(books, buku)
	}

	// Bangun response akhir
	result := response.CategoryResponse{
		Name: name,
		Buku: books,
	}

	return result, nil
}
