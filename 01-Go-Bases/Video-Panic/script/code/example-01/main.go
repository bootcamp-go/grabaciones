package main

import "fmt"

/*
	Notas del Orador:
	- En este ejemplo veremos un panic nativo de Go donde al intentar acceder a un indice de un slice que no existe, Go genera un panic.
*/
// main.go
func main() {
	// Create a new store
	store := &StorageItemsSlice{
		Items: []Item{
			{Name:  "Item 1", Price: 9.99},
			{Name:  "Item 2", Price: 19.99},
		},
	}

	// Get an item from the store
	item := store.GetItem(3)

	// Print the item
	fmt.Printf("%+v\n", item)
}


// ./items/storage.go
type Item struct {
	Name  string
	Price float64
}

type StorageItemsSlice struct {
	Items []Item
}

func (s *StorageItemsSlice) GetItem(id int) Item {
	return s.Items[id]
}