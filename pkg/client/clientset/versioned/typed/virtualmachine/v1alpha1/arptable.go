/*
Copyright 2018 Rancher Labs, Inc.

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
package v1alpha1

import (
	v1alpha1 "github.com/rancher/vm/pkg/apis/ranchervm/v1alpha1"
	scheme "github.com/rancher/vm/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ARPTablesGetter has a method to return a ARPTableInterface.
// A group's client should implement this interface.
type ARPTablesGetter interface {
	ARPTables() ARPTableInterface
}

// ARPTableInterface has methods to work with ARPTable resources.
type ARPTableInterface interface {
	Create(*v1alpha1.ARPTable) (*v1alpha1.ARPTable, error)
	Update(*v1alpha1.ARPTable) (*v1alpha1.ARPTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ARPTable, error)
	List(opts v1.ListOptions) (*v1alpha1.ARPTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ARPTable, err error)
	ARPTableExpansion
}

// aRPTables implements ARPTableInterface
type aRPTables struct {
	client rest.Interface
}

// newARPTables returns a ARPTables
func newARPTables(c *VirtualmachineV1alpha1Client) *aRPTables {
	return &aRPTables{
		client: c.RESTClient(),
	}
}

// Get takes name of the aRPTable, and returns the corresponding aRPTable object, and an error if there is any.
func (c *aRPTables) Get(name string, options v1.GetOptions) (result *v1alpha1.ARPTable, err error) {
	result = &v1alpha1.ARPTable{}
	err = c.client.Get().
		Resource("arptables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ARPTables that match those selectors.
func (c *aRPTables) List(opts v1.ListOptions) (result *v1alpha1.ARPTableList, err error) {
	result = &v1alpha1.ARPTableList{}
	err = c.client.Get().
		Resource("arptables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aRPTables.
func (c *aRPTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("arptables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a aRPTable and creates it.  Returns the server's representation of the aRPTable, and an error, if there is any.
func (c *aRPTables) Create(aRPTable *v1alpha1.ARPTable) (result *v1alpha1.ARPTable, err error) {
	result = &v1alpha1.ARPTable{}
	err = c.client.Post().
		Resource("arptables").
		Body(aRPTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a aRPTable and updates it. Returns the server's representation of the aRPTable, and an error, if there is any.
func (c *aRPTables) Update(aRPTable *v1alpha1.ARPTable) (result *v1alpha1.ARPTable, err error) {
	result = &v1alpha1.ARPTable{}
	err = c.client.Put().
		Resource("arptables").
		Name(aRPTable.Name).
		Body(aRPTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the aRPTable and deletes it. Returns an error if one occurs.
func (c *aRPTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("arptables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aRPTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("arptables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched aRPTable.
func (c *aRPTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ARPTable, err error) {
	result = &v1alpha1.ARPTable{}
	err = c.client.Patch(pt).
		Resource("arptables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
