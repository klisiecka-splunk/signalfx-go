/*
 * Charts API
 *
 * An API for creating, retrieving, updating, and deleting charts
 *
 * API version: 2.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chart

type Chart struct {
	// The time the chart was created. The format is Unix time, UTC. This value is always set by the system.
	Created int64 `json:"created,omitempty"`
	// SignalFx user ID of the user that initially created the chart
	Creator string `json:"creator,omitempty"`
	// Description of the chart. This value appears underneath the chart name in the SignalFx web UI.
	Description string `json:"description,omitempty"`
	// System-defined identifier for the chart
	Id string `json:"id,omitempty"`
	// The last time the chart was updated, in Unix time UTC.
	LastUpdated int64 `json:"lastUpdated,omitempty"`
	// The SignalFx user ID of the last person who updated the chart. If the last update was done by the system, the value is the string literal \"AAAAAAAAAA\".
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
	// The displayed name of the chart in the dashboard
	Name    string   `json:"name,omitempty"`
	Options *Options `json:"options,omitempty"`
	// Specifies one or more SignalFlow packages to import for use with the SignalFlow program specified in the `programText` option. This option must be set to `signalfx`.
	PackageSpecifications string `json:"packageSpecifications,omitempty"`
	// The SignalFlow program that provides data for the chart.  If you use more than one line of SignalFlow, separate the lines with  semicolons (\";\") or newline characters (\"\\n\"). See the  [Charts Overview](https://developers.signalfx.com/v2/reference.html#charts-overview-1) for more information.\"
	ProgramText string `json:"programText,omitempty"`
	// An array that contains tag values. You can use tags to search for or filter charts. One use is to tag charts that are in production with the value `prod`.
	Tags []string `json:"tags,omitempty"`
	// ID of an SLO instance attached to this chart
	SloId string `json:"sloId,omitempty"`
}
