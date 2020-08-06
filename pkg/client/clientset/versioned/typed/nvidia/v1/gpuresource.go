/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/k8s-gpushare/gpushare-crd-controller/pkg/apis/nvidia/v1"
	scheme "github.com/k8s-gpushare/gpushare-crd-controller/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GPUResourcesGetter has a method to return a GPUResourceInterface.
// A group's client should implement this interface.
type GPUResourcesGetter interface {
	GPUResources(namespace string) GPUResourceInterface
}

// GPUResourceInterface has methods to work with GPUResource resources.
type GPUResourceInterface interface {
	Create(ctx context.Context, gPUResource *v1.GPUResource, opts metav1.CreateOptions) (*v1.GPUResource, error)
	Update(ctx context.Context, gPUResource *v1.GPUResource, opts metav1.UpdateOptions) (*v1.GPUResource, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GPUResource, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GPUResourceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GPUResource, err error)
	GPUResourceExpansion
}

// gPUResources implements GPUResourceInterface
type gPUResources struct {
	client rest.Interface
	ns     string
}

// newGPUResources returns a GPUResources
func newGPUResources(c *NvidiaV1Client, namespace string) *gPUResources {
	return &gPUResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gPUResource, and returns the corresponding gPUResource object, and an error if there is any.
func (c *gPUResources) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GPUResource, err error) {
	result = &v1.GPUResource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gpuresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GPUResources that match those selectors.
func (c *gPUResources) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GPUResourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GPUResourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gpuresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gPUResources.
func (c *gPUResources) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gpuresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a gPUResource and creates it.  Returns the server's representation of the gPUResource, and an error, if there is any.
func (c *gPUResources) Create(ctx context.Context, gPUResource *v1.GPUResource, opts metav1.CreateOptions) (result *v1.GPUResource, err error) {
	result = &v1.GPUResource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gpuresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gPUResource).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a gPUResource and updates it. Returns the server's representation of the gPUResource, and an error, if there is any.
func (c *gPUResources) Update(ctx context.Context, gPUResource *v1.GPUResource, opts metav1.UpdateOptions) (result *v1.GPUResource, err error) {
	result = &v1.GPUResource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gpuresources").
		Name(gPUResource.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(gPUResource).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the gPUResource and deletes it. Returns an error if one occurs.
func (c *gPUResources) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gpuresources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gPUResources) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gpuresources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched gPUResource.
func (c *gPUResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GPUResource, err error) {
	result = &v1.GPUResource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gpuresources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}