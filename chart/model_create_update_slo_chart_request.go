/*
 * Charts API
 *
 * An API for creating, retrieving, updating, and deleting SLO charts
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chart

type CreateUpdateSloChartRequest struct {
	// ID of an SLO instance attached to this chart
	SloId string `json:"sloId"`
}
