package repo

import (
	"errors"
	"log"
	"math/rand"
	"quoteList/internal/payload"
)

var ErrInvalidInput = errors.New("invalid input")

type Storage interface {
	Create()
	GetAll()
	GetRandom()
	Filter()
	Delete()
}

type Store struct {
	data   []payload.Quote
	lastID uint64
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) Create(data *payload.Quote) {
	s.lastID++
	data.Id = s.lastID
	s.data = append(s.data, *data)
	log.Println("Quote added successfully", data)

}

func (s *Store) GetAll() []payload.Quote {
	var quotes []payload.Quote
	if len(s.data) == 0 {
		log.Println("Cant get quotes from empty list")
		return quotes
	}

	for _, val := range s.data {
		log.Println(val)
		quotes = append(quotes, val)
	}

	return quotes
}

func (s *Store) GetRandom() (payload.Quote, error) {

	if len(s.data) == 0 {
		log.Println("Cant get random quote from empty list")
		return payload.Quote{}, ErrInvalidInput
	}

	randomInt := rand.Intn(len(s.data))
	log.Println("Random quote is:", s.data[randomInt])
	return s.data[randomInt], nil
}

func (s *Store) Filter(author string) []payload.Quote {
	if author == "" {

		return s.data
	}

	var filtered []payload.Quote
	for _, q := range s.data {
		if q.Author == author {
			filtered = append(filtered, q)
		}
	}
	return filtered
}

func (s *Store) Delete(id uint64) {
	for i := range s.data {
		if s.data[i].Id == id {

			s.data = append(s.data[:i], s.data[i+1:]...)
			log.Printf("Quote deleted successfully with id = %v\n", id)
			return

		}
	}
}
