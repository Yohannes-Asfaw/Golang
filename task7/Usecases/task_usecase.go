package Usecases

import (
	"context"
	"task7/Domain"
)


type TaskUseCase struct {
	TaskRepo Domain.TaskRepository
}

func NewTaskUseCase(repo Domain.TaskRepository) Domain.TaskUseCase {
	return &TaskUseCase{
		TaskRepo: repo,
	}
}

func (uc *TaskUseCase) Create(ctx context.Context, payload *Domain.Task) (*Domain.Task, error) {
	return uc.TaskRepo.Create(ctx, payload)
}

func (uc *TaskUseCase) Update(ctx context.Context, taskId string, payload *Domain.Task) (*Domain.Task, error) {
	return uc.TaskRepo.Update(ctx, taskId, payload)
}

func (uc *TaskUseCase) Delete(ctx context.Context, taskId string) error {
	return uc.TaskRepo.Delete(ctx, taskId)
}

func (uc *TaskUseCase) GetAll(ctx context.Context) (*[]*Domain.Task, error) {
	return uc.TaskRepo.GetAll(ctx)
}

func (uc *TaskUseCase) GetById(ctx context.Context, taskId string) (*Domain.Task, error) {
	return uc.TaskRepo.GetById(ctx, taskId)
}
