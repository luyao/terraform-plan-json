// Package jsonplan implements methods for inputting a plan in a
// machine-readable json format.
// The models defined in this package are mainly copied from
// https://github.com/hashicorp/terraform/blob/main/internal/command/jsonplan,
// changes are made to:
// 1. remove marshaller;
// 2. export models and factory method;
// 3. add a file parser to parse the json to model
package jsonplan
