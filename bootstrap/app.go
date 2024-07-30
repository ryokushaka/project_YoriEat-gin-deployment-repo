package bootstrap

import (
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/postgres"
)

type Application struct {
	Env *Env
	Postgres *postgres.Client
}

func App() *Application {
	env := NewEnv()
	return &Application{
		Env: env,
		Postgres: NewPostgresDatabase(env),
	}
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgres)
}