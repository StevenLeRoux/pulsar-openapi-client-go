/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ConsumerStats struct for ConsumerStats
type ConsumerStats struct {
	MsgRateOut float64 `json:"msgRateOut,omitempty"`
	MsgThroughputOut float64 `json:"msgThroughputOut,omitempty"`
	MsgRateRedeliver float64 `json:"msgRateRedeliver,omitempty"`
	ConsumerName string `json:"consumerName,omitempty"`
	AvailablePermits int32 `json:"availablePermits,omitempty"`
	UnackedMessages int32 `json:"unackedMessages,omitempty"`
	BlockedConsumerOnUnackedMsgs bool `json:"blockedConsumerOnUnackedMsgs,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	Address string `json:"address,omitempty"`
	ConnectedSince string `json:"connectedSince,omitempty"`
	ClientVersion string `json:"clientVersion,omitempty"`
}
