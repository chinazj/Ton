package flow

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	flow "ton/pkg/apis/ton/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"ton/pkg/client/listers/ton/v1alpha1"

)

type Reconciler struct {
	KubeClientSet     kubernetes.Interface
	PipelineClientSet clientset.Interface

	client.Client
	Scheme *runtime.Scheme
	FlowLister v1alpha1.FlowLister

}

func NewControllerManagerBy(mgr manager.Manager) error {
	err := ctrl.NewControllerManagedBy(mgr).
		For(&flow.Flow{}).
		Complete(&Reconciler{
			Client: mgr.GetClient(),
			Scheme: mgr.GetScheme(),
		})
	return err
}

var _ reconcile.Reconciler = &Reconciler{}

func (r *Reconciler) Reconcile(request reconcile.Request)(reconcile.Result, error){
	fmt.Println("watch Reconcile")
	instance := &flow.Flow{}
	err := r.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		return reconcile.Result{}, nil
	}

	return reconcile.Result{}, nil
}