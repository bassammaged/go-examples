package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Server struct
type server struct {
	interfaceIp string
	port        int
	id          int
	serviceStat bool
	/*
		to do:
			- the server will support communication over tls (bool)
			- the server will allow the package user to set the number of connections (int)
	*/
}

// NewServer() initates server struct with default values
// default values: ip=127.0.0.1, port=8080, id=1
func NewServer() *server {
	// Initialize the a new server struct with default values
	s := defaultConfiguration()
	return &s
}

// Set default configuration to server struct
// Returns server struct with the default configuration
func defaultConfiguration() server {
	return server{
		interfaceIp: "127.0.0.1",
		port:        8080,
		id:          1,
		serviceStat: false,
	}
}

// Blueprint server configuration functions
// type serverConf func(*server) error

// SetIpAddress() set the IP address of the struct
// return error message, if the IP address value is invalid textual respresntation of an IP address.
func (s *server) SetIpAddress(ipAddress string) error {
	if err := s.parseIp(ipAddress); err != nil {
		return err
	}
	s.interfaceIp = ipAddress
	return nil
}

// parseIp() Parse the IPv4 and IPv6.
// Return error if the value is invalid textual representation of an IP address
func (s server) parseIp(ipAddress string) error {
	if net.ParseIP(ipAddress) == nil {
		return errors.New("invalid IP address")
	} else {
		return nil
	}
}

// SetPort() set the port of the struct
// return error message, if the port value is not in range 1-65535.
func (s *server) SetPort(port int) error {
	if err := s.parsePort(port); err != nil {
		return err
	}
	s.port = port
	return nil
}

// parsePort() Parse the server port.
// Return error if the value is not in range 1-65535
func (s server) parsePort(port int) error {
	if port > 0 && port <= 65535 {
		return nil
	}
	return errors.New("invalid port number")
}

// parsePort() Parse the server port.
// Return error if the value is not in range 1-65535
func (s *server) SetId(id int) error {
	if err := s.parseId(id); err != nil {
		return err
	}
	s.id = id
	return nil
}

// parsePort() Parse the server port.
// Return error if the value is not in range 1-200
func (s server) parseId(id int) error {
	min := 1
	max := 200
	if id < min || id > max {
		return errors.New("invalid id number")
	}
	return nil
}

func (s *server) Run() error {

	// Assign the socket address
	port := strconv.Itoa(s.port)
	socketAdderss := s.interfaceIp + ":" + port

	// Print out the current status and change the server status
	log.Printf("GoIam service running on %v:%v\n", s.interfaceIp, s.port)
	s.serviceStat = true

	// Declare router/Multiplexer
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandlefunc(s.handleAccount))
	router.HandleFunc("/account/{id}", makeHTTPHandlefunc(s.handleAccount))

	// Return an error if ListenAndServe is failed to start
	if err := http.ListenAndServe(socketAdderss, router); err != nil {
		return err
	}
	return nil
}

// WriteJson() decorates the response into json and sets the http header and http status code
func WriteJson(w http.ResponseWriter, response ApiResponse) error {
	w.Header().Set("Content-Type", "applicaion/json")
	w.WriteHeader(response.StatusCode)
	return json.NewEncoder(w).Encode(response.Content)
}

// Blueprint (function signature) for our handlers
type apiFunc func(w http.ResponseWriter, r *http.Request) error

// makeHTTPHandlefunc() converts (decortes) any function match with function signature `apiFunc` to `http.HandlerFunc`
// Arguments apiFunc
// Returns http.HandlerFunc
func makeHTTPHandlefunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			// handle the error
			WriteJson(w, ApiResponse{StatusCode: http.StatusBadGateway, Content: err})
		}
	}
}

// handleAccount() function handles all requests that crosponding to /account
func (s *server) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetAccount(w, r)
	} else if r.Method == "POST" {
		return s.HandleCreateAccount(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

type ApiResponse struct {
	StatusCode int
	Content    any
}

type ApiError struct {
	Type    string
	Message string
}

func (s *server) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	varId := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(varId)

	user := NewAccount()
	if err := user.GetAccount(id); err != nil {
		response := ApiResponse{StatusCode: http.StatusOK, Content: err}
		return WriteJson(w, response)
	}
	response := ApiResponse{StatusCode: http.StatusOK, Content: user}
	return WriteJson(w, response)
}

func (s *server) HandleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	response := ApiResponse{StatusCode: http.StatusForbidden, Content: ApiError{Type: "Forbidden", Message: "Access is forbidden"}}
	return WriteJson(w, response)
}
