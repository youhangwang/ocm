// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "github.com/openshift/api/config/v1"
	configv1 "github.com/openshift/client-go/config/applyconfigurations/config/v1"
	scheme "github.com/openshift/client-go/config/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SchedulersGetter has a method to return a SchedulerInterface.
// A group's client should implement this interface.
type SchedulersGetter interface {
	Schedulers() SchedulerInterface
}

// SchedulerInterface has methods to work with Scheduler resources.
type SchedulerInterface interface {
	Create(ctx context.Context, scheduler *v1.Scheduler, opts metav1.CreateOptions) (*v1.Scheduler, error)
	Update(ctx context.Context, scheduler *v1.Scheduler, opts metav1.UpdateOptions) (*v1.Scheduler, error)
	UpdateStatus(ctx context.Context, scheduler *v1.Scheduler, opts metav1.UpdateOptions) (*v1.Scheduler, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Scheduler, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SchedulerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Scheduler, err error)
	Apply(ctx context.Context, scheduler *configv1.SchedulerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Scheduler, err error)
	ApplyStatus(ctx context.Context, scheduler *configv1.SchedulerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Scheduler, err error)
	SchedulerExpansion
}

// schedulers implements SchedulerInterface
type schedulers struct {
	client rest.Interface
}

// newSchedulers returns a Schedulers
func newSchedulers(c *ConfigV1Client) *schedulers {
	return &schedulers{
		client: c.RESTClient(),
	}
}

// Get takes name of the scheduler, and returns the corresponding scheduler object, and an error if there is any.
func (c *schedulers) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Get().
		Resource("schedulers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Schedulers that match those selectors.
func (c *schedulers) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SchedulerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SchedulerList{}
	err = c.client.Get().
		Resource("schedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested schedulers.
func (c *schedulers) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("schedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a scheduler and creates it.  Returns the server's representation of the scheduler, and an error, if there is any.
func (c *schedulers) Create(ctx context.Context, scheduler *v1.Scheduler, opts metav1.CreateOptions) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Post().
		Resource("schedulers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduler).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a scheduler and updates it. Returns the server's representation of the scheduler, and an error, if there is any.
func (c *schedulers) Update(ctx context.Context, scheduler *v1.Scheduler, opts metav1.UpdateOptions) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Put().
		Resource("schedulers").
		Name(scheduler.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduler).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *schedulers) UpdateStatus(ctx context.Context, scheduler *v1.Scheduler, opts metav1.UpdateOptions) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Put().
		Resource("schedulers").
		Name(scheduler.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scheduler).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the scheduler and deletes it. Returns an error if one occurs.
func (c *schedulers) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("schedulers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *schedulers) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("schedulers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched scheduler.
func (c *schedulers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Scheduler, err error) {
	result = &v1.Scheduler{}
	err = c.client.Patch(pt).
		Resource("schedulers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied scheduler.
func (c *schedulers) Apply(ctx context.Context, scheduler *configv1.SchedulerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Scheduler, err error) {
	if scheduler == nil {
		return nil, fmt.Errorf("scheduler provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(scheduler)
	if err != nil {
		return nil, err
	}
	name := scheduler.Name
	if name == nil {
		return nil, fmt.Errorf("scheduler.Name must be provided to Apply")
	}
	result = &v1.Scheduler{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("schedulers").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *schedulers) ApplyStatus(ctx context.Context, scheduler *configv1.SchedulerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.Scheduler, err error) {
	if scheduler == nil {
		return nil, fmt.Errorf("scheduler provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(scheduler)
	if err != nil {
		return nil, err
	}

	name := scheduler.Name
	if name == nil {
		return nil, fmt.Errorf("scheduler.Name must be provided to Apply")
	}

	result = &v1.Scheduler{}
	err = c.client.Patch(types.ApplyPatchType).
		Resource("schedulers").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
