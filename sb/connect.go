package sb

import (
	"log"
	"os"

	"github.com/nedpals/supabase-go"
)

var SB *supabase.Client

func Connect() {
	SB = supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_ADMIN_KEY"))
	if SB.BaseURL == "" {
		log.Fatal("can't connect to supabase!")
	}
}
