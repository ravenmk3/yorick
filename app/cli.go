package app

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	AppName = "yorick"
)

func RunCliApp() error {
	app := NewCliApp()
	return app.Run(os.Args)
}

func NewCliApp() *cli.App {
	app := &cli.App{
		Name:        AppName,
		Usage:       AppName,
		Description: "Yorick backup tool",
		Commands: []*cli.Command{
			NewRunCommand(),
		},
	}
	return app
}

func NewRunCommand() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "Run a backup script",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "debug", Required: false, Value: false},
			&cli.StringFlag{Name: "script", Aliases: []string{"s"}, Required: false, Value: ".yorick.js"},
			&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Required: false, Value: ".backup"},
		},
		Action: func(ctx *cli.Context) error {
			debug := ctx.Bool("debug")
			if debug {
				logrus.SetLevel(logrus.DebugLevel)
			}
			scriptFile := ctx.String("script")
			outputDir := ctx.String("output")
			return ExecRunScript(scriptFile, outputDir)
		},
	}
}
