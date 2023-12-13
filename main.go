package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	gl "cloud.google.com/go/ai/generativelanguage/apiv1beta2"
	pb "cloud.google.com/go/ai/generativelanguage/apiv1beta2/generativelanguagepb"
	"github.com/gorilla/mux"
	"google.golang.org/api/option"
)

func generateTextHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// Extract the text prompt from the request body
	var requestBody struct {
		Prompt string `json:"prompt"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	client, err := gl.NewTextRESTClient(ctx, option.WithAPIKey("AIzaSyD3JAzzadnoAnT1zb9A-80Cgx75nfkhFk0"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer client.Close()

	req := &pb.GenerateTextRequest{
		Model: "models/text-bison-001",
		Prompt: &pb.TextPrompt{
			Text: requestBody.Prompt,
		},
	}

	resp, err := client.GenerateText(ctx, req)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the generated text as the response
	response := struct {
		GeneratedText string `json:"generated_text"`
	}{GeneratedText: resp.Candidates[0].Output}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	r := mux.NewRouter()

	// Define your API routes
	r.HandleFunc("/generate-text", generateTextHandler).Methods("POST")

	// Start the server
	port := ":8080"
	srv := &http.Server{
		Handler:      r,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Server listening on port %s...\n", port)
	log.Fatal(srv.ListenAndServe())
}
