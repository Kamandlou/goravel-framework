package console

import (
	"github.com/goravel/framework/console/console"
	consolecontract "github.com/goravel/framework/contracts/console"
	"github.com/goravel/framework/contracts/foundation"
)

const Binding = "goravel.console"

type ServiceProvider struct {
}

func (receiver *ServiceProvider) Register(app foundation.Application) {
	app.Singleton(Binding, func(app foundation.Application) (any, error) {
		name := "Goravel Framework"
		usage := app.Version()
		usageText := "artisan [global options] command [options] [arguments...]"
		return NewApplication(name, usage, usageText, app.Version(), true), nil
	})
}

func (receiver *ServiceProvider) Boot(app foundation.Application) {
	receiver.registerCommands(app)
}

func (receiver *ServiceProvider) registerCommands(app foundation.Application) {
	artisan := app.MakeArtisan()
	config := app.MakeConfig()
	artisan.Register([]consolecontract.Command{
		console.NewListCommand(artisan),
		console.NewKeyGenerateCommand(config),
		console.NewMakeCommand(),
		console.NewBuildCommand(config),
	})
}
