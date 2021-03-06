//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ApplicationPackage) DeepCopyInto(out *ApplicationPackage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(ChartRepoRef)
		**out = **in
	}
	out.Chart = in.Chart
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ApplicationPackage.
func (in *ApplicationPackage) DeepCopy() *ApplicationPackage {
	if in == nil {
		return nil
	}
	out := new(ApplicationPackage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ApplicationPackage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Badge) DeepCopyInto(out *Badge) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Badge.
func (in *Badge) DeepCopy() *Badge {
	if in == nil {
		return nil
	}
	out := new(Badge)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bundle) DeepCopyInto(out *Bundle) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bundle.
func (in *Bundle) DeepCopy() *Bundle {
	if in == nil {
		return nil
	}
	out := new(Bundle)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bundle) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleList) DeepCopyInto(out *BundleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bundle, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleList.
func (in *BundleList) DeepCopy() *BundleList {
	if in == nil {
		return nil
	}
	out := new(BundleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BundleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleOption) DeepCopyInto(out *BundleOption) {
	*out = *in
	out.BundleRef = in.BundleRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleOption.
func (in *BundleOption) DeepCopy() *BundleOption {
	if in == nil {
		return nil
	}
	out := new(BundleOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleOptionView) DeepCopyInto(out *BundleOptionView) {
	*out = *in
	in.PackageMeta.DeepCopyInto(&out.PackageMeta)
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]Feature, len(*in))
		copy(*out, *in)
	}
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]PackageCard, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleOptionView.
func (in *BundleOptionView) DeepCopy() *BundleOptionView {
	if in == nil {
		return nil
	}
	out := new(BundleOptionView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleRef) DeepCopyInto(out *BundleRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleRef.
func (in *BundleRef) DeepCopy() *BundleRef {
	if in == nil {
		return nil
	}
	out := new(BundleRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleSpec) DeepCopyInto(out *BundleSpec) {
	*out = *in
	in.PackageDescriptor.DeepCopyInto(&out.PackageDescriptor)
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]Feature, len(*in))
		copy(*out, *in)
	}
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]PackageRef, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleSpec.
func (in *BundleSpec) DeepCopy() *BundleSpec {
	if in == nil {
		return nil
	}
	out := new(BundleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleStatus) DeepCopyInto(out *BundleStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleStatus.
func (in *BundleStatus) DeepCopy() *BundleStatus {
	if in == nil {
		return nil
	}
	out := new(BundleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BundleView) DeepCopyInto(out *BundleView) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.BundleOptionView.DeepCopyInto(&out.BundleOptionView)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BundleView.
func (in *BundleView) DeepCopy() *BundleView {
	if in == nil {
		return nil
	}
	out := new(BundleView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BundleView) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartCard) DeepCopyInto(out *ChartCard) {
	*out = *in
	out.ChartRef = in.ChartRef
	in.PackageDescriptor.DeepCopyInto(&out.PackageDescriptor)
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]VersionOption, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartCard.
func (in *ChartCard) DeepCopy() *ChartCard {
	if in == nil {
		return nil
	}
	out := new(ChartCard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartOption) DeepCopyInto(out *ChartOption) {
	*out = *in
	out.ChartRef = in.ChartRef
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]VersionDetail, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartOption.
func (in *ChartOption) DeepCopy() *ChartOption {
	if in == nil {
		return nil
	}
	out := new(ChartOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartRef) DeepCopyInto(out *ChartRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartRef.
func (in *ChartRef) DeepCopy() *ChartRef {
	if in == nil {
		return nil
	}
	out := new(ChartRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartRepoRef) DeepCopyInto(out *ChartRepoRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartRepoRef.
func (in *ChartRepoRef) DeepCopy() *ChartRepoRef {
	if in == nil {
		return nil
	}
	out := new(ChartRepoRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartSelection) DeepCopyInto(out *ChartSelection) {
	*out = *in
	out.ChartRef = in.ChartRef
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(ChartRepoRef)
		**out = **in
	}
	if in.ValuesPatch != nil {
		in, out := &in.ValuesPatch, &out.ValuesPatch
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceDefinitions)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitFors != nil {
		in, out := &in.WaitFors, &out.WaitFors
		*out = make([]WaitFlags, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartSelection.
func (in *ChartSelection) DeepCopy() *ChartSelection {
	if in == nil {
		return nil
	}
	out := new(ChartSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContactData) DeepCopyInto(out *ContactData) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContactData.
func (in *ContactData) DeepCopy() *ContactData {
	if in == nil {
		return nil
	}
	out := new(ContactData)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Feature) DeepCopyInto(out *Feature) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Feature.
func (in *Feature) DeepCopy() *Feature {
	if in == nil {
		return nil
	}
	out := new(Feature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureTable) DeepCopyInto(out *FeatureTable) {
	*out = *in
	if in.Plans != nil {
		in, out := &in.Plans, &out.Plans
		*out = make([]Plan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rows != nil {
		in, out := &in.Rows, &out.Rows
		*out = make([]*Row, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Row)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureTable.
func (in *FeatureTable) DeepCopy() *FeatureTable {
	if in == nil {
		return nil
	}
	out := new(FeatureTable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hub) DeepCopyInto(out *Hub) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hub.
func (in *Hub) DeepCopy() *Hub {
	if in == nil {
		return nil
	}
	out := new(Hub)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallOptions) DeepCopyInto(out *InstallOptions) {
	*out = *in
	out.ChartRef = in.ChartRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallOptions.
func (in *InstallOptions) DeepCopy() *InstallOptions {
	if in == nil {
		return nil
	}
	out := new(InstallOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Link) DeepCopyInto(out *Link) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Link.
func (in *Link) DeepCopy() *Link {
	if in == nil {
		return nil
	}
	out := new(Link)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MediaSpec) DeepCopyInto(out *MediaSpec) {
	*out = *in
	out.ImageSpec = in.ImageSpec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MediaSpec.
func (in *MediaSpec) DeepCopy() *MediaSpec {
	if in == nil {
		return nil
	}
	out := new(MediaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneOfBundleOption) DeepCopyInto(out *OneOfBundleOption) {
	*out = *in
	if in.Bundles != nil {
		in, out := &in.Bundles, &out.Bundles
		*out = make([]*BundleOption, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BundleOption)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneOfBundleOption.
func (in *OneOfBundleOption) DeepCopy() *OneOfBundleOption {
	if in == nil {
		return nil
	}
	out := new(OneOfBundleOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OneOfBundleOptionView) DeepCopyInto(out *OneOfBundleOptionView) {
	*out = *in
	if in.Bundles != nil {
		in, out := &in.Bundles, &out.Bundles
		*out = make([]*BundleOptionView, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(BundleOptionView)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OneOfBundleOptionView.
func (in *OneOfBundleOptionView) DeepCopy() *OneOfBundleOptionView {
	if in == nil {
		return nil
	}
	out := new(OneOfBundleOptionView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Order) DeepCopyInto(out *Order) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Order.
func (in *Order) DeepCopy() *Order {
	if in == nil {
		return nil
	}
	out := new(Order)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Order) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderList) DeepCopyInto(out *OrderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Order, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderList.
func (in *OrderList) DeepCopy() *OrderList {
	if in == nil {
		return nil
	}
	out := new(OrderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OrderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderSpec) DeepCopyInto(out *OrderSpec) {
	*out = *in
	if in.Packages != nil {
		in, out := &in.Packages, &out.Packages
		*out = make([]PackageSelection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderSpec.
func (in *OrderSpec) DeepCopy() *OrderSpec {
	if in == nil {
		return nil
	}
	out := new(OrderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrderStatus) DeepCopyInto(out *OrderStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrderStatus.
func (in *OrderStatus) DeepCopy() *OrderStatus {
	if in == nil {
		return nil
	}
	out := new(OrderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageCard) DeepCopyInto(out *PackageCard) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(ChartCard)
		(*in).DeepCopyInto(*out)
	}
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(BundleOptionView)
		(*in).DeepCopyInto(*out)
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = new(OneOfBundleOptionView)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageCard.
func (in *PackageCard) DeepCopy() *PackageCard {
	if in == nil {
		return nil
	}
	out := new(PackageCard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageDescriptor) DeepCopyInto(out *PackageDescriptor) {
	*out = *in
	if in.Icons != nil {
		in, out := &in.Icons, &out.Icons
		*out = make([]ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]ContactData, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]Link, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageDescriptor.
func (in *PackageDescriptor) DeepCopy() *PackageDescriptor {
	if in == nil {
		return nil
	}
	out := new(PackageDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageMeta) DeepCopyInto(out *PackageMeta) {
	*out = *in
	in.PackageDescriptor.DeepCopyInto(&out.PackageDescriptor)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageMeta.
func (in *PackageMeta) DeepCopy() *PackageMeta {
	if in == nil {
		return nil
	}
	out := new(PackageMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageRef) DeepCopyInto(out *PackageRef) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(ChartOption)
		(*in).DeepCopyInto(*out)
	}
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(BundleOption)
		**out = **in
	}
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = new(OneOfBundleOption)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageRef.
func (in *PackageRef) DeepCopy() *PackageRef {
	if in == nil {
		return nil
	}
	out := new(PackageRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageSelection) DeepCopyInto(out *PackageSelection) {
	*out = *in
	if in.Chart != nil {
		in, out := &in.Chart, &out.Chart
		*out = new(ChartSelection)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageSelection.
func (in *PackageSelection) DeepCopy() *PackageSelection {
	if in == nil {
		return nil
	}
	out := new(PackageSelection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PackageView) DeepCopyInto(out *PackageView) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.PackageMeta.DeepCopyInto(&out.PackageMeta)
	if in.ValuesFiles != nil {
		in, out := &in.ValuesFiles, &out.ValuesFiles
		*out = make([]ValuesFile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.OpenAPIV3Schema != nil {
		in, out := &in.OpenAPIV3Schema, &out.OpenAPIV3Schema
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PackageView.
func (in *PackageView) DeepCopy() *PackageView {
	if in == nil {
		return nil
	}
	out := new(PackageView)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PackageView) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plan) DeepCopyInto(out *Plan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plan.
func (in *Plan) DeepCopy() *Plan {
	if in == nil {
		return nil
	}
	out := new(Plan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Plan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanList) DeepCopyInto(out *PlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Plan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanList.
func (in *PlanList) DeepCopy() *PlanList {
	if in == nil {
		return nil
	}
	out := new(PlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanSpec) DeepCopyInto(out *PlanSpec) {
	*out = *in
	if in.Bundle != nil {
		in, out := &in.Bundle, &out.Bundle
		*out = new(ChartRef)
		**out = **in
	}
	if in.IncludedPlans != nil {
		in, out := &in.IncludedPlans, &out.IncludedPlans
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PricingPattern != nil {
		in, out := &in.PricingPattern, &out.PricingPattern
		*out = make(map[ResourceGroup]PricingPattern, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.AggregateUsage != nil {
		in, out := &in.AggregateUsage, &out.AggregateUsage
		*out = new(string)
		**out = **in
	}
	if in.Amount != nil {
		in, out := &in.Amount, &out.Amount
		*out = new(int64)
		**out = **in
	}
	if in.AmountDecimal != nil {
		in, out := &in.AmountDecimal, &out.AmountDecimal
		*out = new(float64)
		**out = **in
	}
	if in.BillingScheme != nil {
		in, out := &in.BillingScheme, &out.BillingScheme
		*out = new(string)
		**out = **in
	}
	if in.Currency != nil {
		in, out := &in.Currency, &out.Currency
		*out = new(string)
		**out = **in
	}
	if in.Interval != nil {
		in, out := &in.Interval, &out.Interval
		*out = new(string)
		**out = **in
	}
	if in.IntervalCount != nil {
		in, out := &in.IntervalCount, &out.IntervalCount
		*out = new(int64)
		**out = **in
	}
	if in.Tiers != nil {
		in, out := &in.Tiers, &out.Tiers
		*out = make([]*PlanTier, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PlanTier)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.TiersMode != nil {
		in, out := &in.TiersMode, &out.TiersMode
		*out = new(string)
		**out = **in
	}
	if in.TransformUsage != nil {
		in, out := &in.TransformUsage, &out.TransformUsage
		*out = new(PlanTransformUsage)
		(*in).DeepCopyInto(*out)
	}
	if in.TrialPeriodDays != nil {
		in, out := &in.TrialPeriodDays, &out.TrialPeriodDays
		*out = new(int64)
		**out = **in
	}
	if in.UsageType != nil {
		in, out := &in.UsageType, &out.UsageType
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanSpec.
func (in *PlanSpec) DeepCopy() *PlanSpec {
	if in == nil {
		return nil
	}
	out := new(PlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanStatus) DeepCopyInto(out *PlanStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanStatus.
func (in *PlanStatus) DeepCopy() *PlanStatus {
	if in == nil {
		return nil
	}
	out := new(PlanStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanTier) DeepCopyInto(out *PlanTier) {
	*out = *in
	if in.FlatAmount != nil {
		in, out := &in.FlatAmount, &out.FlatAmount
		*out = new(int64)
		**out = **in
	}
	if in.FlatAmountDecimal != nil {
		in, out := &in.FlatAmountDecimal, &out.FlatAmountDecimal
		*out = new(float64)
		**out = **in
	}
	if in.UnitAmount != nil {
		in, out := &in.UnitAmount, &out.UnitAmount
		*out = new(int64)
		**out = **in
	}
	if in.UnitAmountDecimal != nil {
		in, out := &in.UnitAmountDecimal, &out.UnitAmountDecimal
		*out = new(float64)
		**out = **in
	}
	if in.UpTo != nil {
		in, out := &in.UpTo, &out.UpTo
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanTier.
func (in *PlanTier) DeepCopy() *PlanTier {
	if in == nil {
		return nil
	}
	out := new(PlanTier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlanTransformUsage) DeepCopyInto(out *PlanTransformUsage) {
	*out = *in
	if in.DivideBy != nil {
		in, out := &in.DivideBy, &out.DivideBy
		*out = new(int64)
		**out = **in
	}
	if in.Round != nil {
		in, out := &in.Round, &out.Round
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlanTransformUsage.
func (in *PlanTransformUsage) DeepCopy() *PlanTransformUsage {
	if in == nil {
		return nil
	}
	out := new(PlanTransformUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PricingPattern) DeepCopyInto(out *PricingPattern) {
	*out = *in
	if in.SizedPrices != nil {
		in, out := &in.SizedPrices, &out.SizedPrices
		*out = make([]SizedPrice, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PricingPattern.
func (in *PricingPattern) DeepCopy() *PricingPattern {
	if in == nil {
		return nil
	}
	out := new(PricingPattern)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Product) DeepCopyInto(out *Product) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Product.
func (in *Product) DeepCopy() *Product {
	if in == nil {
		return nil
	}
	out := new(Product)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Product) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductList) DeepCopyInto(out *ProductList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Product, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductList.
func (in *ProductList) DeepCopy() *ProductList {
	if in == nil {
		return nil
	}
	out := new(ProductList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProductList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductSpec) DeepCopyInto(out *ProductSpec) {
	*out = *in
	if in.Media != nil {
		in, out := &in.Media, &out.Media
		*out = make([]MediaSpec, len(*in))
		copy(*out, *in)
	}
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]ContactData, len(*in))
		copy(*out, *in)
	}
	if in.Keywords != nil {
		in, out := &in.Keywords, &out.Keywords
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = make([]Link, len(*in))
		copy(*out, *in)
	}
	if in.Badges != nil {
		in, out := &in.Badges, &out.Badges
		*out = make([]Badge, len(*in))
		copy(*out, *in)
	}
	if in.Versions != nil {
		in, out := &in.Versions, &out.Versions
		*out = make([]ProductVersion, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductSpec.
func (in *ProductSpec) DeepCopy() *ProductSpec {
	if in == nil {
		return nil
	}
	out := new(ProductSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductStatus) DeepCopyInto(out *ProductStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductStatus.
func (in *ProductStatus) DeepCopy() *ProductStatus {
	if in == nil {
		return nil
	}
	out := new(ProductStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProductVersion) DeepCopyInto(out *ProductVersion) {
	*out = *in
	if in.ReleaseDate != nil {
		in, out := &in.ReleaseDate, &out.ReleaseDate
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProductVersion.
func (in *ProductVersion) DeepCopy() *ProductVersion {
	if in == nil {
		return nil
	}
	out := new(ProductVersion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	if in.Maintainers != nil {
		in, out := &in.Maintainers, &out.Maintainers
		*out = make([]ContactData, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceDefinitions) DeepCopyInto(out *ResourceDefinitions) {
	*out = *in
	if in.Owned != nil {
		in, out := &in.Owned, &out.Owned
		*out = make([]v1.GroupVersionResource, len(*in))
		copy(*out, *in)
	}
	if in.Required != nil {
		in, out := &in.Required, &out.Required
		*out = make([]v1.GroupVersionResource, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceDefinitions.
func (in *ResourceDefinitions) DeepCopy() *ResourceDefinitions {
	if in == nil {
		return nil
	}
	out := new(ResourceDefinitions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Row) DeepCopyInto(out *Row) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Row.
func (in *Row) DeepCopy() *Row {
	if in == nil {
		return nil
	}
	out := new(Row)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SizedPrice) DeepCopyInto(out *SizedPrice) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SizedPrice.
func (in *SizedPrice) DeepCopy() *SizedPrice {
	if in == nil {
		return nil
	}
	out := new(SizedPrice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValuesFile) DeepCopyInto(out *ValuesFile) {
	*out = *in
	if in.Values != nil {
		in, out := &in.Values, &out.Values
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValuesFile.
func (in *ValuesFile) DeepCopy() *ValuesFile {
	if in == nil {
		return nil
	}
	out := new(ValuesFile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionDetail) DeepCopyInto(out *VersionDetail) {
	*out = *in
	in.VersionOption.DeepCopyInto(&out.VersionOption)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(ResourceDefinitions)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitFors != nil {
		in, out := &in.WaitFors, &out.WaitFors
		*out = make([]WaitFlags, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionDetail.
func (in *VersionDetail) DeepCopy() *VersionDetail {
	if in == nil {
		return nil
	}
	out := new(VersionDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VersionOption) DeepCopyInto(out *VersionOption) {
	*out = *in
	if in.ValuesPatch != nil {
		in, out := &in.ValuesPatch, &out.ValuesPatch
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VersionOption.
func (in *VersionOption) DeepCopy() *VersionOption {
	if in == nil {
		return nil
	}
	out := new(VersionOption)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WaitFlags) DeepCopyInto(out *WaitFlags) {
	*out = *in
	out.Resource = in.Resource
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.Timeout = in.Timeout
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WaitFlags.
func (in *WaitFlags) DeepCopy() *WaitFlags {
	if in == nil {
		return nil
	}
	out := new(WaitFlags)
	in.DeepCopyInto(out)
	return out
}
