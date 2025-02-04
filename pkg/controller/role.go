package controller

import (
	"fmt"
	authorizv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type role struct {
	*DependentResourceHelper
}

func (res role) Update(toUpdate runtime.Object) (bool, error) {
	return false, nil
}

func (res role) NewInstanceWith(owner Resource) DependentResource {
	return newOwnedRole(owner)
}

func NewRole() role {
	return newOwnedRole(nil)
}

func newOwnedRole(owner Resource) role {
	dependent := NewDependentResource(&authorizv1.Role{}, owner)
	role := role{DependentResourceHelper: dependent}
	dependent.SetDelegate(role)
	return role
}

func RoleName(owner Resource) string {
	switch owner.(type) {
	case *Component:
		return "image-scc-privileged-role"
	case *Capability:
		return "scc-privileged-role"
	default:
		panic(fmt.Sprintf("unknown type '%s' for role owner", GetObjectName(owner)))
	}
}

func (res role) Name() string {
	return RoleName(res.Owner())
}

func (res role) Build() (runtime.Object, error) {
	c := res.Owner()
	ser := &authorizv1.Role{
		ObjectMeta: metav1.ObjectMeta{
			Name:      res.Name(),
			Namespace: c.GetNamespace(),
		},
		Rules: []authorizv1.PolicyRule{
			{
				APIGroups:     []string{"security.openshift.io"},
				Resources:     []string{"securitycontextconstraints"},
				ResourceNames: []string{"privileged"},
				Verbs:         []string{"use"},
			},
		},
	}

	if _, ok := c.(*Component); ok {
		ser.Rules = append(ser.Rules, authorizv1.PolicyRule{
			APIGroups: []string{"image.openshift.io"},
			Resources: []string{"imagestreams", "imagestreams/layers"},
			Verbs:     []string{"*"},
		})
	}

	return ser, nil
}

func (res role) ShouldWatch() bool {
	return false
}
