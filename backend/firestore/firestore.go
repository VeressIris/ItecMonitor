package firestore

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
)

type FirestoreClient struct {
	Client *firestore.Client
	ctx    context.Context
}

type User struct {
	Email string `json:"email"`
	Name  string `json:"username"`
	UID   string `json:"uid"`
}

type Endpoint struct {
	Path   string `json:"path"`
	Status string `json:"status"`
}

type App struct {
	Developer string `json:"developer"`
	Endpoints []Endpoint
	BaseURL   string `json:"baseurl"`
	Status    string `json:"status"`
}

func GetFirestoreClient(projectID string) (*FirestoreClient, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
		return nil, err
	}
	return &FirestoreClient{Client: client, ctx: ctx}, nil
}

func (firestoreClient *FirestoreClient) WriteToDatabase(collection string, data map[string]interface{}, uid string) {
	_, err := firestoreClient.Client.Collection(collection).Doc(uid).Set(firestoreClient.ctx, data)
	if err != nil {
		log.Fatalf("Failed adding to %s: %v", collection, err)
	}
}

func (firestoreClient *FirestoreClient) WriteUserToDatabase(user User, uid string) {
	data := map[string]interface{}{
		"email": user.Email,
		"name":  user.Name,
	}
	firestoreClient.WriteToDatabase("users", data, uid)
}

func (firestoreClient *FirestoreClient) WriteAppToDatabase(app App, uid string) {
	data := map[string]interface{}{
		"developer": app.Developer,
		"endpoints": app.Endpoints,
		"baseURL":   app.BaseURL,
		"status":    app.Status,
	}
	firestoreClient.WriteToDatabase("apps", data, uid)
}

func (firestoreClient *FirestoreClient) WriteEndpointToDatabase(app string, endpoint Endpoint) {
	data := map[string]interface{}{
		"path":   endpoint.Path,
		"status": endpoint.Status,
	}
	_, err := firestoreClient.Client.Collection("apps").Doc(app).Set(firestoreClient.ctx, map[string]interface{}{
		"endpoints": firestore.ArrayUnion(data),
	}, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed adding endpoint to apps: %v", err)
	}
}
