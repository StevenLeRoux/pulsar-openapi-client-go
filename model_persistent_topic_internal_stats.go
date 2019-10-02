/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PersistentTopicInternalStats struct for PersistentTopicInternalStats
type PersistentTopicInternalStats struct {
	EntriesAddedCounter int64 `json:"entriesAddedCounter,omitempty"`
	NumberOfEntries int64 `json:"numberOfEntries,omitempty"`
	TotalSize int64 `json:"totalSize,omitempty"`
	CurrentLedgerEntries int64 `json:"currentLedgerEntries,omitempty"`
	CurrentLedgerSize int64 `json:"currentLedgerSize,omitempty"`
	LastLedgerCreatedTimestamp string `json:"lastLedgerCreatedTimestamp,omitempty"`
	LastLedgerCreationFailureTimestamp string `json:"lastLedgerCreationFailureTimestamp,omitempty"`
	WaitingCursorsCount int32 `json:"waitingCursorsCount,omitempty"`
	PendingAddEntriesCount int32 `json:"pendingAddEntriesCount,omitempty"`
	LastConfirmedEntry string `json:"lastConfirmedEntry,omitempty"`
	State string `json:"state,omitempty"`
	Ledgers []LedgerInfo `json:"ledgers,omitempty"`
	Cursors map[string]CursorStats `json:"cursors,omitempty"`
}
