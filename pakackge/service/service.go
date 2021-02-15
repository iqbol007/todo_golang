package service

import "todoAppGo/pakackge/reposithory"

type Authorization interface {
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization Authorization
	TodoItem      TodoItem
	TodoList      TodoList
}

func NewService(repos *reposithory.Repository) *Service {
	return &Service{

	}

}
