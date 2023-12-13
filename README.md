This repository contains the code for a Go REST API that integrates with Google's Generative Language Model (GLM) to generate text based on prompts.

## Features
- Exposes an endpoint to receive text prompts as JSON requests.
- Uses the GLM "text-bison-001" model to generate text.
- Returns the generated text as JSON response.

## Usage
### Prerequisites:
- Go version 1.18 or later
- Google Cloud AI Platform account with access to GLM
- API key for GLM access
## Installation:
- Clone this repository.
- Install dependencies: `go mod download`
## Configuration:
- Replace AIzaSyD3JAzzadnoAnT1zb9A-80Cgx75nfkhFk0 in main.go with your GLM API key.
## Running the server:
    go run main.go
## Sending requests:
- Use HTTP client like Postman to send POST request to `http://localhost:8080/generate-text`
- Include a JSON body with the following structure:
### JSON
    {
      "prompt": "Your text prompt here"
    }
Use code with caution. Learn more
## Response:
The server will respond with a JSON object containing the generated text:
### JSON
    {
      "generated_text": "Your generated text here"
    }

## Technical details
This API uses the Google Cloud AI Platform Generative Language Model API with the "text-bison-001" model.
The API is implemented using the Gorilla Mux router and the Go RESTful package.
## Contributing
Contributions are welcome! Please feel free to open pull requests with improvements, bug fixes, or new features.

