package reposithory


type Authorization interface {

}

type TodoList interface {

}
type TodoItem interface {

}
type Repository struct {
	Authorization Authorization
	TodoItem TodoItem
	TodoList TodoList
}

func NewRepository() *Repository {
	return &Repository{

	}

}
