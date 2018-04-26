package main

// Config
type Config struct {
	Scripts []Scripts
}

type Scripts struct {
	Name        string
	Description string
	Executor    Executor
	Actions     string
}

type Executor struct {
	Path         string
	Script       string
	ActionMapper map[string]string
}
