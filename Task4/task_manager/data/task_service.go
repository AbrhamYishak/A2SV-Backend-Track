package data

import (
	"net/http"
	"task4/models"
	"errors"
)
var Store []models.Task
func GetByID(id string) (int , models.Task){
	for _,v := range Store{
		if v.ID == id{
			return http.StatusOK, v
		}
	}
	return http.StatusNotFound, models.Task{}
}
func Addtask(t models.Task) (int ,error){
	Store = append(Store,t)
	return http.StatusCreated, nil
}
func Deltask(id string) (int , error){ 
	for i,v := range Store{
		if v.ID == id{
			Store = append(Store[:i],Store[i+1:]...)	
            return http.StatusOK, nil
		}
	}
	return http.StatusNotFound, errors.New("could not find task with given id")
}
func EditTask(t models.Task,id string) (int, error){
	for i,v := range Store{
		if v.ID == t.ID && v.ID == id{
			Store[i].Title = t.Title
			Store[i].Description = t.Description
			Store[i].Duedate = t.Duedate
			Store[i].Status = t.Status
			return http.StatusOK, nil
		}
	}
	return http.StatusNotFound, errors.New("could not find task with given id")
}
