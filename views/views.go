package views

import (
	"github.com/alexcetto/golang-tests/db"
	"net/http"
	"log"
	"time"
	"html/template"
)

var message string
var homeTemplate *template.Template
var templates *template.Template

// Route /
func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		context := db.GetTasks("PENDING")
		if message != "" {
			context.Message = message
		}
		context.CSRFToken = "abcd"
		message = ""
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires:expiration}
		http.SetCookie(w, &cookie)
		homeTemplate.Execute(w, context)
	} else {
		w.Write([]byte("Method not allowed"))
		http.Redirect(w, r, "/", http.StatusFound)
	}

}

// Route /logout/
func HandleLogout(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Handle logout GET"
	} else {
		message = "Handle logout POST"
	}
	w.Write([]byte(message))
}

// Route /change/
func PostChange(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Post change GET"
	} else {
		message = "Post change POST"
	}
	w.Write([]byte(message))
}

// Route /add_user
func PostAddUser(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Post add user GET"
	} else {
		message = "Post add user POST"
	}
	w.Write([]byte(message))
}

// Route /admin
func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Handle admin GET"
	} else {
		message = "handle admin POST"
	}
	w.Write([]byte(message))
}

// Route /register
func PostRegister(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Post register GET"
	} else {
		message = "Post register POST"
	}
	w.Write([]byte(message))
}

// Route /login/
func GetLogin(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Get login GET"
	} else {
		message = "Get login POST"
	}
	w.Write([]byte(message))
}

// Route /search/
func SearchTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Search task GET"
	} else {
		message = "Search task POST"
	}
	w.Write([]byte(message))
}

// Route /update/
func UpdateTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Update task GET"
	} else {
		message = "Update task POST"
	}
	w.Write([]byte(message))
}

// Route /add/
func AddTaskFunc(w http.ResponseWriter, r *http.Request) {
	title := "Random title"
	content := "Random content"
	truth := db.Addtask(title, content)
	if truth != nil {
		log.Fatal("Error adding task")
	}
	w.Write([]byte("Added task"))
}

// Route /restore/
func RestoreTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Restore task GET"
	} else {
		message = "Restore task POST"
	}
	w.Write([]byte(message))
}

// Route /completed/
func ShowCompleteTasksFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Show complete tasks GET"
	} else {
		message = "Show complete tasks POST"
	}
	w.Write([]byte(message))
}

// Route /edit
func EditTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Edit task GET"
	} else {
		message = "Edit task POST"
	}
	w.Write([]byte(message))
}

// Route /deleted/
func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "show trash tasks GET"
	} else {
		message = "show trash tasks POST"
	}
	w.Write([]byte(message))
}

// Route /trash/
func TrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "trash task GET"
	} else {
		message = "trash task POST"
	}
	w.Write([]byte(message))
}

// Route /delete/
func DeleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "Delete task GET"
	} else {
		message = "Delete task POST"
	}
	w.Write([]byte(message))
}

// Route /complete/
func CompleteTaskFunc(w http.ResponseWriter, r *http.Request) {
	var message string
	if r.Method == "GET" {
		message = "All pending tasks GET"
	} else {
		message = "All pending tasks POST"
	}
	w.Write([]byte(message))
}
