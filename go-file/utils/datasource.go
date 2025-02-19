package utils

type Datasource[T any] interface {
	Read() ([]T, error)
	Write(data []T) error
}