package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"os"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/manager/signals"
	"ton/pkg/reconciler/line"
)

func main(){
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		fmt.Print(err)
	}
	err = line.Add(mgr)

	//controllerruntime.NewControllerManagedBy(mgr)
	//kubernetesInformer :=
	_ = line.Add(mgr)
	if err := mgr.Start(signals.SetupSignalHandler()); err != nil {
		fmt.Println("sasa")
		os.Exit(1)
	}
}

//func AddController(){
//	kubernetesClient, err :=
//}