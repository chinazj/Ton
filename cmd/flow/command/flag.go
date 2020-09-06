package command

import (
	"flag"
	_ "net/http/pprof"
	"os"
)

type Commond struct {
	kubeconfig *string
	debug      *bool
	help       *bool
}

func NewCommond() *Commond{
	c := &Commond{}
	// https://github.com/kubernetes-sigs/controller-runtime/issues/878
	// Do not set `kubeconfig` and `master` parameters. The other has this parameter by default
	c.kubeconfig = flag.String("k","~/.kube/config","kubeconfig")
	c.debug      = flag.Bool("d",false,"debug")
	c.help       = flag.Bool("h",false,"help")
	flag.Parse()
	return c
}

func (c *Commond)Parse(){
	if *c.help {
		flag.Usage()
		os.Exit(1)
	}
}
