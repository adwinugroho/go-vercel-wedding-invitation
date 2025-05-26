package config

import "os"

var (
	// postgres
	DB_HOST     = os.Getenv("DB_HOST")
	DB_PORT     = os.Getenv("DB_PORT")
	DB_USER     = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME     = os.Getenv("DB_NAME")
	// supabase
	SUPABASE_URL      = os.Getenv("SUPABASE_URL")
	SUPABASE_API_KEY  = os.Getenv("SUPABASE_API_KEY")
	SUPABASE_PASSWORD = os.Getenv("SUPABASE_PASSWORD")
	// api key
	API_KEY = os.Getenv("API_KEY")
)
