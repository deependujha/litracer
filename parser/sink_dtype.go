package parser

// each worker will parase a line of the log file, and send the result to the sink
// the result will comprise of the worker id and the content
type SinkDType struct {
	WorkerID int
	Content  string
}
