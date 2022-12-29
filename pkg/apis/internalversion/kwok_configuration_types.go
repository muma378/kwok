/*
Copyright 2022 The Kubernetes Authors.

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

package internalversion

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KwokConfiguration provides configuration for the Kwok.
type KwokConfiguration struct {
	metav1.TypeMeta
	// Standard list metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	metav1.ObjectMeta
	// Options holds information about the default value.
	Options KwokConfigurationOptions
}

type KwokConfigurationOptions struct {
	// The default IP assigned to the Pod on maintained Nodes.
	CIDR string

	// The ip of all nodes maintained by the Kwok
	NodeIP string

	// Default option to manage (i.e., maintain heartbeat/liveness of) all Nodes or not.
	ManageAllNodes bool

	// Default annotations specified on Nodes to demand manage.
	ManageNodesWithAnnotationSelector string

	// Default labels specified on Nodes to demand manage.
	ManageNodesWithLabelSelector string

	// If a Node/Pod is on a managed Node and has this annotation status will not be modified
	DisregardStatusWithAnnotationSelector string

	// If a Node/Pod is on a managed Node and has this label status will not be modified
	DisregardStatusWithLabelSelector string

	// ServerAddress is server address of the Kwok.
	ServerAddress string

	// Experimental support for getting pod ip from CNI, for CNI-related components
	EnableCNI bool
}