package data

import (
	"task4/models"
	"time"
	"errors"
)
func ValidateForCreation(t models.Task)error{
	for _,v := range Store{
		if v.ID == t.ID{
			return errors.New("id already in use")
		}
	}
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
	for _,v := range Store{
		if v.ID == t.ID{
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
	}
	return errors.New("id does not exist")
}
