package module

import (
	"regexp"
	"strings"

	"github.com/fimreal/goutils/ezap"
	"github.com/fimreal/rack/pkg/components/crond"
	"github.com/fimreal/rack/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

var (
	ModVersion   = []string{}
	RouteFuncs   []func(*gin.Engine)
	FlagFuncs    []func(*cobra.Command)
	CliFlagFuncs []func(*cobra.Command)
	CrondFuncs   map[string]func() = make(map[string]func())
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
	CrondFunc   map[string]func()    // crond CrondFunc["* * * * *"] = func() { log.Print "do sth at ts" }()
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

func RunCron() {
	for spec, f := range CrondFuncs {
		if f != nil {
			job := f // 创建一个新的变量，确保每次循环的 `f` 是独立的
			id, err := crond.Run(spec, job)
			if err != nil {
				ezap.Error(err.Error())
			}
			jobName := utils.GetFunctionName(job)
			ezap.Infof("Add cronjob: %s %s %s", id, spec, jobName)
			if strings.HasPrefix(spec, "@") {
				ezap.Infof("Firstly starting cronjob: %s %s", spec, jobName)
				go func() { job() }()
			}
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
	RouteFuncs = append(RouteFuncs, m.RouteFunc)
	ModVersion = append(ModVersion, m.ID)
	FlagFuncs = append(FlagFuncs, m.FlagFunc)
	CliFlagFuncs = append(CliFlagFuncs, m.CliFlagFunc)
	if len(m.CrondFunc) != 0 {
		for spec, f := range m.CrondFunc {
			re := regexp.MustCompile(`^\[.*?\] +`)
			spec = re.ReplaceAllString(spec, "")
			CrondFuncs[spec] = f
		}
	}
}
