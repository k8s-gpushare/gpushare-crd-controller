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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/k8s-gpushare/gpushare-crd-controller/pkg/apis/nvidia/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GPUNodeInfoLister helps list GPUNodeInfos.
// All objects returned here must be treated as read-only.
type GPUNodeInfoLister interface {
	// List lists all GPUNodeInfos in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GPUNodeInfo, err error)
	// GPUNodeInfos returns an object that can list and get GPUNodeInfos.
	GPUNodeInfos(namespace string) GPUNodeInfoNamespaceLister
	GPUNodeInfoListerExpansion
}

// GPUNodeInfoLister implements the GPUNodeInfoLister interface.
type GPUNodeInfoLister struct {
	indexer cache.Indexer
}

// NewGPUNodeInfoLister returns a new GPUNodeInfoLister.
func NewGPUNodeInfoLister(indexer cache.Indexer) GPUNodeInfoLister {
	return &GPUNodeInfoLister{indexer: indexer}
}

// List lists all GPUNodeInfos in the indexer.
func (s *GPUNodeInfoLister) List(selector labels.Selector) (ret []*v1.GPUNodeInfo, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GPUNodeInfo))
	})
	return ret, err
}

// GPUNodeInfos returns an object that can list and get GPUNodeInfos.
func (s *GPUNodeInfoLister) GPUNodeInfos(namespace string) GPUNodeInfoNamespaceLister {
	return GPUNodeInfoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GPUNodeInfoNamespaceLister helps list and get GPUNodeInfos.
// All objects returned here must be treated as read-only.
type GPUNodeInfoNamespaceLister interface {
	// List lists all GPUNodeInfos in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GPUNodeInfo, err error)
	// Get retrieves the GPUNodeInfo from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GPUNodeInfo, error)
	GPUNodeInfoNamespaceListerExpansion
}

// GPUNodeInfoNamespaceLister implements the GPUNodeInfoNamespaceLister
// interface.
type GPUNodeInfoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GPUNodeInfos in the indexer for a given namespace.
func (s GPUNodeInfoNamespaceLister) List(selector labels.Selector) (ret []*v1.GPUNodeInfo, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GPUNodeInfo))
	})
	return ret, err
}

// Get retrieves the GPUNodeInfo from the indexer for a given namespace and name.
func (s GPUNodeInfoNamespaceLister) Get(name string) (*v1.GPUNodeInfo, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("GPUNodeInfo"), name)
	}
	return obj.(*v1.GPUNodeInfo), nil
}
