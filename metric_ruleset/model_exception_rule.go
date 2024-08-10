/*
Metric Ruleset API

 Metric ruleset API 

API version: 3.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metric_ruleset

import (
	"encoding/json"
)

// ExceptionRule Fields for exception rules. 
type ExceptionRule struct {
	// Name of the exception rule. 
	Name string `json:"name"`
	// Whether to reroute the metric with the specified dimension. If false, the rule remains defined but will not reroute metrics. 
	Enabled bool `json:"enabled"`
	// Finds the metric to reroute. 
	Matcher DimensionMatcher `json:"matcher"`
	// Contains fields for the restoration job. The restoration job reroutes metrics from the archival route to the real-time route.  
	Restoration ExceptionRuleRestorationFields `json:"restoration,omitempty"`
	// Information about an exception rule. 
	Description *string `json:"description,omitempty"`
}

type _ExceptionRule ExceptionRule

// NewExceptionRule instantiates a new ExceptionRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExceptionRule(name string, enabled bool, matcher DimensionMatcher) *ExceptionRule {
	this := ExceptionRule{}
	this.Name = name
	this.Enabled = enabled
	this.Matcher = matcher
	return &this
}

// NewExceptionRuleWithDefaults instantiates a new ExceptionRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExceptionRuleWithDefaults() *ExceptionRule {
	this := ExceptionRule{}
	return &this
}

// GetName returns the Name field value
func (o *ExceptionRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExceptionRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExceptionRule) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *ExceptionRule) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ExceptionRule) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *ExceptionRule) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMatcher returns the Matcher field value
func (o *ExceptionRule) GetMatcher() DimensionMatcher {
	if o == nil {
		var ret DimensionMatcher
		return ret
	}

	return o.Matcher
}

// GetMatcherOk returns a tuple with the Matcher field value
// and a boolean to check if the value has been set.
func (o *ExceptionRule) GetMatcherOk() (*DimensionMatcher, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matcher, true
}

// SetMatcher sets field value
func (o *ExceptionRule) SetMatcher(v DimensionMatcher) {
	o.Matcher = v
}

// GetRestoration returns the Restoration field value
func (o *ExceptionRule) GetRestoration() ExceptionRuleRestorationFields {
	if o == nil {
		var ret ExceptionRuleRestorationFields
		return ret
	}

	return o.Restoration
}

// GetRestorationOk returns a tuple with the Restoration field value
// and a boolean to check if the value has been set.
func (o *ExceptionRule) GetRestorationOk() (*ExceptionRuleRestorationFields, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Restoration, true
}

// SetRestoration sets field value
func (o *ExceptionRule) SetRestoration(v ExceptionRuleRestorationFields) {
	o.Restoration = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExceptionRule) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExceptionRule) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExceptionRule) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExceptionRule) SetDescription(v string) {
	o.Description = &v
}

func (o ExceptionRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExceptionRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["enabled"] = o.Enabled
	toSerialize["matcher"] = o.Matcher
	if !isNil(o.Restoration) {
		toSerialize["restoration"] = o.Restoration
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableExceptionRule struct {
	value *ExceptionRule
	isSet bool
}

func (v NullableExceptionRule) Get() *ExceptionRule {
	return v.value
}

func (v *NullableExceptionRule) Set(val *ExceptionRule) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionRule) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionRule(val *ExceptionRule) *NullableExceptionRule {
	return &NullableExceptionRule{value: val, isSet: true}
}

func (v NullableExceptionRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

