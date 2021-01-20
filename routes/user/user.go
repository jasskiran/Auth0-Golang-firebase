package user

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"OAuth/app"
	"OAuth/firebase"
	"OAuth/routes/templates"
)

// User struct
type User struct {
	Name     string `json:"name,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Picture  string `json:"picture,omitempty"`
}

//Profile struct to retrieve data from session
type Profile struct {
	Aud        string `json:"aud"`
	Exp        int    `json:"exp"`
	FamilyName string `json:"family_name"`
	GivenName  string `json:"given_name"`
	Iat        int    `json:"iat"`
	Iss        string `json:"iss"`
	Locale     string `json:"locale"`
	Name       string `json:"name"`
	Nickname   string `json:"nickname"`
	Picture    string `json:"picture"`
	Sub        string `json:"sub"`
	UpdatedAt  string `json:"updated_at"`
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	session, err := app.Store.Get(r, "auth-session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	mapB, _ := json.Marshal(session.Values["profile"])

	var profile Profile
	err = json.Unmarshal(mapB, &profile)
	if err != nil {
		log.Println(err)
	}

	fbClient, err := firebase.NewClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctx := context.Background()
	db, err := fbClient.Database(ctx)
	fmt.Println("db", db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get a database reference
	ref := db.NewRef("server/saving-data/jass")
	usersRef := ref.Child("users")
	err = usersRef.Set(ctx, map[string]*User{
		"User": {
			Name:     profile.Name,
			Nickname: profile.Nickname,
			Picture:  profile.Picture,
		},
	})
	if err != nil {
		log.Fatalln("Error setting value:", err)
	}
	templates.RenderTemplate(w, "user", session.Values["profile"])
}
