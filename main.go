package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	config "github.com/klipitkas/hooktail/config"
	deployment "github.com/klipitkas/hooktail/deployment"
	"github.com/klipitkas/hooktail/logging"
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
		logging.Log.Fatalf("parsing configuration: %v", err)
	}

	// The list of request handlers.
	http.HandleFunc("/", handleRequest)

	// Log the server start.
	logging.Log.Printf("Starting HTTP server on port: %v", conf.Port)

	// The port in string format.
	portStr := fmt.Sprintf("%d", conf.Port)

	// The server configuration.
	if err := http.ListenAndServe(":"+portStr, nil); err != nil {
		logging.Log.Fatalf("listen on port %d failed: %v", conf.Port, err)
	}
}

func handleRequest(w http.ResponseWriter, req *http.Request) {

	// The body of the request.
	body, err := ioutil.ReadAll(req.Body)

	// Extract the content type from the headers.
	contentType := strings.ToLower(req.Header.Get("Content-Type"))
	if contentType != ApplicationJSON {
		logging.Log.Errorf("got invalid request: %v", body)
		w.WriteHeader(200)
		w.Write([]byte("I don't speak this language."))
		return
	}

	// Construct the request struct.
	request := request.Request{
		Headers:  req.Header,
		JSONBody: string(body),
	}

	if err = request.Parse(body); err != nil {
		logging.Log.Errorf("cannot unmarshal string: %v", err)
		w.WriteHeader(500)
		w.Write([]byte("Error parsing request body to request struct."))
		return
	}

	// Check if request matches a deployment.
	match := deployment.FindMatching(conf.Deployments, request)
	notFound := deployment.Deployment{}

	if match == notFound {
		logging.Log.Warnf("A deployment that matches the request cannot be found!")
		w.WriteHeader(404)
		w.Write([]byte("A matching deployment was not found."))
		return
	}

	// Check the validity of the request and deployment.
	if match.Secret != "" {
		validSignature := request.HasValidSignature(match.Secret)
		if !validSignature {
			logging.Log.Errorf("Request integrity check failed, please verify the " +
				"secret!")
			w.WriteHeader(400)
			w.Write([]byte("Invalid secret or signature."))
			return
		}
	}

	// Respond timely to the webook.
	w.WriteHeader(200)
	w.Write([]byte("Deployment has started."))

	// Run the deployment.
	go deployment.Deploy(match)
}
