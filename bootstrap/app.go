package bootstrap

import (
	"github.com/jmoiron/sqlx"
	"gorm.io/gorm"
)

// Application holds the environment and database connections for the application.
type Application struct {
	Env    *Env
	DB     *gorm.DB
	SqlxDB *sqlx.DB
}

// App initializes and returns a new Application instance.
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

// CloseDBConnection closes the database connections.
func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection()
}
