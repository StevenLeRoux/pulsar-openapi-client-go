/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// AllocatorStats struct for AllocatorStats
type AllocatorStats struct {
	NumDirectArenas int32 `json:"numDirectArenas,omitempty"`
	NumHeapArenas int32 `json:"numHeapArenas,omitempty"`
	NumThreadLocalCaches int32 `json:"numThreadLocalCaches,omitempty"`
	NormalCacheSize int32 `json:"normalCacheSize,omitempty"`
	SmallCacheSize int32 `json:"smallCacheSize,omitempty"`
	TinyCacheSize int32 `json:"tinyCacheSize,omitempty"`
	DirectArenas []PoolArenaStats `json:"directArenas,omitempty"`
	HeapArenas []PoolArenaStats `json:"heapArenas,omitempty"`
}
