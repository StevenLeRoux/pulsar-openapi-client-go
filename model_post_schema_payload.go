/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PostSchemaPayload struct for PostSchemaPayload
type PostSchemaPayload struct {
	Type string `json:"type,omitempty"`
	Schema string `json:"schema,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}