package tools

import (
	"fmt"
	"github.com/go-playground/validator"
	"log"
)

func Validate(data interface{}) error {
	validate := validator.New()

	err := validate.Struct(data)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Printf("err validator: %v", err)
			return err
		}

		log.Println("------ List of tag fields with error ---------")

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println("---------------")
		}
		return err
	}

	return err
}
