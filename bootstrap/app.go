package bootstrap

import (
	"github.com/ryokushaka/project_YoriEat-gin-deployment-repo/postgres"
)

type Application struct {
	Env *Env
	Postgres *postgres.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgres)
}