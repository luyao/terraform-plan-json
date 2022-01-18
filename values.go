package jsonplan

// StateValues is the common representation of resolved values for both the
// prior state (which is always complete) and the planned new state.
type StateValues struct {
	Outputs    map[string]output `json:"outputs,omitempty"`
	RootModule module            `json:"root_module,omitempty"`
}

// AttributeValues is the JSON representation of the attribute values of the
// resource, whose structure depends on the resource type schema.
type AttributeValues map[string]interface{}
