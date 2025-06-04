package repository

import (
	"context"
	"log"

	"github.com/adwinugroho/go-vercel-wedding-invitation/api/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/nedpals/supabase-go"
)

type (
	WishesInterface interface {
		Insert(ctx context.Context, data model.Wishes) (string, error)
		List(ctx context.Context, offset, limit int) ([]model.Wishes, error)
		InsertWithSupabaseClient(data model.Wishes) (string, error)
		ListWithSupabaseClient(offset, limit int) ([]model.Wishes, error)
	}
)

type wishesImp struct {
	DB         *pgx.Conn
	DBSupabase *supabase.Client
}

func NewWishesRepository(conn *pgx.Conn, connSupabase *supabase.Client) WishesInterface {
	return &wishesImp{
		DB:         conn,
		DBSupabase: connSupabase,
	}
}

func (c *wishesImp) Insert(ctx context.Context, data model.Wishes) (string, error) {
	data.ID = uuid.New().String()
	defer c.DB.Close(ctx)
	query := `INSERT INTO tb_wishes (id, name, message, is_published, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := c.DB.Exec(ctx, query, data.ID, data.Name, data.Message, data.IsPublished, data.CreatedAt)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return "", err
	}

	log.Printf("Insert new wishes to DB with ID:%s\n", data.ID)
	return data.ID, nil
}

func (c *wishesImp) List(ctx context.Context, offset, limit int) ([]model.Wishes, error) {
	var results []model.Wishes

	// Safety: ensure range is valid
	if limit <= 0 {
		limit = 10 // default limit
	}
	if offset <= 0 {
		offset = 1
	}

	offsetPage := (offset - 1) * limit

	query := `SELECT id, name, message, created_at FROM tb_wishes WHERE is_published = true 
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
		var result model.Wishes

		err = rows.Scan(&result.ID, &result.Name, &result.Message, &result.CreatedAt)
		if err != nil {
			log.Printf("Error while scan:%+v\n", err)
		}

		results = append(results, result)
	}

	return results, nil
}

func (c *wishesImp) InsertWithSupabaseClient(data model.Wishes) (string, error) {
	var results []model.Wishes
	data.ID = uuid.New().String()
	err := c.DBSupabase.DB.From("tb_wishes").Insert(data).Execute(&results)
	if err != nil {
		log.Println("error cause:", err)
		return "", err
	}

	log.Println("successfully inserted new data with supabase:", data.ID)
	return data.ID, nil
}

func (c *wishesImp) ListWithSupabaseClient(offset, limit int) ([]model.Wishes, error) {
	var results []model.Wishes

	// Safety: ensure range is valid
	if limit <= 0 {
		limit = 10 // default limit
	}
	if offset <= 0 {
		offset = 1
	}

	offsetPage := (offset - 1) * limit

	err := c.DBSupabase.
		DB.
		From("tb_wishes").
		Select("*").
		LimitWithOffset(limit, offsetPage).
		Eq("is_published", "true").
		Execute(&results)
	if err != nil {
		log.Printf("error listing wishes: %+v\n", err)
		return nil, err
	}

	log.Println("successfully return data with total:", len(results))
	return results, nil
}
