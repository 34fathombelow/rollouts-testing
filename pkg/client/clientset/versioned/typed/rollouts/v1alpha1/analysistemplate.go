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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/argoproj/argo-rollouts/pkg/apis/rollouts/v1alpha1"
	scheme "github.com/argoproj/argo-rollouts/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AnalysisTemplatesGetter has a method to return a AnalysisTemplateInterface.
// A group's client should implement this interface.
type AnalysisTemplatesGetter interface {
	AnalysisTemplates(namespace string) AnalysisTemplateInterface
}

// AnalysisTemplateInterface has methods to work with AnalysisTemplate resources.
type AnalysisTemplateInterface interface {
	Create(ctx context.Context, analysisTemplate *v1alpha1.AnalysisTemplate, opts v1.CreateOptions) (*v1alpha1.AnalysisTemplate, error)
	Update(ctx context.Context, analysisTemplate *v1alpha1.AnalysisTemplate, opts v1.UpdateOptions) (*v1alpha1.AnalysisTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AnalysisTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AnalysisTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnalysisTemplate, err error)
	AnalysisTemplateExpansion
}

// analysisTemplates implements AnalysisTemplateInterface
type analysisTemplates struct {
	client rest.Interface
	ns     string
}

// newAnalysisTemplates returns a AnalysisTemplates
func newAnalysisTemplates(c *ArgoprojV1alpha1Client, namespace string) *analysisTemplates {
	return &analysisTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the analysisTemplate, and returns the corresponding analysisTemplate object, and an error if there is any.
func (c *analysisTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AnalysisTemplate, err error) {
	result = &v1alpha1.AnalysisTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("analysistemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AnalysisTemplates that match those selectors.
func (c *analysisTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AnalysisTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AnalysisTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("analysistemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested analysisTemplates.
func (c *analysisTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("analysistemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a analysisTemplate and creates it.  Returns the server's representation of the analysisTemplate, and an error, if there is any.
func (c *analysisTemplates) Create(ctx context.Context, analysisTemplate *v1alpha1.AnalysisTemplate, opts v1.CreateOptions) (result *v1alpha1.AnalysisTemplate, err error) {
	result = &v1alpha1.AnalysisTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("analysistemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(analysisTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a analysisTemplate and updates it. Returns the server's representation of the analysisTemplate, and an error, if there is any.
func (c *analysisTemplates) Update(ctx context.Context, analysisTemplate *v1alpha1.AnalysisTemplate, opts v1.UpdateOptions) (result *v1alpha1.AnalysisTemplate, err error) {
	result = &v1alpha1.AnalysisTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("analysistemplates").
		Name(analysisTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(analysisTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the analysisTemplate and deletes it. Returns an error if one occurs.
func (c *analysisTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("analysistemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *analysisTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("analysistemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched analysisTemplate.
func (c *analysisTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AnalysisTemplate, err error) {
	result = &v1alpha1.AnalysisTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("analysistemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
