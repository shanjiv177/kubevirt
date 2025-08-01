//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
This file is part of the KubeVirt project

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Copyright The KubeVirt Authors.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kubevirt.io/api/core/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUInstancetype) DeepCopyInto(out *CPUInstancetype) {
	*out = *in
	if in.NUMA != nil {
		in, out := &in.NUMA, &out.NUMA
		*out = new(v1.NUMA)
		(*in).DeepCopyInto(*out)
	}
	if in.Realtime != nil {
		in, out := &in.Realtime, &out.Realtime
		*out = new(v1.Realtime)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUInstancetype.
func (in *CPUInstancetype) DeepCopy() *CPUInstancetype {
	if in == nil {
		return nil
	}
	out := new(CPUInstancetype)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUPreferences) DeepCopyInto(out *CPUPreferences) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUPreferences.
func (in *CPUPreferences) DeepCopy() *CPUPreferences {
	if in == nil {
		return nil
	}
	out := new(CPUPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClockPreferences) DeepCopyInto(out *ClockPreferences) {
	*out = *in
	if in.PreferredClockOffset != nil {
		in, out := &in.PreferredClockOffset, &out.PreferredClockOffset
		*out = new(v1.ClockOffset)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredTimer != nil {
		in, out := &in.PreferredTimer, &out.PreferredTimer
		*out = new(v1.Timer)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClockPreferences.
func (in *ClockPreferences) DeepCopy() *ClockPreferences {
	if in == nil {
		return nil
	}
	out := new(ClockPreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DevicePreferences) DeepCopyInto(out *DevicePreferences) {
	*out = *in
	if in.PreferredAutoattachGraphicsDevice != nil {
		in, out := &in.PreferredAutoattachGraphicsDevice, &out.PreferredAutoattachGraphicsDevice
		*out = new(bool)
		**out = **in
	}
	if in.PreferredAutoattachMemBalloon != nil {
		in, out := &in.PreferredAutoattachMemBalloon, &out.PreferredAutoattachMemBalloon
		*out = new(bool)
		**out = **in
	}
	if in.PreferredAutoattachPodInterface != nil {
		in, out := &in.PreferredAutoattachPodInterface, &out.PreferredAutoattachPodInterface
		*out = new(bool)
		**out = **in
	}
	if in.PreferredAutoattachSerialConsole != nil {
		in, out := &in.PreferredAutoattachSerialConsole, &out.PreferredAutoattachSerialConsole
		*out = new(bool)
		**out = **in
	}
	if in.PreferredAutoattachInputDevice != nil {
		in, out := &in.PreferredAutoattachInputDevice, &out.PreferredAutoattachInputDevice
		*out = new(bool)
		**out = **in
	}
	if in.PreferredDisableHotplug != nil {
		in, out := &in.PreferredDisableHotplug, &out.PreferredDisableHotplug
		*out = new(bool)
		**out = **in
	}
	if in.PreferredVirtualGPUOptions != nil {
		in, out := &in.PreferredVirtualGPUOptions, &out.PreferredVirtualGPUOptions
		*out = new(v1.VGPUOptions)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredUseVirtioTransitional != nil {
		in, out := &in.PreferredUseVirtioTransitional, &out.PreferredUseVirtioTransitional
		*out = new(bool)
		**out = **in
	}
	if in.PreferredDiskDedicatedIoThread != nil {
		in, out := &in.PreferredDiskDedicatedIoThread, &out.PreferredDiskDedicatedIoThread
		*out = new(bool)
		**out = **in
	}
	if in.PreferredDiskBlockSize != nil {
		in, out := &in.PreferredDiskBlockSize, &out.PreferredDiskBlockSize
		*out = new(v1.BlockSize)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredRng != nil {
		in, out := &in.PreferredRng, &out.PreferredRng
		*out = new(v1.Rng)
		**out = **in
	}
	if in.PreferredBlockMultiQueue != nil {
		in, out := &in.PreferredBlockMultiQueue, &out.PreferredBlockMultiQueue
		*out = new(bool)
		**out = **in
	}
	if in.PreferredNetworkInterfaceMultiQueue != nil {
		in, out := &in.PreferredNetworkInterfaceMultiQueue, &out.PreferredNetworkInterfaceMultiQueue
		*out = new(bool)
		**out = **in
	}
	if in.PreferredTPM != nil {
		in, out := &in.PreferredTPM, &out.PreferredTPM
		*out = new(v1.TPMDevice)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DevicePreferences.
func (in *DevicePreferences) DeepCopy() *DevicePreferences {
	if in == nil {
		return nil
	}
	out := new(DevicePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturePreferences) DeepCopyInto(out *FeaturePreferences) {
	*out = *in
	if in.PreferredAcpi != nil {
		in, out := &in.PreferredAcpi, &out.PreferredAcpi
		*out = new(v1.FeatureState)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredApic != nil {
		in, out := &in.PreferredApic, &out.PreferredApic
		*out = new(v1.FeatureAPIC)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredHyperv != nil {
		in, out := &in.PreferredHyperv, &out.PreferredHyperv
		*out = new(v1.FeatureHyperv)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredKvm != nil {
		in, out := &in.PreferredKvm, &out.PreferredKvm
		*out = new(v1.FeatureKVM)
		**out = **in
	}
	if in.PreferredPvspinlock != nil {
		in, out := &in.PreferredPvspinlock, &out.PreferredPvspinlock
		*out = new(v1.FeatureState)
		(*in).DeepCopyInto(*out)
	}
	if in.PreferredSmm != nil {
		in, out := &in.PreferredSmm, &out.PreferredSmm
		*out = new(v1.FeatureState)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturePreferences.
func (in *FeaturePreferences) DeepCopy() *FeaturePreferences {
	if in == nil {
		return nil
	}
	out := new(FeaturePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirmwarePreferences) DeepCopyInto(out *FirmwarePreferences) {
	*out = *in
	if in.PreferredUseBios != nil {
		in, out := &in.PreferredUseBios, &out.PreferredUseBios
		*out = new(bool)
		**out = **in
	}
	if in.PreferredUseBiosSerial != nil {
		in, out := &in.PreferredUseBiosSerial, &out.PreferredUseBiosSerial
		*out = new(bool)
		**out = **in
	}
	if in.PreferredUseEfi != nil {
		in, out := &in.PreferredUseEfi, &out.PreferredUseEfi
		*out = new(bool)
		**out = **in
	}
	if in.PreferredUseSecureBoot != nil {
		in, out := &in.PreferredUseSecureBoot, &out.PreferredUseSecureBoot
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirmwarePreferences.
func (in *FirmwarePreferences) DeepCopy() *FirmwarePreferences {
	if in == nil {
		return nil
	}
	out := new(FirmwarePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePreferences) DeepCopyInto(out *MachinePreferences) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePreferences.
func (in *MachinePreferences) DeepCopy() *MachinePreferences {
	if in == nil {
		return nil
	}
	out := new(MachinePreferences)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemoryInstancetype) DeepCopyInto(out *MemoryInstancetype) {
	*out = *in
	out.Guest = in.Guest.DeepCopy()
	if in.Hugepages != nil {
		in, out := &in.Hugepages, &out.Hugepages
		*out = new(v1.Hugepages)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemoryInstancetype.
func (in *MemoryInstancetype) DeepCopy() *MemoryInstancetype {
	if in == nil {
		return nil
	}
	out := new(MemoryInstancetype)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterInstancetype) DeepCopyInto(out *VirtualMachineClusterInstancetype) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterInstancetype.
func (in *VirtualMachineClusterInstancetype) DeepCopy() *VirtualMachineClusterInstancetype {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterInstancetype)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterInstancetype) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterInstancetypeList) DeepCopyInto(out *VirtualMachineClusterInstancetypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineClusterInstancetype, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterInstancetypeList.
func (in *VirtualMachineClusterInstancetypeList) DeepCopy() *VirtualMachineClusterInstancetypeList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterInstancetypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterInstancetypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterPreference) DeepCopyInto(out *VirtualMachineClusterPreference) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterPreference.
func (in *VirtualMachineClusterPreference) DeepCopy() *VirtualMachineClusterPreference {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterPreference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterPreference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineClusterPreferenceList) DeepCopyInto(out *VirtualMachineClusterPreferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineClusterPreference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineClusterPreferenceList.
func (in *VirtualMachineClusterPreferenceList) DeepCopy() *VirtualMachineClusterPreferenceList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineClusterPreferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineClusterPreferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineInstancetype) DeepCopyInto(out *VirtualMachineInstancetype) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineInstancetype.
func (in *VirtualMachineInstancetype) DeepCopy() *VirtualMachineInstancetype {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineInstancetype)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineInstancetype) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineInstancetypeList) DeepCopyInto(out *VirtualMachineInstancetypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachineInstancetype, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineInstancetypeList.
func (in *VirtualMachineInstancetypeList) DeepCopy() *VirtualMachineInstancetypeList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineInstancetypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachineInstancetypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachineInstancetypeSpec) DeepCopyInto(out *VirtualMachineInstancetypeSpec) {
	*out = *in
	in.CPU.DeepCopyInto(&out.CPU)
	in.Memory.DeepCopyInto(&out.Memory)
	if in.GPUs != nil {
		in, out := &in.GPUs, &out.GPUs
		*out = make([]v1.GPU, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostDevices != nil {
		in, out := &in.HostDevices, &out.HostDevices
		*out = make([]v1.HostDevice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IOThreadsPolicy != nil {
		in, out := &in.IOThreadsPolicy, &out.IOThreadsPolicy
		*out = new(v1.IOThreadsPolicy)
		**out = **in
	}
	if in.LaunchSecurity != nil {
		in, out := &in.LaunchSecurity, &out.LaunchSecurity
		*out = new(v1.LaunchSecurity)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachineInstancetypeSpec.
func (in *VirtualMachineInstancetypeSpec) DeepCopy() *VirtualMachineInstancetypeSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachineInstancetypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePreference) DeepCopyInto(out *VirtualMachinePreference) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePreference.
func (in *VirtualMachinePreference) DeepCopy() *VirtualMachinePreference {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePreference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachinePreference) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePreferenceList) DeepCopyInto(out *VirtualMachinePreferenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualMachinePreference, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePreferenceList.
func (in *VirtualMachinePreferenceList) DeepCopy() *VirtualMachinePreferenceList {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePreferenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualMachinePreferenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualMachinePreferenceSpec) DeepCopyInto(out *VirtualMachinePreferenceSpec) {
	*out = *in
	if in.Clock != nil {
		in, out := &in.Clock, &out.Clock
		*out = new(ClockPreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.CPU != nil {
		in, out := &in.CPU, &out.CPU
		*out = new(CPUPreferences)
		**out = **in
	}
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = new(DevicePreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(FeaturePreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.Firmware != nil {
		in, out := &in.Firmware, &out.Firmware
		*out = new(FirmwarePreferences)
		(*in).DeepCopyInto(*out)
	}
	if in.Machine != nil {
		in, out := &in.Machine, &out.Machine
		*out = new(MachinePreferences)
		**out = **in
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = new(VolumePreferences)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualMachinePreferenceSpec.
func (in *VirtualMachinePreferenceSpec) DeepCopy() *VirtualMachinePreferenceSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualMachinePreferenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumePreferences) DeepCopyInto(out *VolumePreferences) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumePreferences.
func (in *VolumePreferences) DeepCopy() *VolumePreferences {
	if in == nil {
		return nil
	}
	out := new(VolumePreferences)
	in.DeepCopyInto(out)
	return out
}
