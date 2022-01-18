package jsonplan

import (
	"encoding/json"
)

// FormatVersion represents the version of the json format and will be
// incremented for any change to this format that requires changes to a
// consuming parser.
const FormatVersion = "0.2"

// Plan is the top-level representation of the json format of a plan. It includes
// the complete config and current state.
type Plan struct {
	FormatVersion    string      `json:"format_version,omitempty"`
	TerraformVersion string      `json:"terraform_version,omitempty"`
	Variables        variables   `json:"variables,omitempty"`
	PlannedValues    StateValues `json:"planned_values,omitempty"`
	// ResourceDrift and ResourceChanges are sorted in a user-friendly order
	// that is undefined at this time, but consistent.
	ResourceDrift   []resourceChange  `json:"resource_drift,omitempty"`
	ResourceChanges []resourceChange  `json:"resource_changes,omitempty"`
	OutputChanges   map[string]Change `json:"output_changes,omitempty"`
	PriorState      json.RawMessage   `json:"prior_state,omitempty"`
	Config          json.RawMessage   `json:"configuration,omitempty"`
}

func NewPlan() *Plan {
	return &Plan{
		FormatVersion: FormatVersion,
	}
}

// Change is the representation of a proposed change for an object.
type Change struct {
	// Actions are the actions that will be taken on the object selected by the
	// properties below. Valid actions values are:
	//    ["no-op"]
	//    ["create"]
	//    ["read"]
	//    ["update"]
	//    ["delete", "create"]
	//    ["create", "delete"]
	//    ["delete"]
	// The two "replace" actions are represented in this way to allow callers to
	// e.g. just scan the list for "delete" to recognize all three situations
	// where the object will be deleted, allowing for any new deletion
	// combinations that might be added in future.
	Actions []string `json:"actions,omitempty"`

	// Before and After are representations of the object value both before and
	// after the action. For ["create"] and ["delete"] actions, either "before"
	// or "after" is unset (respectively). For ["no-op"], the before and after
	// values are identical. The "after" value will be incomplete if there are
	// values within it that won't be known until after apply.
	Before json.RawMessage `json:"before,omitempty"`
	After  json.RawMessage `json:"after,omitempty"`

	// AfterUnknown is an object value with similar structure to After, but
	// with all unknown leaf values replaced with true, and all known leaf
	// values omitted.  This can be combined with After to reconstruct a full
	// value after the action, including values which will only be known after
	// apply.
	AfterUnknown json.RawMessage `json:"after_unknown,omitempty"`

	// BeforeSensitive and AfterSensitive are object values with similar
	// structure to Before and After, but with all sensitive leaf values
	// replaced with true, and all non-sensitive leaf values omitted. These
	// objects should be combined with Before and After to prevent accidental
	// display of sensitive values in user interfaces.
	BeforeSensitive json.RawMessage `json:"before_sensitive,omitempty"`
	AfterSensitive  json.RawMessage `json:"after_sensitive,omitempty"`

	// ReplacePaths is an array of arrays representing a set of paths into the
	// object value which resulted in the action being "replace". This will be
	// omitted if the action is not replace, or if no paths caused the
	// replacement (for example, if the resource was tainted). Each path
	// consists of one or more steps, each of which will be a number or a
	// string.
	ReplacePaths json.RawMessage `json:"replace_paths,omitempty"`
}

type output struct {
	Sensitive bool            `json:"sensitive"`
	Value     json.RawMessage `json:"value,omitempty"`
}

// variables is the JSON representation of the variables provided to the current
// plan.
type variables map[string]*variable

type variable struct {
	Value json.RawMessage `json:"value,omitempty"`
}
