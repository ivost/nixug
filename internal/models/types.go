package models

type ReadParam struct {
	Start       string
	End         string
	Limit       int64
	Stream      bool
	ContentType string
}
