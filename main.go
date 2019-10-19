package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	config "github.com/klipitkas/hooktail/config"
	deployment "github.com/klipitkas/hooktail/deployment"
	request "github.com/klipitkas/hooktail/request"
)

// The content type we want
const (
	ApplicationJSON           string = "application/json"
	ApplicationFormURLEncoded string = "application/x-www-form-urlencoded"
)

var conf config.Config

func main() {
	// The path to the configuration file.
	configPath := ""

	flag.StringVar(&configPath, "config", "config.yml",
		"The configuration file path.")
	flag.Parse()

	// The configuration struct.
	if err := config.Parse(&conf, configPath); err != nil {
		log.Fatalf("parsing configuration: %v", err)
	}

	// The list of request handlers.
	http.HandleFunc("/", handleRequest)

	// Log the server start.
	log.Printf("Starting HTTP server on port: %v", conf.Port)

	// The port in string format.
	portStr := fmt.Sprintf("%d", conf.Port)

	// The server configuration.
	if err := http.ListenAndServe(":"+portStr, nil); err != nil {
		log.Fatalf("listen on port %d failed: %v", conf.Port, err)
	}
}

func handleRequest(w http.ResponseWriter, req *http.Request) {

	// The body of the request.
	body, err := ioutil.ReadAll(req.Body)

	// Extract the content type from the headers.
	contentType := strings.ToLower(req.Header.Get("Content-Type"))
	// Extract the given hash if it exists.
	givenHash := strings.ReplaceAll(
		strings.ToLower(req.Header.Get("X-Hub-Signature")),
		"sha1=",
		"")

	if contentType != ApplicationJSON {
		log.Printf("got invalid request content type: %v", contentType)
		return
	}

	// The expected request struct.
	request := request.Request{}
	request.Headers = req.Header

	if err := json.Unmarshal(body, &request.Body); err != nil {
		log.Printf("cannot unmarshal string: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Error parsing request body to request struct."))
		return
	}

	// Check if request matches a deployment.
	match := deployment.FindMatching(conf.Deployments, request)
	notFound := deployment.Deployment{}

	if match == notFound {
		log.Printf("A deployment that matches the request cannot be found!")
		w.WriteHeader(404)
		w.Write([]byte("A matching deployment was not found."))
		return
	}

	// Check the validity of the request and deployment.
	if match.Secret != "" {
		ok := request.HasValidSignature(match.Secret, string(body), givenHash)
		if !ok {
			log.Printf("Request integrity check failed, please verify the secret!")
			w.WriteHeader(400)
			w.Write([]byte("Invalid secret or signature."))
			return
		}
	}

	// Run the deployment.
	if err = deployment.Deploy(match); err != nil {
		log.Printf("run deployment: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Deployment failed."))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("Deployment successfully completed."))
}
