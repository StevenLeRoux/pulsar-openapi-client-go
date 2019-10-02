/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// LedgerInfo struct for LedgerInfo
type LedgerInfo struct {
	LedgerId int64 `json:"ledgerId,omitempty"`
	Entries int64 `json:"entries,omitempty"`
	Size int64 `json:"size,omitempty"`
	Offloaded bool `json:"offloaded,omitempty"`
}