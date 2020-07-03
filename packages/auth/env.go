package auth

import dbx "github.com/go-ozzo/ozzo-dbx"

//Env структура хранящая данные окружения
type Env struct {
	db *dbx.DB
}

// GetEnvDbPointer метод получения указателя к ORM ozzo-dbx
func (env *Env) GetEnvDbPointer() *dbx.DB {
	return env.db
}

// SetEnvDbPointer установка укзателя
func (env *Env) SetEnvDbPointer(dbPointer *dbx.DB) {
	env.db = dbPointer
	return
}
