package firebase

import (
	"context"
	"fmt"
	"log"
	"os"

	fb "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

type Config struct {
	CredentialsFilePath string
	ProjectID           string
	DatabaseURL         string
}

// Client is a Firebase client.
type Client struct {
	AuthClient *auth.Client
	App        *fb.App
}

// NewClient returns a new Firebase client.
func NewClient() (*Client, error) {
	envConfig := Config{
		CredentialsFilePath: os.Getenv("FIREBASE_APPLICATION_CREDENTIALS"),
		ProjectID:           os.Getenv("PROJECT_ID"),
		DatabaseURL:         os.Getenv("DATABASE_URL"),
	}
	conf := &fb.Config{
		ProjectID:   envConfig.ProjectID,
		DatabaseURL: envConfig.DatabaseURL,
	}
	opts := option.WithCredentialsFile(envConfig.CredentialsFilePath)
	app, err := fb.NewApp(context.Background(), conf, opts)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}
	return &Client{
		AuthClient: client,
		App:        app,
	}, nil
}

// GenerateCustomToken generates a firebase token.
func (c *Client) GenerateCustomToken(uid string) (string, error) {
	token, err := c.AuthClient.CustomToken(context.Background(), uid)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
		return "", err
	}
	log.Printf("Got custom token: %v\n", token)
	return token, err
}

// Create a database client
func (c *Client) Database(ctx context.Context) (*db.Client, error) {
	client, err := c.App.Database(ctx)
	fmt.Println("client", client)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}

	// Get a database reference
	//ref := client.NewRef("server/saving-data/fireblog")
	return client, nil
}
