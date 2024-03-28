package main

import "fmt"

type validationError struct {
	Message string
}

type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{
			Message: "id cannot be empty",
		}
	}

	if id != "123" {
		return &notFoundError{
			Message: "data not found",
		}
	}

	return nil
}

func main() {
	err := SaveData("2222222", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println(validationErr)
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println(notFoundErr)
		// } else {
		// 	fmt.Println("unknown error", err.Error())
		// }

		//versi switch
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error :", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error :", finalError.Error())
		default:
			fmt.Println("error tidak diketahui :", finalError.Error())
		}
	} else {
		fmt.Println("success")
	}
}