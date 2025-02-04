package controller

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type DependentResource interface {
	Name() string
	Fetch(helper ReconcilerHelper) (runtime.Object, error)
	Build() (runtime.Object, error)
	Update(toUpdate runtime.Object) (bool, error)
	NewInstanceWith(owner Resource) DependentResource
	Owner() Resource
	Prototype() runtime.Object
	ShouldWatch() bool
	CanBeCreatedOrUpdated() bool
	ShouldBeOwned() bool
}

type DependentResourceHelper struct {
	_owner     Resource
	_prototype runtime.Object
	_delegate  DependentResource
}

func (res DependentResourceHelper) ShouldWatch() bool {
	return true
}

func (res DependentResourceHelper) ShouldBeOwned() bool {
	return true
}

func (res DependentResourceHelper) CanBeCreatedOrUpdated() bool {
	return true
}

func NewDependentResource(primaryResourceType runtime.Object, owner Resource) *DependentResourceHelper {
	return &DependentResourceHelper{_prototype: primaryResourceType, _owner: owner}
}

func (res *DependentResourceHelper) SetDelegate(delegate DependentResource) {
	res._delegate = delegate
}

func (res DependentResourceHelper) Name() string {
	return DefaultDependentResourceNameFor(res.Owner())
}

func (res DependentResourceHelper) Fetch(helper ReconcilerHelper) (runtime.Object, error) {
	delegate := res._delegate
	into := delegate.Prototype()
	if err := helper.Client.Get(context.TODO(), types.NamespacedName{Name: delegate.Name(), Namespace: delegate.Owner().GetNamespace()}, into); err != nil {
		return nil, err
	}
	return into, nil
}

func (res DependentResourceHelper) Owner() Resource {
	return res._owner
}

func (res DependentResourceHelper) Prototype() runtime.Object {
	return res._prototype.DeepCopyObject()
}
