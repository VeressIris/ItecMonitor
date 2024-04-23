package firestore

import (
	"context"
	"fmt"
	"log"

	firestore "cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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
	Name   string `json:"name"`
	URL    string `json:"url"`
	Status string `json:"status"`
}

type App struct {
	Name      string `json:"name"`
	Developer string `json:"developer"`
	BaseURL   string `json:"baseurl"`
	Endpoints []map[string]interface{}
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

func (firestoreClient *FirestoreClient) WriteAppToDatabase(app App, name string) {
	data := map[string]interface{}{
		"developer": app.Developer,
		"baseURL":   app.BaseURL,
		"endpoints": app.Endpoints,
		"status":    app.Status,
	}
	firestoreClient.WriteToDatabase("apps", data, name)
}

func (firestoreClient *FirestoreClient) WriteEndpointToDatabase(app string, endpoint Endpoint) {
	data := map[string]interface{}{
		"name":   endpoint.Name,
		"url":    endpoint.URL,
		"status": endpoint.Status,
	}
	_, err := firestoreClient.Client.Collection("apps").Doc(app).Set(firestoreClient.ctx, map[string]interface{}{
		"endpoints": firestore.ArrayUnion(data),
	}, firestore.MergeAll)
	if err != nil {
		log.Fatalf("Failed adding endpoint to apps: %v", err)
	}
}

func (firestoreClient *FirestoreClient) GetDevApps(dev string) ([]App, error) {
	apps := make([]App, 0)
	iter := firestoreClient.Client.Collection("apps").Where("developer", "==", dev).Documents(firestoreClient.ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		//TODO: don't read all the endpoints, just the number of endpoints
		// parse endpoints data
		endpointsData, ok := doc.Data()["endpoints"].([]interface{})
		if !ok {
			log.Fatalf("Failed to retrieve endpoints data")
		}
		endpoints := make([]map[string]interface{}, len(endpointsData))
		for i, endpointData := range endpointsData {
			endpoint, ok := endpointData.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("invalid endpoint data type")
			}
			endpoints[i] = endpoint
		}

		// create app
		app := App{
			Name:      doc.Ref.ID,
			Developer: doc.Data()["developer"].(string),
			BaseURL:   doc.Data()["baseURL"].(string),
			Endpoints: endpoints,
			Status:    doc.Data()["status"].(string),
		}

		apps = append(apps, app)
	}
	return apps, nil
}

func (firestoreClient *FirestoreClient) GetEndpoints(appName string) ([]Endpoint, error) {
	doc, err := firestoreClient.Client.Collection("apps").Doc(appName).Get(firestoreClient.ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println(doc.Data()["endpoints"])

	// parse endpoints data
	endpointsData, ok := doc.Data()["endpoints"].([]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to retrieve endpoints data")
	}
	endpoints := make([]Endpoint, len(endpointsData))
	for i, endpointData := range endpointsData {
		endpoint, ok := endpointData.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid endpoint data format")
		}
		endpoints[i] = Endpoint{
			Name:   endpoint["name"].(string),
			URL:    endpoint["url"].(string),
			Status: endpoint["status"].(string),
		}
	}

	return endpoints, nil
}
