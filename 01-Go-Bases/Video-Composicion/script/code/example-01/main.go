package main

import "fmt"

// Publisher is a struct that represents a publisher
type Publisher struct {
	// Name represents the name of the publisher
	Name     string
	// Founders represents the founders of the publisher
	Founders []string
}

func (p Publisher) TotalFounders() int {
	return len(p.Founders)
}

// Book is a struct that represents a book
type Book struct {
	// Title represents the title of the book
	Title string
	// Author represents the author of the book
	Author string
	// Pages represents the number of pages of the book
	Pages int
	// Embedded Publisher having info about the publisher
	Publisher
}
func (b Book) ShowInfo() (info string) {
	info = fmt.Sprintf(
		"Title: %s\nAuthor: %s\nPages: %d\nTotal publisher founders: %v\n",
		b.Title,
		b.Author,
		b.Pages,
		b.TotalFounders(),
	)
	return
}

// Movie is a struct that represents a movie
type Movie struct {
	// Title represents the title of the movie
	Title string
	// Director represents the director of the movie
	Director string
	// IMDBRating represents the IMDB rating of the movie
	IMDBRating float64
	// Embedded Publisher having info about the publisher
	Publisher
}
func (m Movie) ShowInfo() (info string) {
	info = fmt.Sprintf(
		"Title: %s\nDirector: %s\nIMDB Rating: %.1f\nTotal publisher founders: %v\n",
		m.Title,
		m.Director,
		m.IMDBRating,
		m.TotalFounders(),
	)
	return
}

func main() {
	// a movie where the publisher is warner bross pictures
	movie := Movie{
		Title:      "The Dark Knight",
		Director:   "Christopher Nolan",
		IMDBRating: 9.0,
		Publisher: Publisher{
			Name: "Warner Bross Pictures",
			Founders: []string{
				"Albert Warner",
				"Harry Warner",
				"Samuel Warner",
			},
		},
	}

	// a book where the publisher is paramount pictures
	book := Book{
		Title:  "Titanic",
		Author: "James Cameron",
		Pages:  560,
		Publisher: Publisher{
			Name: "Paramount Pictures",
			Founders: []string{
				"Adolph Zukor",
				"Jesse L. Lasky",
			},
		},
	}

	// print the info of the movie
	fmt.Println(movie.ShowInfo())

	// print the info of the book
	fmt.Println(book.ShowInfo())
}