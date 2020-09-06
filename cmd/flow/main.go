package main

import (
	"fmt"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"ton/cmd/flow/command"
	flowScheme "ton/pkg/apis/ton/v1alpha1"
	"ton/pkg/reconciler/flow"
)

var (

)

func main(){
	c := command.NewCommond()
	c.Parse()

	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		fmt.Print(err)
	}
	// add more Scheme
	err = flowScheme.AddToScheme(mgr.GetScheme())
	if err != nil {
		fmt.Print(err)
	}

	// add more controller
	err = flow.NewControllerManagerBy(mgr)
	if err != nil {
		fmt.Print(err)
	}

	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}