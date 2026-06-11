package email

import (
	"errors"
	"strings"
)

// ValidarEmail valida si un email tiene un formato básico correcto.
// Retorna un error si el formato es inválido.
func ValidarEmail(email string) error {

	if !strings.Contains(email, "@") {
		return errors.New("el email no contiene el símbolo @")
	}

	if !strings.Contains(email, ".") {
		return errors.New("el email no contiene un punto")
	}

	return nil
}
