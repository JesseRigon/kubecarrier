/*
Copyright 2019 The KubeCarrier Authors.

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

package controllers

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"

	operatorv1alpha1 "github.com/kubermatic/kubecarrier/pkg/apis/operator/v1alpha1"
	apiserver "github.com/kubermatic/kubecarrier/pkg/internal/resources/api-server"
)

// +kubebuilder:rbac:groups=operator.kubecarrier.io,resources=apiservers,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=operator.kubecarrier.io,resources=apiservers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterrolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=clusterroles,verbs=get;list;watch;create;update;patch;delete;escalate;bind
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=roles,verbs=get;list;watch;create;update;patch;delete;escalate;bind
// +kubebuilder:rbac:groups=rbac.authorization.k8s.io,resources=rolebindings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=serviceaccounts,verbs=get;list;watch;create;update;patch;delete

type APIServerStrategy struct {
}

func (c *APIServerStrategy) GetObj() Component {
	return &operatorv1alpha1.APIServer{}
}

func (c *APIServerStrategy) GetDeletionObjectTypes() []runtime.Object {
	return []runtime.Object{
		&rbacv1.ClusterRole{},
		&rbacv1.ClusterRoleBinding{},
	}
}

func (c *APIServerStrategy) GetManifests(ctx context.Context, component Component) ([]unstructured.Unstructured, error) {
	apiServer, ok := component.(*operatorv1alpha1.APIServer)
	if !ok {
		return nil, fmt.Errorf("can't assert to APIServer: %v", component)
	}
	return apiserver.Manifests(
		apiserver.Config{
			Namespace: apiServer.Namespace,
			Name:      apiServer.Name,
			Spec:      apiServer.Spec,
		})
}

func (c *APIServerStrategy) AddWatches(builder *builder.Builder, scheme *runtime.Scheme) *builder.Builder {
	return builder.
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.ServiceAccount{}).
		Owns(&corev1.Service{}).
		Owns(&rbacv1.Role{}).
		Owns(&rbacv1.RoleBinding{})
}