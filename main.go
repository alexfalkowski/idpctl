package main

import (
	"os"

	sc "github.com/alexfalkowski/go-service/cmd"
	"github.com/alexfalkowski/go-service/flags"
	"github.com/alexfalkowski/idpctl/cmd"
	"github.com/alexfalkowski/idpctl/cmd/pipeline"
)

func main() {
	if err := command().Run(); err != nil {
		os.Exit(1)
	}
}

type fn func(*sc.Command)

func command() *sc.Command {
	c := sc.New(cmd.Version)
	c.RegisterInput(c.Root(), "")

	// Add more commands here to extend the cli.
	fns := []fn{pipelineCommand}
	for _, f := range fns {
		f(c)
	}

	return c
}

func pipelineCommand(c *sc.Command) {
	r := c.AddClientCommand("pipeline", "Manage pipelines.", cmd.Module, pipeline.Module)
	flags.StringVar(r, pipeline.CreateFlag, "create", "c", "", "create a pipeline (path to file)")
	flags.StringVar(r, pipeline.GetFlag, "get", "g", "", "retrieve a pipeline (the id of the pipeline)")
	flags.StringVar(r, pipeline.UpdateFlag, "update", "u", "", "update a pipeline (id of the pipeline and the path to file, e.g. id:path)")
	flags.StringVar(r, pipeline.DeleteFlag, "delete", "d", "", "delete a pipeline (the id of the pipeline)")
	flags.StringVar(r, pipeline.TriggerFlag, "trigger", "t", "", "trigger a pipeline (the id of the pipeline)")
}
