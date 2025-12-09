// Package domain содержит доменные модели и бизнес-сущности.
package domain

// Metric представляет метрику с идентификатором, именем, значением и типом.
type Metric struct {
	ID    string
	Name  string
	Value float64
	Type  string
}
