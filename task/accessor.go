package task

import (
	"errors"
)

// ErrTaskNotExist occurs when the task with the ID was not found.
var ErrTaskNotExist = errors.New("task does not exist")

// ID is a data type to identify the task.
type ID string

// Accessor is an interface to access the tasks.
type Accessor interface {
	Get(id ID) (Task, error)
	Put(id ID, t Task) error
	Post(t Task) (ID, error)
	Delete(id ID) error
}
