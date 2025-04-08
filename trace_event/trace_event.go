package trace_event

import (
	"encoding/json"
)

// TraceEvent represents a trace event.
type TraceEvent struct {
	Name        string            `json:"name"`
	EventType   TraceEventType    `json:"ph"`
	TimestampUs float64           `json:"ts"`
	DurationUs  float64           `json:"dur,omitempty"`
	ProcessID   int               `json:"pid"`
	ThreadID    *int              `json:"tid,omitempty"`
	Categories  string            `json:"cat,omitempty"`
	Cname       string            `json:"cname,omitempty"`
	Args        map[string]string `json:"args,omitempty"`
	Scope       *string           `json:"s,omitempty"`
}

func (te *TraceEvent) ToJSON() (string, error) {
	data, err := json.Marshal(te)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func FromJSON(data string) (*TraceEvent, error) {
	var te TraceEvent
	err := json.Unmarshal([]byte(data), &te)
	if err != nil {
		return nil, err
	}
	return &te, nil
}

func NewDurationBegin(name string, timestampUs float64, processID int, threadID *int, categories []string, args map[string]string) TraceEvent {
	return TraceEvent{
		Name:        name,
		EventType:   BEGIN,
		TimestampUs: timestampUs,
		ProcessID:   processID,
		ThreadID:    threadID,
		Categories:  categoriesToString(categories),
		Args:        args,
	}
}

func NewDurationEnd(name string, timestampUs float64, processID int, threadID *int, categories []string, args map[string]string) TraceEvent {
	return TraceEvent{
		Name:        name,
		EventType:   END,
		TimestampUs: timestampUs,
		ProcessID:   processID,
		ThreadID:    threadID,
		Categories:  categoriesToString(categories),
		Args:        args,
	}
}

func NewComplete(name string, timestampUs, durationUs float64, processID int, threadID *int, categories []string, args map[string]string) TraceEvent {
	return TraceEvent{
		Name:        name,
		EventType:   COMPLETE,
		TimestampUs: timestampUs,
		DurationUs:  durationUs,
		ProcessID:   processID,
		ThreadID:    threadID,
		Categories:  categoriesToString(categories),
		Args:        args,
	}
}

func NewInstantGlobalScope(name string, timestampUs float64, categories []string) TraceEvent {
	scope := "g"
	return TraceEvent{
		Name:        name,
		EventType:   INSTANT,
		TimestampUs: timestampUs,
		Categories:  categoriesToString(categories),
		Scope:       &scope,
	}
}

func NewProcessName(processID int, processName string) TraceEvent {
	return TraceEvent{
		Name:        "process_name",
		ProcessID:   processID,
		Args:        map[string]string{"name": processName},
		EventType:   METADATA,
		TimestampUs: 0,
	}
}

func NewThreadName(processID int, threadID int, threadName string) TraceEvent {
	return TraceEvent{
		Name:        "thread_name",
		ProcessID:   processID,
		ThreadID:    &threadID,
		EventType:   METADATA,
		TimestampUs: 0,
		Args:        map[string]string{"name": threadName},
	}
}

func categoriesToString(categories []string) string {
	if len(categories) == 0 {
		return ""
	}
	return categories[0]
}
