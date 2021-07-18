package server

// Runner requires Run()
type Runner interface {
	Run() error
}

// Closer requires Close()
type Closer interface {
	Close() error
}

// RunCloser needs Runner and Closer
type RunCloser interface {
	Runner
	Closer
}
