module ton

go 1.13

replace (
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190920225731-5eefd052ad72
	k8s.io/apimachinery => k8s.io/apimachinery v0.17.6
	k8s.io/client-go => k8s.io/client-go v0.17.6
	sigs.k8s.io/controller-runtime => sigs.k8s.io/controller-runtime v0.5.0
)

require (
	github.com/tektoncd/pipeline v0.15.2
	k8s.io/api v0.18.6
	k8s.io/apimachinery v0.18.6
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	knative.dev/pkg v0.0.0-20200702222342-ea4d6e985ba0
	sigs.k8s.io/controller-runtime v0.5.0
)
