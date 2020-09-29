package suma

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestSumar usa el paquete que viene para testing en golang.
func TestSumar(t *testing.T) {
	var suma int = Sumar(3, 8)

	if suma != 11 {
		t.Error("El resultado de la suma debería ser 11")
	}
}

// TestSumarConTestify usa el paquete assert de testify para probar,
// esta dependencia la manejamos con go modules.
func TestSumarConTestify(t *testing.T) {
	var suma int = Sumar(3, 8)
	assert.Equal(t, 11, suma, "El resultado de la suma debería ser 11")
}
