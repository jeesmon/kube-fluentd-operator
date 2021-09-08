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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/vmware/kube-fluentd-operator/config-reloader/datasource/kubedatasource/fluentdconfig/apis/logs.vdp.vmware.com/v1beta1"
	scheme "github.com/vmware/kube-fluentd-operator/config-reloader/datasource/kubedatasource/fluentdconfig/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FluentdConfigsGetter has a method to return a FluentdConfigInterface.
// A group's client should implement this interface.
type FluentdConfigsGetter interface {
	FluentdConfigs(namespace string) FluentdConfigInterface
}

// FluentdConfigInterface has methods to work with FluentdConfig resources.
type FluentdConfigInterface interface {
	Create(*v1beta1.FluentdConfig) (*v1beta1.FluentdConfig, error)
	Update(*v1beta1.FluentdConfig) (*v1beta1.FluentdConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.FluentdConfig, error)
	List(opts v1.ListOptions) (*v1beta1.FluentdConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FluentdConfig, err error)
	FluentdConfigExpansion
}

// fluentdConfigs implements FluentdConfigInterface
type fluentdConfigs struct {
	client rest.Interface
	ns     string
}

// newFluentdConfigs returns a FluentdConfigs
func newFluentdConfigs(c *LogsV1beta1Client, namespace string) *fluentdConfigs {
	return &fluentdConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the fluentdConfig, and returns the corresponding fluentdConfig object, and an error if there is any.
func (c *fluentdConfigs) Get(name string, options v1.GetOptions) (result *v1beta1.FluentdConfig, err error) {
	result = &v1beta1.FluentdConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentdconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.TODO()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FluentdConfigs that match those selectors.
func (c *fluentdConfigs) List(opts v1.ListOptions) (result *v1beta1.FluentdConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FluentdConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("fluentdconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.TODO()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested fluentdConfigs.
func (c *fluentdConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("fluentdconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.TODO())
}

// Create takes the representation of a fluentdConfig and creates it.  Returns the server's representation of the fluentdConfig, and an error, if there is any.
func (c *fluentdConfigs) Create(fluentdConfig *v1beta1.FluentdConfig) (result *v1beta1.FluentdConfig, err error) {
	result = &v1beta1.FluentdConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("fluentdconfigs").
		Body(fluentdConfig).
		Do(context.TODO()).
		Into(result)
	return
}

// Update takes the representation of a fluentdConfig and updates it. Returns the server's representation of the fluentdConfig, and an error, if there is any.
func (c *fluentdConfigs) Update(fluentdConfig *v1beta1.FluentdConfig) (result *v1beta1.FluentdConfig, err error) {
	result = &v1beta1.FluentdConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("fluentdconfigs").
		Name(fluentdConfig.Name).
		Body(fluentdConfig).
		Do(context.TODO()).
		Into(result)
	return
}

// Delete takes name of the fluentdConfig and deletes it. Returns an error if one occurs.
func (c *fluentdConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentdconfigs").
		Name(name).
		Body(options).
		Do(context.TODO()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *fluentdConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("fluentdconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.TODO()).
		Error()
}

// Patch applies the patch and returns the patched fluentdConfig.
func (c *fluentdConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.FluentdConfig, err error) {
	result = &v1beta1.FluentdConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("fluentdconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(context.TODO()).
		Into(result)
	return
}
