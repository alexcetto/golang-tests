package types

type Task struct {
	Id      int
	Title   string
	Content string
	Created string
}

type Context struct {
	Tasks      []Task
	Navigation string
	Search     string
	Message    string
	CSRFToken  string
	Referer    string
}

type Status struct {
	StatusCode int
	Message    string
}