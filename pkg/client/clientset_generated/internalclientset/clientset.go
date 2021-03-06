/*
Copyright 2017 The Kubernetes Authors.

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

package internalclientset

import (
	"github.com/golang/glog"
	"k8s.io/client-go/pkg/util/flowcontrol"
	internalversionapps "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/apps/internalversion"
	internalversionauthentication "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authentication/internalversion"
	internalversionauthorization "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/authorization/internalversion"
	internalversionautoscaling "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/autoscaling/internalversion"
	internalversionbatch "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/batch/internalversion"
	internalversioncertificates "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/certificates/internalversion"
	internalversioncore "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/core/internalversion"
	internalversionextensions "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/extensions/internalversion"
	internalversionpolicy "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/policy/internalversion"
	internalversionrbac "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/rbac/internalversion"
	internalversionstorage "k8s.io/kubernetes/pkg/client/clientset_generated/internalclientset/typed/storage/internalversion"
	restclient "k8s.io/kubernetes/pkg/client/restclient"
	discovery "k8s.io/kubernetes/pkg/client/typed/discovery"
	_ "k8s.io/kubernetes/plugin/pkg/client/auth"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	Core() internalversioncore.CoreInterface

	Apps() internalversionapps.AppsInterface

	Authentication() internalversionauthentication.AuthenticationInterface

	Authorization() internalversionauthorization.AuthorizationInterface

	Autoscaling() internalversionautoscaling.AutoscalingInterface

	Batch() internalversionbatch.BatchInterface

	Certificates() internalversioncertificates.CertificatesInterface

	Extensions() internalversionextensions.ExtensionsInterface

	Policy() internalversionpolicy.PolicyInterface

	Rbac() internalversionrbac.RbacInterface

	Storage() internalversionstorage.StorageInterface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	*internalversioncore.CoreClient
	*internalversionapps.AppsClient
	*internalversionauthentication.AuthenticationClient
	*internalversionauthorization.AuthorizationClient
	*internalversionautoscaling.AutoscalingClient
	*internalversionbatch.BatchClient
	*internalversioncertificates.CertificatesClient
	*internalversionextensions.ExtensionsClient
	*internalversionpolicy.PolicyClient
	*internalversionrbac.RbacClient
	*internalversionstorage.StorageClient
}

// Core retrieves the CoreClient
func (c *Clientset) Core() internalversioncore.CoreInterface {
	if c == nil {
		return nil
	}
	return c.CoreClient
}

// Apps retrieves the AppsClient
func (c *Clientset) Apps() internalversionapps.AppsInterface {
	if c == nil {
		return nil
	}
	return c.AppsClient
}

// Authentication retrieves the AuthenticationClient
func (c *Clientset) Authentication() internalversionauthentication.AuthenticationInterface {
	if c == nil {
		return nil
	}
	return c.AuthenticationClient
}

// Authorization retrieves the AuthorizationClient
func (c *Clientset) Authorization() internalversionauthorization.AuthorizationInterface {
	if c == nil {
		return nil
	}
	return c.AuthorizationClient
}

// Autoscaling retrieves the AutoscalingClient
func (c *Clientset) Autoscaling() internalversionautoscaling.AutoscalingInterface {
	if c == nil {
		return nil
	}
	return c.AutoscalingClient
}

// Batch retrieves the BatchClient
func (c *Clientset) Batch() internalversionbatch.BatchInterface {
	if c == nil {
		return nil
	}
	return c.BatchClient
}

// Certificates retrieves the CertificatesClient
func (c *Clientset) Certificates() internalversioncertificates.CertificatesInterface {
	if c == nil {
		return nil
	}
	return c.CertificatesClient
}

// Extensions retrieves the ExtensionsClient
func (c *Clientset) Extensions() internalversionextensions.ExtensionsInterface {
	if c == nil {
		return nil
	}
	return c.ExtensionsClient
}

// Policy retrieves the PolicyClient
func (c *Clientset) Policy() internalversionpolicy.PolicyInterface {
	if c == nil {
		return nil
	}
	return c.PolicyClient
}

// Rbac retrieves the RbacClient
func (c *Clientset) Rbac() internalversionrbac.RbacInterface {
	if c == nil {
		return nil
	}
	return c.RbacClient
}

// Storage retrieves the StorageClient
func (c *Clientset) Storage() internalversionstorage.StorageInterface {
	if c == nil {
		return nil
	}
	return c.StorageClient
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *restclient.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.CoreClient, err = internalversioncore.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AppsClient, err = internalversionapps.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AuthenticationClient, err = internalversionauthentication.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AuthorizationClient, err = internalversionauthorization.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.AutoscalingClient, err = internalversionautoscaling.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.BatchClient, err = internalversionbatch.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.CertificatesClient, err = internalversioncertificates.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.ExtensionsClient, err = internalversionextensions.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.PolicyClient, err = internalversionpolicy.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.RbacClient, err = internalversionrbac.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.StorageClient, err = internalversionstorage.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *restclient.Config) *Clientset {
	var cs Clientset
	cs.CoreClient = internalversioncore.NewForConfigOrDie(c)
	cs.AppsClient = internalversionapps.NewForConfigOrDie(c)
	cs.AuthenticationClient = internalversionauthentication.NewForConfigOrDie(c)
	cs.AuthorizationClient = internalversionauthorization.NewForConfigOrDie(c)
	cs.AutoscalingClient = internalversionautoscaling.NewForConfigOrDie(c)
	cs.BatchClient = internalversionbatch.NewForConfigOrDie(c)
	cs.CertificatesClient = internalversioncertificates.NewForConfigOrDie(c)
	cs.ExtensionsClient = internalversionextensions.NewForConfigOrDie(c)
	cs.PolicyClient = internalversionpolicy.NewForConfigOrDie(c)
	cs.RbacClient = internalversionrbac.NewForConfigOrDie(c)
	cs.StorageClient = internalversionstorage.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c restclient.Interface) *Clientset {
	var cs Clientset
	cs.CoreClient = internalversioncore.New(c)
	cs.AppsClient = internalversionapps.New(c)
	cs.AuthenticationClient = internalversionauthentication.New(c)
	cs.AuthorizationClient = internalversionauthorization.New(c)
	cs.AutoscalingClient = internalversionautoscaling.New(c)
	cs.BatchClient = internalversionbatch.New(c)
	cs.CertificatesClient = internalversioncertificates.New(c)
	cs.ExtensionsClient = internalversionextensions.New(c)
	cs.PolicyClient = internalversionpolicy.New(c)
	cs.RbacClient = internalversionrbac.New(c)
	cs.StorageClient = internalversionstorage.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
