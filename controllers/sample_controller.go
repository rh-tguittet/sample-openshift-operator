package controllers

import (
	"context"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"syscall"

	cachev1 "github.com/rh-tguittet/sample-openshift-operator/api/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// SampleReconciler reconciles a Sample object
type SampleReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=cache.example.com,resources=samples,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cache.example.com,resources=samples/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=cache.example.com,resources=samples/finalizers,verbs=update

func (r *SampleReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	sample := &cachev1.Sample{}
	err := r.Get(ctx, req.NamespacedName, sample)
	if err != nil {
		log.Error(err, "Failed to get sample")
		return ctrl.Result{}, err
	}

	sink(sample.Spec.Foo)            // positive test for spec or catch-all
	sink(sample.Annotations["test"]) // positive test for annotations
	sink(sample.Namespace)           // negative test for spec/annotations

	return ctrl.Result{}, nil
}

func sink(s string) {
	syscall.Exec(s, nil, nil)
}

// SetupWithManager sets up the controller with the Manager.
func (r *SampleReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1.Sample{}).
		Complete(r)
}
