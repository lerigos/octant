/*
Copyright (c) 2019 the Octant contributors. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/

package clusteroverview

import (
	"path"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/vmware-tanzu/octant/internal/gvk"
)

var (
	supportedGVKs = []schema.GroupVersionKind{
		gvk.ClusterRoleBinding,
		gvk.ClusterRole,
		gvk.Node,
		gvk.PersistentVolume,
		gvk.Namespace,
	}
)

const rbacAPIVersion = "rbac.authorization.k8s.io/v1"

func crdPath(namespace, crdName, version, name string) (string, error) {
	return path.Join("/cluster-overview/custom-resources", crdName, version, name), nil
}

func gvkPath(namespace, apiVersion, kind, name string) (string, error) {
	var p string

	switch {
	case apiVersion == rbacAPIVersion && kind == "ClusterRole":
		p = "/rbac/cluster-roles"
	case apiVersion == rbacAPIVersion && kind == "ClusterRoleBinding":
		p = "/rbac/cluster-role-bindings"
	case apiVersion == "v1" && kind == "Node":
		p = "/nodes"
	case apiVersion == "v1" && kind == "PersistentVolume":
		p = "/storage/persistent-volumes"
	case apiVersion == "v1" && kind == "Namespace":
		p = "/namespaces"
	default:
		return "", errors.Errorf("unknown object %s %s", apiVersion, kind)
	}

	return path.Join("/cluster-overview", p, name), nil
}
