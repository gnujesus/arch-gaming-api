package services

type GenericService[T any] interface {
	GetAll() ([]T, error)
	GetById(id string) (T, error)
	Save(T) error
	Delete(id string) error
}
