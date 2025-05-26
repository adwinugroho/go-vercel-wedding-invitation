package config

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	supa "github.com/nedpals/supabase-go"
)

func InitPostgresConnection(ctx context.Context, host, port, user, password, dbName string) (*pgx.Conn, error) {
	// connection string
	uri := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", user, password, host, port, dbName)
	// open database
	db, err := pgx.Connect(ctx, uri)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil, err
	}

	// check db
	err = db.Ping(ctx)
	if err != nil {
		log.Printf("Error cause:%+v\n", err)
		return nil, err
	}

	log.Println("Connected!")

	return db, nil
}

func InitSupabaseConnection(url, key, password string) *supa.Client {
	supaClient := supa.CreateClient(url, key, false)
	return supaClient
}
