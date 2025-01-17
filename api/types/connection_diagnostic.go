/*
Copyright 2022 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package types

import (
	"github.com/gravitational/teleport/api/utils"
	"github.com/gravitational/trace"
)

// ConnectionDiagnostic represents a Connection Diagnostic.
type ConnectionDiagnostic interface {
	// ResourceWithLabels provides common resource methods.
	ResourceWithLabels

	// Whether the connection was successful
	IsSuccess() bool

	// The underlying message
	GetMessage() string
}

type ConnectionsDiagnostic []ConnectionDiagnostic

var _ ConnectionDiagnostic = &ConnectionDiagnosticV1{}

// NewConnectionDiagnosticV1 creates a new ConnectionDiagnosticV1 resource.
func NewConnectionDiagnosticV1(name string, labels map[string]string, spec ConnectionDiagnosticSpecV1) (*ConnectionDiagnosticV1, error) {
	c := &ConnectionDiagnosticV1{
		ResourceHeader: ResourceHeader{
			Version: V1,
			Kind:    KindConnectionDiagnostic,
			Metadata: Metadata{
				Name:   name,
				Labels: labels,
			},
		},
		Spec: spec,
	}

	if err := c.CheckAndSetDefaults(); err != nil {
		return nil, trace.Wrap(err)
	}

	return c, nil
}

// CheckAndSetDefaults checks and sets default values for any missing fields.
func (c *ConnectionDiagnosticV1) CheckAndSetDefaults() error {
	if c.Spec.Message == "" {
		return trace.BadParameter("ConnectionDiagnosticV1.Spec missing Message field")
	}

	return nil
}

// GetAllLabels returns combined static and dynamic labels.
func (c *ConnectionDiagnosticV1) GetAllLabels() map[string]string {
	return CombineLabels(c.Metadata.Labels, nil)
}

// GetStaticLabels returns the connection diagnostic static labels.
func (c *ConnectionDiagnosticV1) GetStaticLabels() map[string]string {
	return c.Metadata.Labels
}

// IsSuccess returns whether the connection was successful
func (c *ConnectionDiagnosticV1) IsSuccess() bool {
	return c.Spec.Success
}

// GetMessage returns the connection diagnostic message.
func (c *ConnectionDiagnosticV1) GetMessage() string {
	return c.Spec.Message
}

// MatchSearch goes through select field values and tries to
// match against the list of search values.
func (c *ConnectionDiagnosticV1) MatchSearch(values []string) bool {
	fieldVals := append(utils.MapToStrings(c.GetAllLabels()), c.GetName())
	return MatchSearch(fieldVals, values, nil)
}

// Origin returns the origin value of the resource.
func (c *ConnectionDiagnosticV1) Origin() string {
	return c.Metadata.Labels[OriginLabel]
}

// SetOrigin sets the origin value of the resource.
func (c *ConnectionDiagnosticV1) SetOrigin(o string) {
	c.Metadata.Labels[OriginLabel] = o
}

// SetStaticLabels sets the connection diagnostic static labels.
func (c *ConnectionDiagnosticV1) SetStaticLabels(sl map[string]string) {
	c.Metadata.Labels = sl
}
