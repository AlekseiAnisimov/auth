package auth

import dbx "github.com/go-ozzo/ozzo-dbx"

//Env структура хранящая данные окружения
type Env struct {
	db *dbx.DB
}

// GetEnvDbPointer метод полученияуказателя к ORM ozzo-dbx
func (env *Env) GetEnvDbPointer() *dbx.DB {
	return env.db
}
