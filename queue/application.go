package queue

import (
	configcontract "github.com/goravel/framework/contracts/config"
	"github.com/goravel/framework/contracts/log"
	"github.com/goravel/framework/contracts/queue"
)

type Application struct {
	config *Config
	jobs   []queue.Job
	log    log.Log
}

func NewApplication(config configcontract.Config, log log.Log) *Application {
	return &Application{
		config: NewConfig(config),
		log:    log,
	}
}

func (app *Application) Worker(args *queue.Args) queue.Worker {
	defaultConnection := app.config.DefaultConnection()

	if args == nil {
		return NewWorker(app.config, app.log, 1, defaultConnection, app.jobs, app.config.Queue(defaultConnection, ""))
	}
	if args.Connection == "" {
		args.Connection = defaultConnection
	}

	return NewWorker(app.config, app.log, args.Concurrent, args.Connection, app.jobs, app.config.Queue(args.Connection, args.Queue))
}

func (app *Application) Register(jobs []queue.Job) {
	app.jobs = append(app.jobs, jobs...)
}

func (app *Application) GetJobs() []queue.Job {
	return app.jobs
}

func (app *Application) Job(job queue.Job, args []queue.Arg) queue.Task {
	return NewTask(app.config, app.log, job, args)
}

func (app *Application) Chain(jobs []queue.Jobs) queue.Task {
	return NewChainTask(app.config, app.log, jobs)
}
