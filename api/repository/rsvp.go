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
	RSVPInterface interface {
		Insert(ctx context.Context, data model.Reservation) (string, error)
		List(ctx context.Context, offset, limit int, isAttending bool) ([]model.Reservation, error)
		InsertWithSupabaseClient(data model.Reservation) (string, error)
		ListWithSupabaseClient(offset, limit int, isAttending string) ([]model.Reservation, error)
	}
)

type rsvpImp struct {
	DB         *pgx.Conn
	DBSupabase *supabase.Client
}

func NewReservationRepository(conn *pgx.Conn, connSupabase *supabase.Client) RSVPInterface {
	return &rsvpImp{
		DB:         conn,
		DBSupabase: connSupabase,
	}
}

func (c *rsvpImp) Insert(ctx context.Context, data model.Reservation) (string, error) {
	data.ID = uuid.New().String()
	defer c.DB.Close(ctx)
	query := `INSERT INTO tb_rsvp (id, name, is_attending, guest_count, created_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := c.DB.Exec(ctx, query, data.ID, data.Name, data.IsAttending, data.GuestCount, data.CreatedAt)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return "", err
	}

	log.Printf("Insert new reservation to DB with ID:%s\n", data.ID)
	return data.ID, nil
}

func (c *rsvpImp) List(ctx context.Context, offset, limit int, isAttending bool) ([]model.Reservation, error) {
	var results []model.Reservation

	// Safety: ensure range is valid
	if limit <= 0 {
		limit = 10 // default limit
	}
	if offset <= 0 {
		offset = 1
	}

	offsetPage := (offset - 1) * limit

	query := `SELECT id, name, is_attending, guest_count, created_at FROM tb_rsvp WHERE is_attending = $3 LIMIT $1 OFFSET $2`

	defer c.DB.Close(ctx)
	rows, err := c.DB.Query(ctx, query, limit, offsetPage, isAttending)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var result model.Reservation

		err = rows.Scan(&result.ID, &result.Name, &result.IsAttending, &result.GuestCount, &result.CreatedAt)
		if err != nil {
			log.Printf("Error while scan:%+v\n", err)
		}

		results = append(results, result)
	}

	return results, nil
}

func (c *rsvpImp) InsertWithSupabaseClient(data model.Reservation) (string, error) {
	var results []model.Wishes
	data.ID = uuid.New().String()
	err := c.DBSupabase.DB.From("tb_rsvp").Insert(data).Execute(&results)
	if err != nil {
		log.Println("error cause:", err)
		return "", err
	}

	log.Println("successfully inserted new rsvp with supabase:", data.ID)
	return data.ID, nil
}

func (c *rsvpImp) ListWithSupabaseClient(offset, limit int, isAttending string) ([]model.Reservation, error) {
	var results []model.Reservation

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
		From("tb_rsvp").
		Select("*").
		LimitWithOffset(limit, offsetPage).
		Eq("is_attending", isAttending).
		Execute(&results)
	if err != nil {
		log.Printf("error listing wishes: %+v\n", err)
		return nil, err
	}

	log.Println("successfully return data rsvp with total:", len(results))
	return results, nil
}
