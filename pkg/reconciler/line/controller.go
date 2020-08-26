package line

import (
	"fmt"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
	"ton/pkg/client/listers/ton/v1alpha1"
	tontype "ton/pkg/apis/ton/v1alpha1"
)

type Reconciler struct {
	KubeClientSet     kubernetes.Interface
	PipelineClientSet clientset.Interface

	Client client.Client
	Scheme *runtime.Scheme
	// listers index properties about resources
	line v1alpha1.LineLister

}

// Add creates a new Workspace Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &Reconciler{}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("ton-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to
	err = c.Watch(&source.Kind{Type: &tontype.Line{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &Reconciler{}

func (c *Reconciler) Reconcile(request reconcile.Request)(reconcile.Result, error){
	fmt.Println("reconcile")
	return reconcile.Result{}, nil
}