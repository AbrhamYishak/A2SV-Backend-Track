package Usecases

import (
	"task7/Domain"
    "time"
	"errors"
)
type TaskUsecase struct{
	Repo TaskRepoI
}
func NewtaskUsecase (repo TaskRepoI) *TaskUsecase{
	return &TaskUsecase{
		Repo: repo,
	}
}
func (tu *TaskUsecase) CheckTask (task *Domain.Task)error{
	if task.Description == "" || task.Title == ""{
		return errors.New("description and title can not empty")
	}
	if task.Duedate.Before(time.Now()){
		return errors.New("could not set time before now")
	}
	return nil
}
func (tu *TaskUsecase) CreateTask ( task *Domain.Task)error{
	if err := tu.CheckTask(task); err != nil{
		return err
	}
	err := tu.Repo.CreateTasks(task)
	return err
}
func (tu *TaskUsecase) GetTasks () ([]Domain.Task,error){
	tasks, err := tu.Repo.GetTasks()
	return tasks, err
}
