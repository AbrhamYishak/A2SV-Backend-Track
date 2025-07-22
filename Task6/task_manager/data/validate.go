package data

import (
	"task6/models"
	"time"
	"errors"
)
func ValidateForCreation(t models.Task)error{
    if t.Description == ""{
		return errors.New("can not set Description empty")
	}
	if t.Duedate.Before(time.Now()){
	   return errors.New("can not set a time before now")	
	}
	if t.Title == ""{
		return errors.New("can not set Title empty")
	}
	return nil
}
func ValidateForEdit(t models.Task)error{
    if t.Description == ""{
		return errors.New("can not set Description empty")
	}
	if t.Duedate.Before(time.Now()){
	   	return errors.New("can not set a time before now")	
	}
	if t.Title == ""{
		return errors.New("can not set Title empty")
	}
	return nil
}
