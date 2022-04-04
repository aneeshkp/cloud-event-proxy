package utils

const (
	// NamespaceForPTP .. ptp namespace
	NamespaceForPTP = "openshift-ptp"
	// DeploymentForPTP  ptp deployment
	DeploymentForPTP = "ptp-operator"
	// NamespaceForTesting ...
	NamespaceForTesting = "cne-testing"
	// NamespaceProducerTesting contains the name of the testing namespace
	NamespaceProducerTesting = "cloud-event-producer-testing"
	// NamespaceConsumerTesting ...
	NamespaceConsumerTesting = "cloud-event-consumer-testing"
	// NamespaceAMQTesting ...
	NamespaceAMQTesting = "amq-router-testing"
	// CloudEventProducerDeploymentName ...
	CloudEventProducerDeploymentName = "cloud-producer-deployment"
	// CloudEventConsumerDeploymentName ...
	CloudEventConsumerDeploymentName = "cloud-consumer-deployment"
	// AMQDeploymentName ...
	AMQDeploymentName = "interconnect-deployment"

	// EventProxyContainerName Event Proxy container name
	EventProxyContainerName = "cloud-event-proxy"
	// AmqInstanceName ...
	AmqInstanceName = "amq-router"
	// ConsumerContainerName ...
	ConsumerContainerName = "cloud-event-consumer"
)
