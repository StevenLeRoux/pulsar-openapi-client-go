/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PoolSubpageStats struct for PoolSubpageStats
type PoolSubpageStats struct {
	MaxNumElements int32 `json:"maxNumElements,omitempty"`
	NumAvailable int32 `json:"numAvailable,omitempty"`
	ElementSize int32 `json:"elementSize,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
}