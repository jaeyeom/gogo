// Package task is an example for struct and JSON code.
package task

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

type status int

const (
	UNKNOWN status = iota
	TODO
	DONE
)

// String returns the string representation of s. This can be generated
// by stringer tool, though this function is hand-written.
func (s status) String() string {
	switch s {
	case UNKNOWN:
		return "UNKNOWN"
	case TODO:
		return "TODO"
	case DONE:
		return "DONE"
	default:
		return ""
	}
}

// MarshalJSON returns the string representation of s.
func (s status) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == "" {
		return nil, errors.New("status.MarshalJSON: unknown value")
	}
	return []byte(fmt.Sprintf("\"%s\"", str)), nil
}

// UnmarshalJSON parses the string representation of status and stores
// it to s.
func (s *status) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"UNKNOWN"`:
		*s = UNKNOWN
	case `"TODO"`:
		*s = TODO
	case `"DONE"`:
		*s = DONE
	default:
		return errors.New("status.UnmarshalJSON: unknown value")
	}
	return nil
}

// Deadline is a struct to hold the deadline time.
type Deadline struct {
	time.Time
}

// NewDeadline returns a newly created Deadline with time t.
func NewDeadline(t time.Time) *Deadline {
	return &Deadline{t}
}

// MarshalJSON returns the Unix time of d.
func (d Deadline) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, d.Unix(), 10), nil
}

// UnmarshalJSON parses the Unix time and stores the result in d.
func (d *Deadline) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	d.Time = time.Unix(unix, 0)
	return nil
}

// Task is a struct to hold a single task.
type Task struct {
	Title    string    `json:"title,omitempty"`
	Status   status    `json:"status,omitempty"`
	Deadline *Deadline `json:"deadline,omitempty"`
	Priority int       `json:"priority,omitempty"`
	SubTasks []Task    `json:"subTasks,omitempty"`
}

// String returns the string representation of t excluding sub tasks.
func (t Task) String() string {
	check := "v"
	if t.Status != DONE {
		check = " "
	}
	return fmt.Sprintf("[%s] %s %s", check, t.Title, t.Deadline)
}

// IncludeSubTasks is a Task but its String method returns the string
// including the sub tasks.
type IncludeSubTasks Task

func (t IncludeSubTasks) indentedString(prefix string) string {
	str := prefix + Task(t).String()
	for _, st := range t.SubTasks {
		str += "\n" + IncludeSubTasks(st).indentedString(prefix+"  ")
	}
	return str
}

// String returns the string representation of t including the sub
// tasks.
func (t IncludeSubTasks) String() string {
	return t.indentedString("")
}
