package flow

import (

	"fmt"
	pipeline "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	clientset "k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/runtime/log"
	flow "ton/pkg/apis/ton/v1alpha1"
	"ton/pkg/client/listers/ton/v1alpha1"
	"ton/pkg/util/constants"
	"ton/pkg/util/sliceutil"
)

const (
	FlowFinalizerName = "flow.finalizers.k8s.io"
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
	err := r.Get(constants.RootContext, request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	if instance.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object.
		if !sliceutil.ContainsString(instance.ObjectMeta.Finalizers, FlowFinalizerName, nil) {
			instance.ObjectMeta.Finalizers = append(instance.ObjectMeta.Finalizers, FlowFinalizerName)
			if err := r.Update(constants.RootContext, instance); err != nil {
				return reconcile.Result{}, err
			}
		}
	} else {
		if sliceutil.ContainsString(instance.ObjectMeta.Finalizers, FlowFinalizerName, nil) {
			// If other resources depend on flow, please delete it
			instance.ObjectMeta.Finalizers = sliceutil.RemoveString(instance.ObjectMeta.Finalizers, FlowFinalizerName, nil)
			if err := r.Update(constants.RootContext, instance); err != nil {
				return reconcile.Result{}, err
			}
		}
		return reconcile.Result{}, nil
	}

	//Resources
	if instance.Spec.Resources != nil{
		for _,v := range instance.Spec.Resources{
			resource := CreateTektonResource(v,instance.Name,instance.Namespace)
			err = r.Create(constants.RootContext, resource)
		}
	}

	//task
	pipelinetasks := make(pipeline.PipelineTaskList,0)
	if instance.Spec.Tasks == nil || len(instance.Spec.Tasks) == 0{
		log.Log.WithName("flow-controller").Error(err, "have no task")
	}

	for _,v := range instance.Spec.Tasks{
		pipelinetasks = append(pipelinetasks, CreateTektonTaskRunSpec(v,instance.Name,instance.Namespace))
	}
	resource := CreateTektonPipelineRun(pipelinetasks,instance.Name,instance.Namespace)
	err = r.Create(constants.RootContext, resource)

	return reconcile.Result{}, nil
}