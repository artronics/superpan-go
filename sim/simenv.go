package sim

import (
	"github.com/artronics/superpan/graph"
	"os"
)

var (
	envfile string
)
const(
	defaultEnvFile ="./env.json"
)

type simenv struct {
	graph graph.Graph
}

func envFile(p string)(env *simenv,err error)  {
	fi,err:=os.Open(p)
	defer fi.Close()

	env=&simenv{}

	return env,err
}

