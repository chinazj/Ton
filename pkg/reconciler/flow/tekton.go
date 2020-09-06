package flow

import (
	pipeline "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	resource "github.com/tektoncd/pipeline/pkg/apis/resource/v1alpha1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"ton/pkg/apis/ton/v1alpha1"
	flow "ton/pkg/apis/ton/v1alpha1"
)

func CreateTektonPipelineRun(task pipeline.PipelineTaskList, flowname, namespace string) *pipeline.PipelineRun {
	pipelineTask := &pipeline.PipelineRun{
		ObjectMeta:metav1.ObjectMeta{
			Name: flowname+"-pipelineRun",
			Namespace:namespace,
		},
		Spec:pipeline.PipelineRunSpec{
			PipelineSpec:&pipeline.PipelineSpec{
				Tasks: task,
			},
		},
	}
	return pipelineTask
}

func CreateTektonStepSpec(task v1alpha1.Task, flowname, namespace string) pipeline.PipelineTask{
	tektonpipeline := pipeline.PipelineTask{
		Name:task.Name,
		TaskSpec:&pipeline.TaskSpec{
			Steps: []pipeline.Step{
				{},
			},
		},
	}
	return tektonpipeline
}

func CreateTektonTaskRunSpec(task v1alpha1.Task, flowname, namespace string) pipeline.PipelineTask{
	tektonpipeline := pipeline.PipelineTask{
		Name:task.Name,
		TaskSpec:&pipeline.TaskSpec{
			Steps: []pipeline.Step{
				{
					Container:v1.Container{
						Name:task.Name+"step",
						Image:task.Image,
						Command:[]string{task.Commond},
					},
				},
			},
		},
	}
	return tektonpipeline
}

func CreateTektonResource(flowresources flow.Resources, flowname , namespace string) *resource.PipelineResource{
	tektonResource := &resource.PipelineResource{
		ObjectMeta: metav1.ObjectMeta{
			Name:      flowresources.Name+"-"+flowname+"-pipelineResource",
			Namespace: namespace,
		},
		Spec: resource.PipelineResourceSpec{
			Type: flowresources.ResourcesType,
			Params: flowresources.Params,
		},
	}
	return tektonResource
}
