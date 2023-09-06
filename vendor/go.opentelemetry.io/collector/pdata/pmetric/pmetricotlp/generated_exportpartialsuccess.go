// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

// Code generated by "pdata/internal/cmd/pdatagen/main.go". DO NOT EDIT.
// To regenerate this file run "make genpdata".

package pmetricotlp

import (
	otlpcollectormetrics "go.opentelemetry.io/collector/pdata/internal/data/protogen/collector/metrics/v1"
)

// ExportPartialSuccess represents the details of a partially successful export request.
//
// This is a reference type, if passed by value and callee modifies it the
// caller will see the modification.
//
// Must use NewExportPartialSuccess function to create new instances.
// Important: zero-initialized instance is not valid for use.
type ExportPartialSuccess struct {
	orig *otlpcollectormetrics.ExportMetricsPartialSuccess
}

func newExportPartialSuccess(orig *otlpcollectormetrics.ExportMetricsPartialSuccess) ExportPartialSuccess {
	return ExportPartialSuccess{orig}
}

// NewExportPartialSuccess creates a new empty ExportPartialSuccess.
//
// This must be used only in testing code. Users should use "AppendEmpty" when part of a Slice,
// OR directly access the member if this is embedded in another struct.
func NewExportPartialSuccess() ExportPartialSuccess {
	return newExportPartialSuccess(&otlpcollectormetrics.ExportMetricsPartialSuccess{})
}

// MoveTo moves all properties from the current struct overriding the destination and
// resetting the current instance to its zero value
func (ms ExportPartialSuccess) MoveTo(dest ExportPartialSuccess) {
	*dest.orig = *ms.orig
	*ms.orig = otlpcollectormetrics.ExportMetricsPartialSuccess{}
}

// RejectedDataPoints returns the rejecteddatapoints associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) RejectedDataPoints() int64 {
	return ms.orig.RejectedDataPoints
}

// SetRejectedDataPoints replaces the rejecteddatapoints associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) SetRejectedDataPoints(v int64) {
	ms.orig.RejectedDataPoints = v
}

// ErrorMessage returns the errormessage associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) ErrorMessage() string {
	return ms.orig.ErrorMessage
}

// SetErrorMessage replaces the errormessage associated with this ExportPartialSuccess.
func (ms ExportPartialSuccess) SetErrorMessage(v string) {
	ms.orig.ErrorMessage = v
}

// CopyTo copies all properties from the current struct overriding the destination.
func (ms ExportPartialSuccess) CopyTo(dest ExportPartialSuccess) {
	dest.SetRejectedDataPoints(ms.RejectedDataPoints())
	dest.SetErrorMessage(ms.ErrorMessage())
}