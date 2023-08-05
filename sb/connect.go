package sb

import (
	"os"

	"github.com/nedpals/supabase-go"
)

var SB *supabase.Client

func Connect() {
	SB = supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
}
