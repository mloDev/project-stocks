package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"cloud.google.com/go/pubsub"
)

const topicName = "stocks"

type app struct {
	pubsubClient    *pubsub.Client
	pubsubTopic     *pubsub.Topic
	tmpl            *template.Template
}

func main() {
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	if projectID == "" {
		log.Fatalf("GOOGLE_CLOUD_PROJECT must be set")
	}

	a, err := newApp(projectID, "")
	if err != nil {
		log.Fatalf("newApp: %v", err)
	}

	http.HandleFunc("/", a.index)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Listening on localhost:%v", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func newApp(projectID, templateDir string) (*app, error) {
	ctx := context.Background()

	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		return nil, fmt.Errorf("pubsub.NewClient: %v", err)
	}

	pubsubTopic := pubsubClient.Topic(topicName)

	tmpl, err := template.ParseFiles(filepath.Join(templateDir, "index.html"))
	if err != nil {
		return nil, fmt.Errorf("template.New: %v", err)
	}

	return &app{
		pubsubClient: pubsubClient,
		pubsubTopic:  pubsubTopic,

		tmpl:            tmpl,
	}, nil
}

func (a *app) index(w http.ResponseWriter, r *http.Request) {
	if err := a.tmpl.Execute(w, nil); err != nil {
		log.Printf("tmpl.Execute: %v", err)
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}