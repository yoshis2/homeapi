package firebases

import (
	"context"
	"os"

	"cloud.google.com/go/firestore"
	"firebase.google.com/go"
	"google.golang.org/api/option"
)

type Firestore struct {
}

func (firestores *Firestore) Open() (*firestore.Client, error) {
	opt := option.WithCredentialsFile(os.Getenv("SERVICE_ACCOUNT_KEY"))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}

	return client, err
}
