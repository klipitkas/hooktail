package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	conf, err := config.Parse(configPath)
	if err != nil {
		log.Fatalf("parsing configuration: %v", err)
	}

	// The list of request handlers.
	http.HandleFunc("/", handleRequest)

	// Log the server start.
	log.Printf("Starting HTTP server on port: %v", conf.Port)

	// The port.
	portStr := strconv.Itoa(conf.Port)

	// The server configuration.
	if err = http.ListenAndServe(":"+portStr, nil); err != nil {
		log.Fatalf("listen on port %d failed: %v", conf.Port, err)
	}

}

func handleRequest(w http.ResponseWriter, req *http.Request) {

	// The body of the request.
	body, err := ioutil.ReadAll(req.Body)

	// Extract the content type from the headers.
	contentType := strings.ToLower(req.Header.Get("Content-Type"))

	if contentType != ApplicationJSON {
		log.Printf("got invalid request content type: %v", contentType)
		return
	}

	// The expected request struct.
	request := request.Request{}

	if err := json.Unmarshal(body, &request); err != nil {
		log.Printf("cannot unmarshal string: %v", err)
		return
	}

	// Log the request
	log.Printf("got request: %+v", request)

	// Check if request matches a deployment.
	match := deployment.FindMatching(conf, request)

	// Run the deployment.
	if err = deployment.Deploy(match); err != nil {
		log.Printf("run deployment: %v", err)
	}

}
