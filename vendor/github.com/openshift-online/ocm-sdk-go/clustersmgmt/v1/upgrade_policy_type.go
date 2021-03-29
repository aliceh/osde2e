/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	time "time"
)

// UpgradePolicyKind is the name of the type used to represent objects
// of type 'upgrade_policy'.
const UpgradePolicyKind = "UpgradePolicy"

// UpgradePolicyLinkKind is the name of the type used to represent links
// to objects of type 'upgrade_policy'.
const UpgradePolicyLinkKind = "UpgradePolicyLink"

// UpgradePolicyNilKind is the name of the type used to nil references
// to objects of type 'upgrade_policy'.
const UpgradePolicyNilKind = "UpgradePolicyNil"

// UpgradePolicy represents the values of the 'upgrade_policy' type.
//
// Representation of an upgrade policy that can be set for a cluster.
type UpgradePolicy struct {
	id           *string
	href         *string
	link         bool
	clusterID    *string
	nextRun      *time.Time
	schedule     *string
	scheduleType *string
	upgradeType  *string
	version      *string
}

// Kind returns the name of the type of the object.
func (o *UpgradePolicy) Kind() string {
	if o == nil {
		return UpgradePolicyNilKind
	}
	if o.link {
		return UpgradePolicyLinkKind
	}
	return UpgradePolicyKind
}

// ID returns the identifier of the object.
func (o *UpgradePolicy) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *UpgradePolicy) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *UpgradePolicy) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *UpgradePolicy) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *UpgradePolicy) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *UpgradePolicy) Empty() bool {
	return o == nil || (o.id == nil &&
		o.clusterID == nil &&
		o.nextRun == nil &&
		o.schedule == nil &&
		o.scheduleType == nil &&
		o.upgradeType == nil &&
		o.version == nil &&
		true)
}

// ClusterID returns the value of the 'cluster_ID' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Cluster ID this upgrade policy is defined for.
func (o *UpgradePolicy) ClusterID() string {
	if o != nil && o.clusterID != nil {
		return *o.clusterID
	}
	return ""
}

// GetClusterID returns the value of the 'cluster_ID' attribute and
// a flag indicating if the attribute has a value.
//
// Cluster ID this upgrade policy is defined for.
func (o *UpgradePolicy) GetClusterID() (value string, ok bool) {
	ok = o != nil && o.clusterID != nil
	if ok {
		value = *o.clusterID
	}
	return
}

// NextRun returns the value of the 'next_run' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Next time the upgrade should run.
func (o *UpgradePolicy) NextRun() time.Time {
	if o != nil && o.nextRun != nil {
		return *o.nextRun
	}
	return time.Time{}
}

// GetNextRun returns the value of the 'next_run' attribute and
// a flag indicating if the attribute has a value.
//
// Next time the upgrade should run.
func (o *UpgradePolicy) GetNextRun() (value time.Time, ok bool) {
	ok = o != nil && o.nextRun != nil
	if ok {
		value = *o.nextRun
	}
	return
}

// Schedule returns the value of the 'schedule' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Schedule cron expression that defines automatic upgrade scheduling.
func (o *UpgradePolicy) Schedule() string {
	if o != nil && o.schedule != nil {
		return *o.schedule
	}
	return ""
}

// GetSchedule returns the value of the 'schedule' attribute and
// a flag indicating if the attribute has a value.
//
// Schedule cron expression that defines automatic upgrade scheduling.
func (o *UpgradePolicy) GetSchedule() (value string, ok bool) {
	ok = o != nil && o.schedule != nil
	if ok {
		value = *o.schedule
	}
	return
}

// ScheduleType returns the value of the 'schedule_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Schedule type can be either "manual" (single execution) or "automatic" (re-occurring).
func (o *UpgradePolicy) ScheduleType() string {
	if o != nil && o.scheduleType != nil {
		return *o.scheduleType
	}
	return ""
}

// GetScheduleType returns the value of the 'schedule_type' attribute and
// a flag indicating if the attribute has a value.
//
// Schedule type can be either "manual" (single execution) or "automatic" (re-occurring).
func (o *UpgradePolicy) GetScheduleType() (value string, ok bool) {
	ok = o != nil && o.scheduleType != nil
	if ok {
		value = *o.scheduleType
	}
	return
}

// UpgradeType returns the value of the 'upgrade_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Upgrade type specify the type of the upgrade. Can be "OSD" or "CVE".
func (o *UpgradePolicy) UpgradeType() string {
	if o != nil && o.upgradeType != nil {
		return *o.upgradeType
	}
	return ""
}

// GetUpgradeType returns the value of the 'upgrade_type' attribute and
// a flag indicating if the attribute has a value.
//
// Upgrade type specify the type of the upgrade. Can be "OSD" or "CVE".
func (o *UpgradePolicy) GetUpgradeType() (value string, ok bool) {
	ok = o != nil && o.upgradeType != nil
	if ok {
		value = *o.upgradeType
	}
	return
}

// Version returns the value of the 'version' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Version is the desired upgrade version.
func (o *UpgradePolicy) Version() string {
	if o != nil && o.version != nil {
		return *o.version
	}
	return ""
}

// GetVersion returns the value of the 'version' attribute and
// a flag indicating if the attribute has a value.
//
// Version is the desired upgrade version.
func (o *UpgradePolicy) GetVersion() (value string, ok bool) {
	ok = o != nil && o.version != nil
	if ok {
		value = *o.version
	}
	return
}

// UpgradePolicyListKind is the name of the type used to represent list of objects of
// type 'upgrade_policy'.
const UpgradePolicyListKind = "UpgradePolicyList"

// UpgradePolicyListLinkKind is the name of the type used to represent links to list
// of objects of type 'upgrade_policy'.
const UpgradePolicyListLinkKind = "UpgradePolicyListLink"

// UpgradePolicyNilKind is the name of the type used to nil lists of objects of
// type 'upgrade_policy'.
const UpgradePolicyListNilKind = "UpgradePolicyListNil"

// UpgradePolicyList is a list of values of the 'upgrade_policy' type.
type UpgradePolicyList struct {
	href  *string
	link  bool
	items []*UpgradePolicy
}

// Kind returns the name of the type of the object.
func (l *UpgradePolicyList) Kind() string {
	if l == nil {
		return UpgradePolicyListNilKind
	}
	if l.link {
		return UpgradePolicyListLinkKind
	}
	return UpgradePolicyListKind
}

// Link returns true iif this is a link.
func (l *UpgradePolicyList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *UpgradePolicyList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *UpgradePolicyList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *UpgradePolicyList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *UpgradePolicyList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *UpgradePolicyList) Get(i int) *UpgradePolicy {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *UpgradePolicyList) Slice() []*UpgradePolicy {
	var slice []*UpgradePolicy
	if l == nil {
		slice = make([]*UpgradePolicy, 0)
	} else {
		slice = make([]*UpgradePolicy, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *UpgradePolicyList) Each(f func(item *UpgradePolicy) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *UpgradePolicyList) Range(f func(index int, item *UpgradePolicy) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}