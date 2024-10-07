package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

type Application struct {
	Env    *Env
	DB     *gorm.DB
	SqlxDB *sqlx.DB
}

func App() (*Application, error) {
	env, err := NewEnv()
	if err != nil {
		return nil, err
	}

	db, sqlxDB, err := NewPostgresDatabase()
	if err != nil {
		return nil, err
	}

	return &Application{
		Env:    env,
		DB:     db,
		SqlxDB: sqlxDB,
	}, nil
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection()
}
