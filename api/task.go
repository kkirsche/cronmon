package api

import "net/http"

// TaskIndex displays a list of all tasks
func TaskIndex(w http.ResponseWriter, req *http.Request) {}

// TaskCreate creates a new task
func TaskCreate(w http.ResponseWriter, req *http.Request) {}

// TaskShow returns a specific task
func TaskShow(w http.ResponseWriter, req *http.Request) {}

// TaskUpdate updates a specific task
func TaskUpdate(w http.ResponseWriter, req *http.Request) {}

// TaskDestroy deletes a specific task
func TaskDestroy(w http.ResponseWriter, req *http.Request) {}
