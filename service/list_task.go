package service

import (
	"context"

	"github.com/yukinkoNo10/go_todo_app/entity"
	"github.com/yukinkoNo10/go_todo_app/store"
)

type ListTask struct {
	DB   store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error) {
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil {
		return nil, err
	}
	return ts, nil
}
