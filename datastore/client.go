package datastore

import (
	"context"
	"log"

	"cloud.google.com/go/datastore"
)

var DSClient *datastore.Client

func InitDatastoreClient() {
	ctx := context.Background()

	client, err := datastore.NewClient(ctx, "")
	if err != nil {
		log.Fatalf("Failed to create Datastore client: %v", err)
	}
	DSClient = client
}
