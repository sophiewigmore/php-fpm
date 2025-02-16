package main

import (
	"os"

	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/chronos"
	"github.com/paketo-buildpacks/packit/v2/draft"
	"github.com/paketo-buildpacks/packit/v2/scribe"
	phpfpm "github.com/paketo-buildpacks/php-fpm"
)

func main() {
	logEmitter := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))
	config := phpfpm.NewConfig()
	entryResolver := draft.NewPlanner()
	packit.Run(phpfpm.Detect(), phpfpm.Build(entryResolver, config, chronos.DefaultClock, logEmitter))
}
