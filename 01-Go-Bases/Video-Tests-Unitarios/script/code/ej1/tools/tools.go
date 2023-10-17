package tools

func ElementAtIndex(slice []int, index int) int {
	return slice[index]
}

// TODO: nota al orador
/*
	descomentar para mostrar que pasa si el test anterior
	paniquea al recibir un indice fuera del rango,
	como mejora para probar casos borde
*/

// func ElementAtIndex(slice []int, index int) (int, error) {
// 	if index < 0 || index >= len(slice) {
// 		return 0, fmt.Errorf("index out of boundaries")
// 	}
// 	return slice[index], nil
// }
