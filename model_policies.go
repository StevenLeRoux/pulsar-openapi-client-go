/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Policies struct for Policies
type Policies struct {
	AuthPolicies AuthPolicies `json:"auth_policies,omitempty"`
	ReplicationClusters []string `json:"replication_clusters,omitempty"`
	Bundles BundlesData `json:"bundles,omitempty"`
	BacklogQuotaMap map[string]BacklogQuota `json:"backlog_quota_map,omitempty"`
	TopicDispatchRate map[string]DispatchRate `json:"topicDispatchRate,omitempty"`
	SubscriptionDispatchRate map[string]DispatchRate `json:"subscriptionDispatchRate,omitempty"`
	ReplicatorDispatchRate map[string]DispatchRate `json:"replicatorDispatchRate,omitempty"`
	ClusterSubscribeRate map[string]SubscribeRate `json:"clusterSubscribeRate,omitempty"`
	Persistence PersistencePolicies `json:"persistence,omitempty"`
	DeduplicationEnabled bool `json:"deduplicationEnabled,omitempty"`
	LatencyStatsSampleRate map[string]int32 `json:"latency_stats_sample_rate,omitempty"`
	MessageTtlInSeconds int32 `json:"message_ttl_in_seconds,omitempty"`
	RetentionPolicies RetentionPolicies `json:"retention_policies,omitempty"`
	Deleted bool `json:"deleted,omitempty"`
	AntiAffinityGroup string `json:"antiAffinityGroup,omitempty"`
	EncryptionRequired bool `json:"encryption_required,omitempty"`
	SubscriptionAuthMode string `json:"subscription_auth_mode,omitempty"`
	MaxProducersPerTopic int32 `json:"max_producers_per_topic,omitempty"`
	MaxConsumersPerTopic int32 `json:"max_consumers_per_topic,omitempty"`
	MaxConsumersPerSubscription int32 `json:"max_consumers_per_subscription,omitempty"`
	CompactionThreshold int64 `json:"compaction_threshold,omitempty"`
	OffloadThreshold int64 `json:"offload_threshold,omitempty"`
	OffloadDeletionLagMs int64 `json:"offload_deletion_lag_ms,omitempty"`
	SchemaAutoUpdateCompatibilityStrategy string `json:"schema_auto_update_compatibility_strategy,omitempty"`
	SchemaValidationEnforced bool `json:"schema_validation_enforced,omitempty"`
}
