//go:build e2e
// +build e2e

package azure_managed_prometheus

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/joho/godotenv"
)

// Load environment variables from .env file
var _ = godotenv.Load("../../.env")

const (
	testNamePodIdentity = "azure-managed-prometheus-pod-identity-test"
)

// Pod Identity test vars
var (
	randomNumberPod          = rand.Int()
	testNamespacePod         = fmt.Sprintf("%s-ns-%d", testNamePodIdentity, randomNumberPod)
	deploymentNamePod        = fmt.Sprintf("%s-deployment-%d", testNamePodIdentity, randomNumberPod)
	monitoredAppNamePod      = fmt.Sprintf("%s-monitored-app-%d", testNamePodIdentity, randomNumberPod)
	publishDeploymentNamePod = fmt.Sprintf("%s-publish-%d", testNamePodIdentity, randomNumberPod)
	scaledObjectNamePod      = fmt.Sprintf("%s-so-%d", testNamePodIdentity, randomNumberPod)
	podIdentityProvider      = "azure"
)

// TestAzureManagedPrometheusScalerWithPodIdentity creates deployments - there are two deployments - both using the same image but one deployment
// is directly tied to the KEDA HPA while the other is isolated that can be used for metrics
// even when the KEDA deployment is at zero - the service points to both deployments
func TestAzureManagedPrometheusScalerWithPodIdentity(t *testing.T) {
	testAzureManagedPrometheusScaler(t, getTemplateDataForPodIdentityTest())
}

func getTemplateDataForPodIdentityTest() templateData {
	return templateData{
		TestNamespace:           testNamespacePod,
		DeploymentName:          deploymentNamePod,
		PublishDeploymentName:   publishDeploymentNamePod,
		ScaledObjectName:        scaledObjectNamePod,
		MonitoredAppName:        monitoredAppNamePod,
		PodIdentityProvider:     podIdentityProvider,
		PrometheusQueryEndpoint: prometheusQueryEndpoint,
		MinReplicaCount:         minReplicaCount,
		MaxReplicaCount:         maxReplicaCount,
	}
}
