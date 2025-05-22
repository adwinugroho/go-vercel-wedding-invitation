package repository

import (
	"database/sql"
	"log"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/model"
	"github.com/google/uuid"
)

type (
	WishesInterface interface {
		Insert(data model.Wishes) (string, error)
		List(offset, limit int) ([]model.Wishes, error)
	}
)

type wishesImp struct {
	DB *sql.DB
}

func NewWishesRepository(conn *sql.DB) WishesInterface {
	return &wishesImp{
		DB: conn,
	}
}

func (c *wishesImp) Insert(data model.Wishes) (string, error) {
	data.ID = uuid.New().String()
	// defer c.DB.Close()
	query := `INSERT INTO tb_wishes (id, name, message, is_published, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := c.DB.Exec(query, data.Name, data.Message, data.IsPublished, data.CreatedAt)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return "", err
	}

	log.Printf("Insert new wishes to DB with ID:%s\n", data.ID)
	return data.ID, nil
}

func (c *wishesImp) List(offset, limit int) ([]model.Wishes, error) {
	var results []model.Wishes
	query := `SELECT id, name, message FROM tb_wishes WHERE is_published = true LIMIT $1 OFFSET $2`
	rows, err := c.DB.Query(query, limit, offset)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.Wishes

		err = rows.Scan(&result.ID, &result.Name, &result.Message)
		if err != nil {
			log.Printf("Error while scan:%+v\n", err)
		}

		results = append(results, result)
	}

	return results, nil
}
