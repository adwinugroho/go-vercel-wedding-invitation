package repository

import (
	"context"
	"log"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/model"
	"github.com/jackc/pgx/v5"
)

type (
	GiftInterface interface {
		Insert(ctx context.Context, data model.Gifts) (int32, error)
		List(ctx context.Context, offset, limit int) ([]model.Gifts, error)
		UpdateIsBooked(ctx context.Context, data model.Gifts) error
	}
)

type giftImp struct {
	DB *pgx.Conn
}

func NewGiftRepository(conn *pgx.Conn) GiftInterface {
	return &giftImp{
		DB: conn,
	}
}

func (c *giftImp) Insert(ctx context.Context, data model.Gifts) (int32, error) {
	query := `INSERT INTO tb_gifts(name, description, price, image_url, category, store_link, is_booked, created_at) 
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	_, err := c.DB.Exec(ctx, query,
		data.Name,
		data.Description,
		data.Price,
		data.ImageURL,
		data.Category,
		data.StoreLink,
		data.IsBooked,
		data.CreatedAt,
	)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return 0, err
	}

	defer c.DB.Close(ctx)

	log.Printf("Insert new gift to DB with ID:%d\n", data.ID)
	return data.ID, nil
}

func (c *giftImp) List(ctx context.Context, offset, limit int) ([]model.Gifts, error) {
	var results []model.Gifts

	// Safety: ensure range is valid
	if limit <= 0 {
		limit = 10 // default limit
	}
	if offset <= 0 {
		offset = 1
	}

	offsetPage := (offset - 1) * limit

	query := `SELECT * 
	FROM tb_gifts
	ORDER BY created_at DESC 
	LIMIT $1 OFFSET $2`

	defer c.DB.Close(ctx)
	rows, err := c.DB.Query(ctx, query, limit, offsetPage)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.Gifts

		err = rows.Scan(
			&result.ID,
			&result.Name,
			&result.Description,
			&result.Price,
			&result.ImageURL,
			&result.Category,
			&result.StoreLink,
			&result.IsBooked,
			&result.CreatedAt,
			&result.UpdatedAt)
		if err != nil {
			log.Printf("Error while scan:%+v\n", err)
		}

		results = append(results, result)
	}

	return results, nil
}

func (c *giftImp) UpdateIsBooked(ctx context.Context, data model.Gifts) error {
	query := `UPDATE tb_gifts SET is_booked = $2, updated_at = $3 WHERE id = $1`
	_, err := c.DB.Exec(ctx, query,
		data.ID,
		data.IsBooked,
		data.UpdatedAt,
	)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return err
	}

	defer c.DB.Close(ctx)

	log.Printf("Booked item gift with ID:%d\n", data.ID)
	return nil
}
