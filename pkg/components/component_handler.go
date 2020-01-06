// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package components

import components_v1alpha1 "dapr/pkg/apis/components/v1alpha1"

// ComponentHandler is an interface for reacting on component changes
type ComponentHandler interface {
	OnComponentUpdated(component components_v1alpha1.Component)
}
