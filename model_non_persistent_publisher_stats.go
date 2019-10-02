/*
 * Pulsar Admin REST API
 *
 * This provides the REST API for admin operations
 *
 * API version: v2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// NonPersistentPublisherStats struct for NonPersistentPublisherStats
type NonPersistentPublisherStats struct {
	MsgRateIn float64 `json:"msgRateIn,omitempty"`
	MsgThroughputIn float64 `json:"msgThroughputIn,omitempty"`
	AverageMsgSize float64 `json:"averageMsgSize,omitempty"`
	ProducerId int64 `json:"producerId,omitempty"`
	Metadata map[string]string `json:"metadata,omitempty"`
	MsgDropRate float64 `json:"msgDropRate,omitempty"`
	Address string `json:"address,omitempty"`
	ProducerName string `json:"producerName,omitempty"`
	ConnectedSince string `json:"connectedSince,omitempty"`
	ClientVersion string `json:"clientVersion,omitempty"`
}
