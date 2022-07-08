package validations

import (
	"log"

	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

// ConfigurarValidadores registra os validadores configurados nos
// handlers
func ConfigurarValidadores() {
	binding.Validator = new(defaultValidator)

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("customerDocument", Document); err != nil {
			log.Fatal(err)
		}
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("stringField", StringField); err != nil {
			log.Fatal(err)
		}
	}
}
