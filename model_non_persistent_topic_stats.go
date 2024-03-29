/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NonPersistentTopicStats struct for NonPersistentTopicStats
type NonPersistentTopicStats struct {
	MsgRateIn float64 `json:"msgRateIn,omitempty"`
	MsgThroughputIn float64 `json:"msgThroughputIn,omitempty"`
	MsgRateOut float64 `json:"msgRateOut,omitempty"`
	MsgThroughputOut float64 `json:"msgThroughputOut,omitempty"`
	AverageMsgSize float64 `json:"averageMsgSize,omitempty"`
	StorageSize int64 `json:"storageSize,omitempty"`
	BacklogSize int64 `json:"backlogSize,omitempty"`
	Publishers []NonPersistentPublisherStats `json:"publishers,omitempty"`
	Subscriptions map[string]NonPersistentSubscriptionStats `json:"subscriptions,omitempty"`
	Replication map[string]NonPersistentReplicatorStats `json:"replication,omitempty"`
	DeduplicationStatus string `json:"deduplicationStatus,omitempty"`
	MsgDropRate float64 `json:"msgDropRate,omitempty"`
}
