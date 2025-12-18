package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func createTaskPage(w http.ResponseWriter, r *http.Request) {
	taskName := "Complete Golang project"
	deadline := "2024-09-15"
	priority := "High"
	user := "Alice"

	log.WithFields(logrus.Fields{
		"page":     "createTask",
		"method":   r.Method,
		"path":     r.URL.Path,
		"taskName": taskName,
		"deadline": deadline,
		"priority": priority,
		"user":     user,
	}).Info("Task Created")

	fmt.Fprintf(w, "Task created successfully!")
}

func viewTasksPage(w http.ResponseWriter, r *http.Request) {
	user := "Alice"
	taskCount := 5

	log.WithFields(logrus.Fields{
		"page":      "viewTasks",
		"method":    r.Method,
		"path":      r.URL.Path,
		"user":      user,
		"taskCount": taskCount,
	}).Info("Tasks Viewed")

	fmt.Fprintf(w, "Displaying %d tasks.", taskCount)
}

func deleteTaskPage(w http.ResponseWriter, r *http.Request) {
	taskName := "Complete Golang project"
	user := "Alice"

	log.WithFields(logrus.Fields{
		"page":     "deleteTask",
		"method":   r.Method,
		"path":     r.URL.Path,
		"taskName": taskName,
		"user":     user,
	}).Info("Task Deleted")

	fmt.Fprintf(w, "Task '%s' deleted successfully!", taskName)
}

func errorTaskPage(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{
		"page":   "errorTask",
		"method": r.Method,
		"path":   r.URL.Path,
	}).Info("Error Page Accessed - Simulating Error")
	// Simulate an error
	if _, err := fmt.Fprintf(w, "An error occurred while managing tasks!"); err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Error("Failed to write response")
	}
}

func handleRequests() {
	http.HandleFunc("/", createTaskPage)
	http.HandleFunc("/view", viewTasksPage)
	http.HandleFunc("/delete", deleteTaskPage)
	http.HandleFunc("/error", errorTaskPage)
	log.Info("Starting Task Management Server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}

func main() {
	// Configure logrus
	log.Out = os.Stdout // Explicitly setting output to stdout
	log.Formatter = &logrus.JSONFormatter{}
	log.Level = logrus.InfoLevel

	log.Info("Starting Task Management Server on port 8080")
	handleRequests()
}
