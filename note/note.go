// Package note implements a note, and managers how the data is managed
package note

import "errors"

// NewNote creates a note with a title, and optional metadata
func NewNote(title string, metadata Metadata) (*Note, error) {
	if title == "" {
		return nil, ErrEmptyTitle
	}

	return &Note{
		title:    title,
		metadata: metadata,
	}, nil
}

// Title returns the note's title
func (n *Note) Title() string {
	return n.title
}

// Metadata returns the note's metadata
func (n *Note) Metadata() Metadata {
	return n.metadata
}

// Content returns the note's content
func (n *Note) Content() string {
	return n.content
}

// Error messages
var (
	ErrEmptyTitle = errors.New("empty title is not allowed")
)
