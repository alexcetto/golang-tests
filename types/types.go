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
}
