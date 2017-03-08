package api

import "net/http"

// HostIndex displays a list of all hosts
func HostIndex(w http.ResponseWriter, req *http.Request) {}

// HostCreate creates a new host
func HostCreate(w http.ResponseWriter, req *http.Request) {}

// HostShow returns a specific host
func HostShow(w http.ResponseWriter, req *http.Request) {}

// HostUpdate updates a specific host
func HostUpdate(w http.ResponseWriter, req *http.Request) {}

// HostDestroy deletes a specific host
func HostDestroy(w http.ResponseWriter, req *http.Request) {}
