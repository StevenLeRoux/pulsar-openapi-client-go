/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NamespaceOwnershipStatus struct for NamespaceOwnershipStatus
type NamespaceOwnershipStatus struct {
	BrokerAssignment string `json:"broker_assignment,omitempty"`
	IsControlled bool `json:"is_controlled,omitempty"`
	IsActive bool `json:"is_active,omitempty"`
}