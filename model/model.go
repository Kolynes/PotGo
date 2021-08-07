package model

type Model struct {
	objects ModelManager
}

type IModel interface {
	Save()
	Delete()
}
