package handler

import (
	"context"

	"github.com/yukinkoNo10/go_todo_app/entity"
)

type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name string, password string, role string) (*entity.User, error)
}
