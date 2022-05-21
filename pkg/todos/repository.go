package todos

import (
	"errors"
	"github.com/lithammer/shortuuid"
)

type InMemoryTodoRepository struct {
	Todos []*Todo
}

func NewInMemoryTodoRepository() InMemoryTodoRepository {
	t := new(InMemoryTodoRepository)
	t.Todos = make([]*Todo, 0)
	return *t
}

func (r *InMemoryTodoRepository) Create(todo *Todo) {
	todo.Id = shortuuid.New()
	r.Todos = append(r.Todos, todo)
}

func (r *InMemoryTodoRepository) GetAll() []*Todo {
	return r.Todos
}

func (r *InMemoryTodoRepository) DeleteAll() {
	r.Todos = make([]*Todo, 0)
}

func (r *InMemoryTodoRepository) Get(id string) (t *Todo, err error) {
	for _, t = range r.Todos {
		if t.Id == id {
			return t, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (r *InMemoryTodoRepository) Delete(id string) (err error) {
	for i, t := range r.Todos {
		if t.Id == id {
			r.Todos = append(r.Todos[:i], r.Todos[i+1:]...)
			return nil
		}
	}
	return errors.New("todo not found")
}

func (r *InMemoryTodoRepository) Update(todo *Todo) (err error) {
	for i, t := range r.Todos {
		if t.Id == todo.Id {
			r.Todos[i] = todo
			return nil
		}
	}
	return errors.New("todo not found")
}
