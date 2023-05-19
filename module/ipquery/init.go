package ipquery

import (
	"github.com/fimreal/rack/module"
)

const (
	ID          = "ipquery"
	Comment     = "ip query tools"
	RoutePrefix = "/ipquery"
)

var Module = module.Module{
	ID:      ID,
	Comment: Comment,
	// gin route
	RouteFunc:   AddRoute,
	RoutePrefix: RoutePrefix,
	// cobra flag
	FlagFunc: ServeFlag,
}
