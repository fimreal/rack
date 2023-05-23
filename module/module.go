package module

import (
	"github.com/fimreal/goutils/ezap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	ModVersion   = []string{}
	RouteFuncs   []func(*gin.Engine)
	FlagFuncs    []func(*cobra.Command)
	CliFlagFuncs []func(*cobra.Command)
)

type ModuleInterface interface {
	Apply()
	Name() string
	Info() string
}

var _ ModuleInterface = &Module{}

type Module struct {
	ID          string               // name
	Comment     string               // usage
	RouteFunc   func(*gin.Engine)    // add route
	RoutePrefix string               // route
	FlagFunc    func(*cobra.Command) // serve flag
	CliFlagFunc func(*cobra.Command) // flag
}

func Register(modules []*Module) {
	for _, module := range modules {
		module.Apply()
	}
}

func GinLoad(r *gin.Engine) {
	for _, f := range RouteFuncs {
		f(r)
	}
}

func FlagParse(serveCmd *cobra.Command) {
	for _, f := range FlagFuncs {
		if f != nil {
			f(serveCmd)
		}
	}
}

func NewFlag(rootCmd *cobra.Command) {
	for _, f := range CliFlagFuncs {
		if f != nil {
			f(rootCmd)
		}
	}
}

func (m *Module) Name() string {
	return m.ID
}

func (m *Module) Info() string {
	return m.Comment
}

func (m *Module) RouterGroup() string {
	return m.RoutePrefix
}

func (m *Module) Apply() {
	ezap.Info("[module] Load " + m.ID)
	RouteFuncs = append(RouteFuncs, m.RouteFunc)
	ModVersion = append(ModVersion, m.ID)
	FlagFuncs = append(FlagFuncs, m.FlagFunc)
	CliFlagFuncs = append(CliFlagFuncs, m.CliFlagFunc)
}
