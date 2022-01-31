/*
Copyright 2022 The KCP Authors.

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

package bootstrap

import (
	rbacv1 "k8s.io/api/rbac/v1"
	rbacv1helpers "k8s.io/kubernetes/pkg/apis/rbac/v1"
	rbacrest "k8s.io/kubernetes/pkg/registry/rbac/rest"
)

// ClusterRoleBindings return default rolebindings to the default roles
func clusterRoleBindings() []rbacv1.ClusterRoleBinding {
	return []rbacv1.ClusterRoleBinding{
		clusterRoleBindingCustomName(rbacv1helpers.NewClusterBinding("edit").Groups("system:kcp:workspace:edit").BindingOrDie(), "system:kcp:workspace:edit"),
		clusterRoleBindingCustomName(rbacv1helpers.NewClusterBinding("view").Groups("system:kcp:workspace:view").BindingOrDie(), "system:kcp:workspace:view"),
		clusterRoleBindingCustomName(rbacv1helpers.NewClusterBinding("admin").Groups("system:kcp:workspace:admin").BindingOrDie(), "system:kcp:workspace:admin"),
	}
}

func clusterRoleBindingCustomName(b rbacv1.ClusterRoleBinding, name string) rbacv1.ClusterRoleBinding {
	b.ObjectMeta.Name = name
	return b
}

func Policy() *rbacrest.PolicyData {
	return &rbacrest.PolicyData{
		ClusterRoleBindings: clusterRoleBindings(),
	}
}
