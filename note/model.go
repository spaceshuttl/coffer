package note

// Metadata is an internal map of metadata such as timestamps, machine name, guid
type Metadata map[string]string

// Note holds data for notes and internal metadata
type Note struct {
	title    string
	metadata Metadata

	content string
}
