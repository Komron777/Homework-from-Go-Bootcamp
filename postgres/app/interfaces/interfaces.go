package funcs

import (
	f "app/funcs"
)

type methods interface {
	GetByID(int) f.User
	Get() []f.User
	Update(f.User) bool
	Delete(int) error
	Insert(f.User) bool
}

func INSERT(m methods, usr f.User) bool {
	return m.Insert(usr)
}
func DELETE(m methods, id int) error {
	return m.Delete(id)
}
func UPDATE(m methods, usr f.User) bool {
	return m.Update(usr)
}
func GET(m methods) []f.User {
	return m.Get()
}
func GETBYID(m methods,id int) f.User {
	return m.GetByID(id)
}
