package config

import "os"

var (
	SUPABASE_URL     = os.Getenv("SUPABASE_URL")
	SUPABASE_API_KEY = os.Getenv("SUPABASE_API_KEY")
	SUPABASE_TOKEN   = os.Getenv("SUPABASE_TOKEN")
)
