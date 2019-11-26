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

package verify

import (
	"context"
	"time"

	"k8s.io/client-go/tools/clientcmd"

	"github.com/stretchr/testify/suite"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"

	e2ev1alpha2 "github.com/kubermatic/kubecarrier/pkg/apis/e2e/v1alpha2"
	"github.com/kubermatic/kubecarrier/test/framework"
)

var _ suite.SetupAllSuite = (*VerifySuite)(nil)

// VerifySuite verifies if we can reach both kubernetes clusters (master and service).
// and whether they are configured for our e2e tests.
type VerifySuite struct {
	suite.Suite
	*framework.Framework

	masterClient  client.Client
	serviceClient client.Client
}

func (s *VerifySuite) SetupSuite() {
	t := s.T()
	t.Logf("master cluster external kubeconfig location: %s", s.Framework.Config().MasterExternalKubeconfigPath)
	t.Logf("master cluster internal kubeconfig location: %s", s.Framework.Config().MasterInternalKubeconfigPath)
	t.Logf("svc cluster external kubeconfig location: %s", s.Framework.Config().ServiceExternalKubeconfigPath)
	t.Logf("svc cluster internal kubeconfig location: %s", s.Framework.Config().ServiceInternalKubeconfigPath)

	var err error
	s.Require().NoError(err, "creating testing framework")
	s.masterClient, err = s.MasterClient()
	s.Require().NoError(err, "creating master client")
	s.serviceClient, err = s.ServiceClient()
	s.Require().NoError(err, "creating service client")
}

func (s *VerifySuite) logHost(kubeconfig []byte) {
	cfg, err := clientcmd.NewClientConfigFromBytes(kubeconfig)
	s.Require().NoError(err)
	c, err := cfg.ClientConfig()
	s.Require().NoError(err)
	s.T().Logf("cluster-info kubeconfig host: %s", c.Host)
}

func (s *VerifySuite) TestValidMasterKubeconfig() {
	cm := &corev1.ConfigMap{}
	s.Require().NoError(s.masterClient.Get(context.Background(), types.NamespacedName{
		Name:      "cluster-info",
		Namespace: "kube-public",
	}, cm), "cannot fetch cluster-info")
	s.logHost([]byte(cm.Data["kubeconfig"]))
}

func (s *VerifySuite) TestValidServiceKubeconfig() {
	cm := &corev1.ConfigMap{}
	s.Require().NoError(s.serviceClient.Get(context.Background(), types.NamespacedName{
		Name:      "cluster-info",
		Namespace: "kube-public",
	}, cm), "cannot fetch cluster-info")
	s.logHost([]byte(cm.Data["kubeconfig"]))
}

func (s *VerifySuite) TestJokeOperatorSuccess() {
	s.T().Parallel()
	s.EnsureJokeOperator(s.T())

	jokes := []e2ev1alpha2.JokeItem{
		{
			// https://twitter.com/wm/status/1172654176742105089?lang=en
			Text: "A devops engineer walks into a bar, puts the bartender in a docker container, put kubernetes behind the bar, spins up 1000 bartenders, orders 1 beer.",
			Type: "kubernetes",
		},
		{
			// https://www.reddit.com/r/sysadmin/comments/625mk9/sysadmindevops_jokes/dfjwac5/
			Text: "I'd tell you the one about UDP, but you wouldn't get it.",
			Type: "devops",
		},
	}

	ctx := context.Background()
	c := s.serviceClient
	joke := &e2ev1alpha2.Joke{
		ObjectMeta: v1.ObjectMeta{
			Name:      "dummy",
			Namespace: "default",
		},
		Spec: e2ev1alpha2.JokeSpec{
			JokeDatabase: jokes,
		},
	}
	defer s.Assert().NoError(c.Delete(ctx, joke))
	s.Require().NoError(c.Create(ctx, joke))
	s.Assert().NoError(wait.Poll(time.Second, 5*time.Second, func() (done bool, err error) {
		if err := c.Get(ctx, types.NamespacedName{
			Namespace: joke.Namespace,
			Name:      joke.Name,
		}, joke); err != nil {
			return false, err
		}
		cond, ok := joke.Status.GetCondition(e2ev1alpha2.JokeReady)
		return ok && cond.Status == e2ev1alpha2.ConditionTrue && joke.Status.ObservedGeneration == joke.Generation, nil
	}), "joke wasn't ready within the timeframe")
	s.T().Log("selected joke: " + joke.Status.SelectedJoke.Text)
}

func (s *VerifySuite) TestJokeFailure() {
	s.T().Parallel()
	jokes := []e2ev1alpha2.JokeItem{}

	ctx := context.Background()
	c := s.serviceClient
	joke := &e2ev1alpha2.Joke{
		ObjectMeta: v1.ObjectMeta{
			Name:      "dummy-2",
			Namespace: "default",
		},
		Spec: e2ev1alpha2.JokeSpec{
			JokeDatabase: jokes,
		},
	}
	defer s.Assert().NoError(c.Delete(ctx, joke))
	s.Require().NoError(c.Create(ctx, joke))
	s.Require().NoError(wait.Poll(time.Second, 5*time.Second, func() (done bool, err error) {
		if err := c.Get(ctx, types.NamespacedName{
			Namespace: joke.Namespace,
			Name:      joke.Name,
		}, joke); err != nil {
			return false, err
		}
		cond, ok := joke.Status.GetCondition(e2ev1alpha2.JokeReady)
		return ok && cond.Status == e2ev1alpha2.ConditionFalse && joke.Status.ObservedGeneration == joke.Generation, nil
	}), "joke wasn't marked as failed within the timeframe")
	cond, _ := joke.Status.GetCondition(e2ev1alpha2.JokeReady)
	s.T().Log("joke status message" + cond.Message)
}