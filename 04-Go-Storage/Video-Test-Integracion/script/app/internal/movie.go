package internal

// MovieAttributes is an struct that represents the attributes of a movie
type MovieAttributes struct {
	// Title is the title of the movie
	Title string
	// Year is the year of the movie
	Year int
	// Director is the director of the movie
	Director string
}

// Movie is an struct that represents a movie
type Movie struct {
	// Id is the unique identifier of the movie
	Id int
	// Attributes is the attributes of the movie
	Attributes MovieAttributes
}