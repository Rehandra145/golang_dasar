package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (v *notFoundError) Error() string {
	return v.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"validation error"}
	}

	if id != "Rehan" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("abid", nil)
	if err != nil {
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation eror :", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("not found error :", notFoundErr.Message)
		} else {
			fmt.Println("Unknown eror")
		}
	} else {
		fmt.Println("Sukses")
	}
}
