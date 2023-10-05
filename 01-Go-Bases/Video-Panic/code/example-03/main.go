package main

// main.go
func main() {
	// Create a new store
	store := &StorageItemsSlice{}

	// Add an item to the store
	store.AddItem(&Item{Name: "Item 1", Price: 9.99})
}


// ./items/storage.go
type Item struct {
	Name  string
	Price float64
}

type StorageItemsSlice struct {
	Items  map[int]Item
	LastId int
}

func (s *StorageItemsSlice) GetItem(id int) Item {
	return s.Items[id]
}

func (s *StorageItemsSlice) AddItem(item *Item) {
	// validate item
	if item.Name == "" {
		return
	}

	// add item to map
	s.LastId++
	s.Items[s.LastId] = *item
}