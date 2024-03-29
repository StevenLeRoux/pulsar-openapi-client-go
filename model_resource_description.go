/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ResourceDescription struct for ResourceDescription
type ResourceDescription struct {
	UsagePct int32 `json:"usagePct,omitempty"`
	ResourceUsage map[string]ResourceUsage `json:"resourceUsage,omitempty"`
}
