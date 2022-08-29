package main

import (
	"context"
	"os"
	"os/signal"
	"sort"

	"github.com/ivanmeca/emptyApp/application"
	"github.com/urfave/cli"
)

func runApplication(cli *cli.Context) error {
	c := context.Background()
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	appMan := application.NewApp()

	err := appMan.Init(ctx)
	if err != nil {
		return err
	}

	err = appMan.Run(ctx)
	if err != nil {
		return err
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	return nil
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{},
	}
	app.Version = Version + "(" + GitCommit + ")"
	app.Name = ApplicationName
	app.Usage = ""
	app.Description = ""
	app.Copyright = ""
	app.EnableBashCompletion = true
	app.Action = runApplication
	app.Commands = []cli.Command{}
	sort.Sort(cli.FlagsByName(app.Flags))
	err := app.Run(os.Args)
	if err != nil {
		panic(err.Error())
	}
}
