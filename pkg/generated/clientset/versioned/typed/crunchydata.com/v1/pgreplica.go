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
	"time"

	v1 "github.com/crunchydata/postgres-operator/apis/crunchydata.com/v1"
	scheme "github.com/crunchydata/postgres-operator/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PgreplicasGetter has a method to return a PgreplicaInterface.
// A group's client should implement this interface.
type PgreplicasGetter interface {
	Pgreplicas(namespace string) PgreplicaInterface
}

// PgreplicaInterface has methods to work with Pgreplica resources.
type PgreplicaInterface interface {
	Create(*v1.Pgreplica) (*v1.Pgreplica, error)
	Update(*v1.Pgreplica) (*v1.Pgreplica, error)
	UpdateStatus(*v1.Pgreplica) (*v1.Pgreplica, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.Pgreplica, error)
	List(opts metav1.ListOptions) (*v1.PgreplicaList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Pgreplica, err error)
	PgreplicaExpansion
}

// pgreplicas implements PgreplicaInterface
type pgreplicas struct {
	client rest.Interface
	ns     string
}

// newPgreplicas returns a Pgreplicas
func newPgreplicas(c *CrunchydataV1Client, namespace string) *pgreplicas {
	return &pgreplicas{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the pgreplica, and returns the corresponding pgreplica object, and an error if there is any.
func (c *pgreplicas) Get(name string, options metav1.GetOptions) (result *v1.Pgreplica, err error) {
	result = &v1.Pgreplica{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pgreplicas").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Pgreplicas that match those selectors.
func (c *pgreplicas) List(opts metav1.ListOptions) (result *v1.PgreplicaList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PgreplicaList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("pgreplicas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested pgreplicas.
func (c *pgreplicas) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("pgreplicas").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a pgreplica and creates it.  Returns the server's representation of the pgreplica, and an error, if there is any.
func (c *pgreplicas) Create(pgreplica *v1.Pgreplica) (result *v1.Pgreplica, err error) {
	result = &v1.Pgreplica{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("pgreplicas").
		Body(pgreplica).
		Do().
		Into(result)
	return
}

// Update takes the representation of a pgreplica and updates it. Returns the server's representation of the pgreplica, and an error, if there is any.
func (c *pgreplicas) Update(pgreplica *v1.Pgreplica) (result *v1.Pgreplica, err error) {
	result = &v1.Pgreplica{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pgreplicas").
		Name(pgreplica.Name).
		Body(pgreplica).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *pgreplicas) UpdateStatus(pgreplica *v1.Pgreplica) (result *v1.Pgreplica, err error) {
	result = &v1.Pgreplica{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("pgreplicas").
		Name(pgreplica.Name).
		SubResource("status").
		Body(pgreplica).
		Do().
		Into(result)
	return
}

// Delete takes name of the pgreplica and deletes it. Returns an error if one occurs.
func (c *pgreplicas) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pgreplicas").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *pgreplicas) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("pgreplicas").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched pgreplica.
func (c *pgreplicas) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Pgreplica, err error) {
	result = &v1.Pgreplica{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("pgreplicas").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}