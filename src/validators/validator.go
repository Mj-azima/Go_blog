package validators

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=3,max=32"`
}

type Post struct {
	Body string `validate:"required,min=3,max=10000"`
}

func ValidateStruct(payload interface{}) {

	var validate *validator.Validate
	validate = validator.New()

	err := validate.Struct(payload)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}
	println("success!")

}
