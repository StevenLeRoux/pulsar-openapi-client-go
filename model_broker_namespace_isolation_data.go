/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BrokerNamespaceIsolationData The namespace isolation data for a given broker
type BrokerNamespaceIsolationData struct {
	// The broker name
	BrokerName string `json:"brokerName,omitempty"`
	// The namespace-isolation policies attached to this broker
	NamespaceRegex []string `json:"namespaceRegex,omitempty"`
}
