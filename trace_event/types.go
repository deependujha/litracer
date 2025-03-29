package trace_event

// TraceEventType represents the type of a trace event.
type TraceEventType string

const (
	// Duration trace events
	BEGIN TraceEventType = "B"
	END   TraceEventType = "E"

	// Complete trace event
	COMPLETE TraceEventType = "X"

	// Instant trace event
	INSTANT TraceEventType = "I"

	// Counter trace event
	COUNTER TraceEventType = "C"

	// Async trace events
	NESTABLE_ASYNC_BEGIN   TraceEventType = "b"
	NESTABLE_ASYNC_END     TraceEventType = "e"
	NESTABLE_ASYNC_INSTANT TraceEventType = "n"

	// Flow trace events
	FLOW_BEGIN TraceEventType = "s"
	FLOW_STEP  TraceEventType = "t"
	FLOW_END   TraceEventType = "f"

	// Metadata trace events
	METADATA TraceEventType = "M"

	// Sample trace event
	SAMPLE TraceEventType = "P"

	// Object trace events
	CREATE_OBJECT   TraceEventType = "N"
	SNAPSHOT_OBJECT TraceEventType = "O"
	DELETE_OBJECT   TraceEventType = "D"

	// Memory dump trace events
	MEMORY_DUMP_GLOBAL TraceEventType = "V"
	MEMORY_DUMP        TraceEventType = "v"

	// Mark trace event
	MARK TraceEventType = "R"

	// Clock sync event
	CLOCK_SYNC TraceEventType = "c"
)
