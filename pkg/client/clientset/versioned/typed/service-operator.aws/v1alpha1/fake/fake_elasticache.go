/*
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

package fake

import (
	v1alpha1 "github.com/awslabs/aws-service-operator/pkg/apis/service-operator.aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeElastiCaches implements ElastiCacheInterface
type FakeElastiCaches struct {
	Fake *FakeServiceoperatorV1alpha1
	ns   string
}

var elasticachesResource = schema.GroupVersionResource{Group: "serviceoperator.aws", Version: "v1alpha1", Resource: "elasticaches"}

var elasticachesKind = schema.GroupVersionKind{Group: "serviceoperator.aws", Version: "v1alpha1", Kind: "ElastiCache"}

// Get takes name of the elastiCache, and returns the corresponding elastiCache object, and an error if there is any.
func (c *FakeElastiCaches) Get(name string, options v1.GetOptions) (result *v1alpha1.ElastiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(elasticachesResource, c.ns, name), &v1alpha1.ElastiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastiCache), err
}

// List takes label and field selectors, and returns the list of ElastiCaches that match those selectors.
func (c *FakeElastiCaches) List(opts v1.ListOptions) (result *v1alpha1.ElastiCacheList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(elasticachesResource, elasticachesKind, c.ns, opts), &v1alpha1.ElastiCacheList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ElastiCacheList{ListMeta: obj.(*v1alpha1.ElastiCacheList).ListMeta}
	for _, item := range obj.(*v1alpha1.ElastiCacheList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested elastiCaches.
func (c *FakeElastiCaches) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(elasticachesResource, c.ns, opts))

}

// Create takes the representation of a elastiCache and creates it.  Returns the server's representation of the elastiCache, and an error, if there is any.
func (c *FakeElastiCaches) Create(elastiCache *v1alpha1.ElastiCache) (result *v1alpha1.ElastiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(elasticachesResource, c.ns, elastiCache), &v1alpha1.ElastiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastiCache), err
}

// Update takes the representation of a elastiCache and updates it. Returns the server's representation of the elastiCache, and an error, if there is any.
func (c *FakeElastiCaches) Update(elastiCache *v1alpha1.ElastiCache) (result *v1alpha1.ElastiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(elasticachesResource, c.ns, elastiCache), &v1alpha1.ElastiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastiCache), err
}

// Delete takes name of the elastiCache and deletes it. Returns an error if one occurs.
func (c *FakeElastiCaches) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(elasticachesResource, c.ns, name), &v1alpha1.ElastiCache{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeElastiCaches) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(elasticachesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ElastiCacheList{})
	return err
}

// Patch applies the patch and returns the patched elastiCache.
func (c *FakeElastiCaches) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ElastiCache, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(elasticachesResource, c.ns, name, data, subresources...), &v1alpha1.ElastiCache{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ElastiCache), err
}
