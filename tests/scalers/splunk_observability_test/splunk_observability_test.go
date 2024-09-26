//go:build e2e
// +build e2e

package splunk_observability_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes"

	. "github.com/kedacore/keda/v2/tests/helper"
)

const (
	testName = "splunk-observability-test"
)

var (
	testNamespace          = fmt.Sprintf("%s-ns", testName)
	deploymentName         = fmt.Sprintf("%s-deployment", testName)
	scaledObjectName       = fmt.Sprintf("%s-so", testName)
	authName               = fmt.Sprintf("%s-auth", testName)
	duration               = "10"
	maxReplicaCount        = 10
	minReplicaCount        = 1
	scaleInTargetValue     = "400"
	scaleInActivationValue = "1.1"
)

type templateData struct {
	TestNamespace         string
	DeploymentName        string
	ScaledObjectName      string
	AuthName              string
	Duration              string
	MinReplicaCount       string
	MaxReplicaCount       string
	TargetValue           string
	ActivationTargetValue string
}

const (
	authTemplate = `
apiVersion: v1
kind: Secret
metadata:
  name: splunk-secrets
  namespace: {{.TestNamespace}}
data:
  accessToken: YW1JeUpqVHRJd185cDhOWG01X21KQQ==  # one time through-away access token used just for testing
  realm: dXMw
---
apiVersion: keda.sh/v1alpha1
kind: TriggerAuthentication
metadata:
  name: keda-trigger-auth-splunk-secret
  namespace: {{.TestNamespace}}
spec:
  secretTargetRef:
  - parameter: accessToken
    name: splunk-secrets
    key: accessToken
  - parameter: realm
    name: splunk-secrets
    key: realm
`

	deploymentTemplate = `
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  namespace: {{.TestNamespace}}
spec:
  selector:
    matchLabels:
      app: nginx
  replicas: 1 
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80
`

	scaledObjectTemplate = `
apiVersion: keda.sh/v1alpha1
kind: ScaledObject
metadata:
  name: keda
  namespace: {{.TestNamespace}}
spec:
  scaleTargetRef:
    apiVersion: apps/v1
    kind: Deployment
    name: nginx
  pollingInterval: 3
  cooldownPeriod: 1
  minReplicaCount: {{.MinReplicaCount}}
  maxReplicaCount: {{.MaxReplicaCount}}
  triggers:
  - type: splunk-observability
    metricType: Value
    metadata:
      query: "data('fdse-1989-tenable-test-metric').publish()"
      duration: "10"
      targetValue: "250" 
      activationTargetValue: "1.1"
      queryAggregator: "max" # 'min', 'max', or 'avg'
    authenticationRef:
      name: keda-trigger-auth-splunk-secret
`
)

func TestSplunkObservabilityScaler(t *testing.T) {
	kc := GetKubernetesClient(t)
	data, templates := getTemplateData()
	t.Cleanup(func() {
		DeleteKubernetesResources(t, testNamespace, data, templates)
	})

	// Create kubernetes resources
	CreateKubernetesResources(t, kc, testNamespace, data, templates)

	// Ensure nginx deployment is ready
	assert.True(t, WaitForAllPodRunningInNamespace(t, kc, testNamespace, minReplicaCount, 120),
		"replica count should be %d after 2 minutes", minReplicaCount)

	// test scaling
	testScaleOut(t, kc, testNamespace)
	testScaleIn(t, kc)
}

func getPodCount(kc *kubernetes.Clientset, namespace string) int {
	pods, err := kc.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	return len(pods.Items)
}

func testActivation(t *testing.T, kc *kubernetes.Clientset) {
	t.Log("--- testing activation ---")

	AssertReplicaCountNotChangeDuringTimePeriod(t, kc, deploymentName, testNamespace, minReplicaCount, 60)
}

func testScaleOut(t *testing.T, kc *kubernetes.Clientset, namespace string) {
	t.Log("--- testing scale out ---")
	t.Log("waiting for 3 minutes")
	time.Sleep(time.Duration(180) * time.Second)

	assert.True(t, getPodCount(kc, testNamespace) > minReplicaCount, "number of pods in deployment should be more than %d after 3 minutes", minReplicaCount)
}

func testScaleIn(t *testing.T, kc *kubernetes.Clientset) {
	t.Log("--- testing scale in ---")

	t.Log("waiting for 10 minutes")
	time.Sleep(time.Duration(600) * time.Second)

	assert.True(t, getPodCount(kc, testNamespace) > minReplicaCount, "number of pods in deployment should be less than %d after 10 minutes", maxReplicaCount)
}

func getTemplateData() (templateData, []Template) {
	return templateData{
			TestNamespace:         testNamespace,
			DeploymentName:        deploymentName,
			ScaledObjectName:      scaledObjectName,
			AuthName:              authName,
			Duration:              duration,
			MinReplicaCount:       fmt.Sprintf("%v", minReplicaCount),
			MaxReplicaCount:       fmt.Sprintf("%v", maxReplicaCount),
			TargetValue:           scaleInTargetValue,
			ActivationTargetValue: scaleInActivationValue,
		}, []Template{
			{Name: "authTemplate", Config: authTemplate},
			{Name: "scaledObjectTemplate", Config: scaledObjectTemplate},
			{Name: "deploymentTemplate", Config: deploymentTemplate},
		}
}

func getScaledObjectTemplateData(targetValue, activationTargetValue string) templateData {
	return templateData{
		TestNamespace:         testNamespace,
		DeploymentName:        deploymentName,
		ScaledObjectName:      scaledObjectName,
		MinReplicaCount:       fmt.Sprintf("%v", minReplicaCount),
		MaxReplicaCount:       fmt.Sprintf("%v", maxReplicaCount),
		TargetValue:           targetValue,
		ActivationTargetValue: activationTargetValue,
	}
}
