package viewmodel

import (
	"context"
	"fmt"
	
	"github.com/CosmicPredator/chibi/internal/kvdb"
	"github.com/CosmicPredator/chibi/internal/ui"
)

// handler func to log user out from AniList
// this is achieved by just deleting the config/chibi folder (for now)

// TODO: Implement proper logout operations
func HandleLogout() error {
	db, err := kvdb.Open()
	if err != nil {
		return err
	}
	defer db.Close()
	
	err1 := db.Delete(context.TODO(), "auth_token")
	err2 := db.Delete(context.TODO(), "user_id")
	err3 := db.Delete(context.TODO(), "user_name")

	if err1 != nil || err2 != nil || err3 != nil {
		return fmt.Errorf("errors occurred during logout: token=%v userId=%v userName=%v", err1, err2, err3)
	}
	fmt.Println(ui.SuccessText("Logged out successfully!"))
	return nil
}