package basemodel

import "time"

// Base базовая структура новых моделей
type Base struct {
	CreateAt time.Time
	UpdateAt time.Time
}

// BaseUpdate базовая структура моделей для обновления
type BaseUpdate struct {
	UpdateAt time.Time
}

// NewBase новая базовая структура
func NewBase() Base {
	return Base{
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

// NewBaseUpdate новая базовая структура для обновления
func NewBaseUpdate() BaseUpdate {
	return BaseUpdate{
		UpdateAt: time.Now(),
	}
}

// UpdateField шаблон для обновлений
type UpdateField[T any] struct {
	NeedUpdate bool
	Value      T
}

// NewUpdateField новый шаблон для обновлений
func NewUpdateField[T any](value T) UpdateField[T] {
	return UpdateField[T]{
		NeedUpdate: true,
		Value:      value,
	}
}
