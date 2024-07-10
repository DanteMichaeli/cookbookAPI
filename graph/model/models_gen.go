// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type Query struct {
}

type Recipe struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       *string  `json:"image,omitempty"`
	Ingredients []string `json:"ingredients"`
	Steps       []string `json:"steps"`
}
