package tools_test

import (
	"fmt"
	"intro-unit-testing/ej1/tools"
	"testing"
)

func TestElementAtIndexWithExistingIndex(t *testing.T) {
	// arrange
	slice := []int{1, 2, 3, 4, 5}
	idx := 0
	// TODO: Nota al orador:
	/*
		como el test es de caja blanca y sabemos que es lo que hace la funcion,
		podemos asignar expected como el resultado de slice[idx],
		para que quede más claro de donde viene ese valor esperado
	*/
	expected := 1 // TODO: Nota al orador: jugar con los valores para que falle

	// act
	result := tools.ElementAtIndex(slice, idx)

	// assert
	if result != expected {
		// TODO: Nota al orador:
		/*
			contar que si t.Errorf() no se ejecuta nunca,
			el test resulta ok.
			OBSERVACION:
			acotar que t.Errorf() no corta la ejecución del test
			si queremos cortar ante el primer t.Errorf(), debemos poner
			un return inmediatamente despues,
			o envolver el error en una función
		*/
		t.Errorf("test failed. expected %d but %d was given instead", expected, result)
	}

	// TODO: Nota al orador:
	/*
		para mostrar que si pasa por t.Errorf(), el test continua
	*/

	// the test continues
	fmt.Println("test ended")
}

// TODO: Nota al orador
// descomentar el siguiente bloque de código despues de mostrar
// que pasa si la función a probar paniquea

// func TestElementAtIndexWithNonExistingIndex(t *testing.T) {
// 	// arrange
// 	slice := []int{1, 2, 3, 4, 5}
// 	idx := len(slice)

// 	// act
// 	result, err := tools.ElementAtIndex(slice, idx)

// 	// assert
// 	if err == nil {
// 		t.Errorf("test failed. expected error but %d was given", result)
// 	}
// }
