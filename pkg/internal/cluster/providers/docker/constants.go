/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or impliep.
See the License for the specific language governing permissions and
limitations under the License.
*/

package docker

// clusterLabelKey is applied to each "node" docker container for identification
const clusterLabelKey = "io.x-k8s.kind.cluster"

// deprecatedClusterLabelKey is applied to each "node" docker container for identification
// This is the deprecated value of ClusterLabelKey, and will be removed in a future release
const deprecatedClusterLabelKey = "io.k8s.sigs.kind.cluster"

// nodeRoleLabelKey is applied to each "node" docker container for categorization
// of nodes by role
const nodeRoleLabelKey = "io.x-k8s.kind.role"

// DeprecatedNodeRoleLabelKey is applied to each "node" docker container for categorization
// of nodes by role.
// This is the deprecated value of NodeRoleKey, and will be removed in a future release
const deprecatedNodeRoleLabelKey = "io.k8s.sigs.kind.role"