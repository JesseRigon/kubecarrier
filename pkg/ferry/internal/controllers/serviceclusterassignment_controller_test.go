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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	fakeclient "sigs.k8s.io/controller-runtime/pkg/client/fake"

	"k8c.io/utils/pkg/testutil"

	corev1alpha1 "k8c.io/kubecarrier/pkg/apis/core/v1alpha1"
)

func TestServiceClusterAssignmentReconciler(t *testing.T) {
	serviceClusterAssignment := &corev1alpha1.ServiceClusterAssignment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo.eu-west-1",
			Namespace: "provider-bar",
		},
		Spec: corev1alpha1.ServiceClusterAssignmentSpec{
			ManagementClusterNamespace: corev1alpha1.ObjectReference{
				Name: "foo",
			},
			ServiceCluster: corev1alpha1.ObjectReference{
				Name: "eu-west-1",
			},
		},
	}

	r := ServiceClusterAssignmentReconciler{
		Log:              testutil.NewLogger(t),
		ManagementClient: fakeclient.NewFakeClientWithScheme(testScheme, serviceClusterAssignment),
		ManagementScheme: testScheme,
		ServiceClient:    fakeclient.NewFakeClientWithScheme(testScheme),
	}
	_, err := r.Reconcile(ctrl.Request{
		NamespacedName: types.NamespacedName{
			Name:      serviceClusterAssignment.Name,
			Namespace: serviceClusterAssignment.Namespace,
		},
	})
	require.NoError(t, err, "Reconcile")

	ctx := context.Background()
	namespaceList := &corev1.NamespaceList{}
	require.NoError(t, r.ServiceClient.List(ctx, namespaceList), "listing Namespaces")
	if assert.Len(t, namespaceList.Items, 1) {
		assert.Equal(t, "foo-", namespaceList.Items[0].GenerateName)
	}
}
