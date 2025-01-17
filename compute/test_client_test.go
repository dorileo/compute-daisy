//  Copyright 2017 Google Inc. All Rights Reserved.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package compute

import (
	"fmt"
	"net/http"
	"testing"

	"google.golang.org/api/compute/v1"
	"google.golang.org/api/googleapi"
)

func TestTestClient(t *testing.T) {
	var fakeCalled, realCalled bool
	var wantFakeCalled, wantRealCalled bool
	var url string
	listOpts := []ListCallOption{Filter("foo"), OrderBy("foo")}
	_, c, _ := NewTestClient(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		realCalled = true
		url = r.URL.String()
		w.WriteHeader(400)
		fmt.Fprintln(w, "Not Implemented")
	}))

	tests := []struct {
		desc string
		op   func()
		wURL string
	}{
		{"retry", func() {
			c.Retry(func(_ ...googleapi.CallOption) (*compute.Operation, error) { realCalled = true; return nil, nil })
		}, ""},
		{"attach disk", func() { c.AttachDisk("a", "b", "c", &compute.AttachedDisk{}) }, "/projects/a/zones/b/instances/c/attachDisk?alt=json&prettyPrint=false"},
		{"detach disk", func() { c.DetachDisk("a", "b", "c", "d") }, "/projects/a/zones/b/instances/c/detachDisk?alt=json&deviceName=d&prettyPrint=false"},
		{"resize disk", func() { c.ResizeDisk("a", "b", "c", &compute.DisksResizeRequest{SizeGb: 128}) }, "/projects/a/zones/b/disks/c/resize?alt=json&prettyPrint=false"},
		{"create disk", func() { c.CreateDisk("a", "b", &compute.Disk{}) }, "/projects/a/zones/b/disks?alt=json&prettyPrint=false"},
		{"create firewall rule", func() { c.CreateFirewallRule("a", &compute.Firewall{}) }, "/projects/a/global/firewalls?alt=json&prettyPrint=false"},
		{"create image", func() { c.CreateImage("a", &compute.Image{}) }, "/projects/a/global/images?alt=json&prettyPrint=false"},
		{"create instance", func() { c.CreateInstance("a", "b", &compute.Instance{}) }, "/projects/a/zones/b/instances?alt=json&prettyPrint=false"},
		{"create network", func() { c.CreateNetwork("a", &compute.Network{}) }, "/projects/a/global/networks?alt=json&prettyPrint=false"},
		{"create subnetwork", func() { c.CreateSubnetwork("a", "b", &compute.Subnetwork{}) }, "/projects/a/regions/b/subnetworks?alt=json&prettyPrint=false"},
		{"instances start", func() { c.StartInstance("a", "b", "c") }, "/projects/a/zones/b/instances/c/start?alt=json&prettyPrint=false"},
		{"instances stop", func() { c.StopInstance("a", "b", "c") }, "/projects/a/zones/b/instances/c/stop?alt=json&prettyPrint=false"},
		{"delete disk", func() { c.DeleteDisk("a", "b", "c") }, "/projects/a/zones/b/disks/c?alt=json&prettyPrint=false"},
		{"delete firewall rule", func() { c.DeleteFirewallRule("a", "b") }, "/projects/a/global/firewalls/b?alt=json&prettyPrint=false"},
		{"delete image", func() { c.DeleteImage("a", "b") }, "/projects/a/global/images/b?alt=json&prettyPrint=false"},
		{"delete instance", func() { c.DeleteInstance("a", "b", "c") }, "/projects/a/zones/b/instances/c?alt=json&prettyPrint=false"},
		{"delete network", func() { c.DeleteNetwork("a", "b") }, "/projects/a/global/networks/b?alt=json&prettyPrint=false"},
		{"delete subnetwork", func() { c.DeleteSubnetwork("a", "b", "c") }, "/projects/a/regions/b/subnetworks/c?alt=json&prettyPrint=false"},
		{"deprecate image", func() { c.DeprecateImage("a", "b", &compute.DeprecationStatus{}) }, "/projects/a/global/images/b/deprecate?alt=json&prettyPrint=false"},
		{"get serial port", func() { c.GetSerialPortOutput("a", "b", "c", 1, 2) }, "/projects/a/zones/b/instances/c/serialPort?alt=json&port=1&prettyPrint=false&start=2"},
		{"get project", func() { c.GetProject("a") }, "/projects/a?alt=json&prettyPrint=false"},
		{"get machine type", func() { c.GetMachineType("a", "b", "c") }, "/projects/a/zones/b/machineTypes/c?alt=json&prettyPrint=false"},
		{"list machine types", func() { c.ListMachineTypes("a", "b", listOpts...) }, "/projects/a/zones/b/machineTypes?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"get firewall rule", func() { c.GetFirewallRule("a", "b") }, "/projects/a/global/firewalls/b?alt=json&prettyPrint=false"},
		{"list firewall rules", func() { c.ListFirewallRules("a", listOpts...) }, "/projects/a/global/firewalls?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"get zone", func() { c.GetZone("a", "b") }, "/projects/a/zones/b?alt=json&prettyPrint=false"},
		{"list zones", func() { c.ListZones("a", listOpts...) }, "/projects/a/zones?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"get instance", func() { c.GetInstance("a", "b", "c") }, "/projects/a/zones/b/instances/c?alt=json&prettyPrint=false"},
		{"aggregated list instances", func() { c.AggregatedListInstances("a", listOpts...) }, "/projects/a/aggregated/instances?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"list instances", func() { c.ListInstances("a", "b", listOpts...) }, "/projects/a/zones/b/instances?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"get image from family", func() { c.GetImageFromFamily("a", "b") }, "/projects/a/global/images/family/b?alt=json&prettyPrint=false"},
		{"get image", func() { c.GetImage("a", "b") }, "/projects/a/global/images/b?alt=json&prettyPrint=false"},
		{"list images", func() { c.ListImages("a", listOpts...) }, "/projects/a/global/images?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"get license", func() { c.GetLicense("a", "b") }, "/projects/a/global/licenses/b?alt=json&prettyPrint=false"},
		{"get network", func() { c.GetNetwork("a", "b") }, "/projects/a/global/networks/b?alt=json&prettyPrint=false"},
		{"list networks", func() { c.ListNetworks("a", listOpts...) }, "/projects/a/global/networks?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"get subnetwork", func() { c.GetSubnetwork("a", "b", "c") }, "/projects/a/regions/b/subnetworks/c?alt=json&prettyPrint=false"},
		{"aggregated list subnetworks", func() { c.AggregatedListSubnetworks("a", listOpts...) }, "/projects/a/aggregated/subnetworks?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"list subnetworks", func() { c.ListSubnetworks("a", "b", listOpts...) }, "/projects/a/regions/b/subnetworks?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"get disk", func() { c.GetDisk("a", "b", "c") }, "/projects/a/zones/b/disks/c?alt=json&prettyPrint=false"},
		{"aggregated list disks", func() { c.AggregatedListDisks("a", listOpts...) }, "/projects/a/aggregated/disks?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"list disks", func() { c.ListDisks("a", "b", listOpts...) }, "/projects/a/zones/b/disks?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"instance status", func() { c.InstanceStatus("a", "b", "c") }, "/projects/a/zones/b/instances/c?alt=json&prettyPrint=false"},
		{"instance stopped", func() { c.InstanceStopped("a", "b", "c") }, "/projects/a/zones/b/instances/c?alt=json&prettyPrint=false"},
		{"set instance metadata", func() { c.SetInstanceMetadata("a", "b", "c", nil) }, "/projects/a/zones/b/instances/c/setMetadata?alt=json&prettyPrint=false"},
		{"set project metadata", func() { c.SetCommonInstanceMetadata("a", nil) }, "/projects/a/setCommonInstanceMetadata?alt=json&prettyPrint=false"},
		{"zone operation wait", func() { c.zoneOperationsWait("a", "b", "c") }, "/projects/a/zones/b/operations/c/wait?alt=json&prettyPrint=false"},
		{"region operation wait", func() { c.regionOperationsWait("a", "b", "c") }, "/projects/a/regions/b/operations/c/wait?alt=json&prettyPrint=false"},
		{"global operation wait", func() { c.globalOperationsWait("a", "b") }, "/projects/a/global/operations/b/wait?alt=json&prettyPrint=false"},
		{"get guest attributes", func() { c.GetGuestAttributes("a", "b", "c", "d", "e") }, "/projects/a/zones/b/instances/c/getGuestAttributes?alt=json&prettyPrint=false&queryPath=d&variableKey=e"},
		{"create machine image", func() { c.CreateMachineImage("a", &compute.MachineImage{}) }, "/projects/a/global/machineImages?alt=json&prettyPrint=false"},
		{"get machine image", func() { c.GetMachineImage("a", "b") }, "/projects/a/global/machineImages/b?alt=json&prettyPrint=false"},
		{"list machine images", func() { c.ListMachineImages("a", listOpts...) }, "/projects/a/global/machineImages?alt=json&filter=foo&orderBy=foo&pageToken=&prettyPrint=false"},
		{"delete machine image", func() { c.DeleteMachineImage("a", "b") }, "/projects/a/global/machineImages/b?alt=json&prettyPrint=false"},
	}

	runTests := func() {
		for _, tt := range tests {
			fakeCalled = false
			realCalled = false
			url = ""
			tt.op()
			if fakeCalled != wantFakeCalled || realCalled != wantRealCalled {
				t.Errorf("%s case: incorrect calls: wanted fakeCalled=%t realCalled=%t; got fakeCalled=%t realCalled=%t", tt.desc, wantFakeCalled, wantRealCalled, fakeCalled, realCalled)
			}
			if wantRealCalled {
				if tt.wURL != url {
					t.Errorf("%s case: want called url not equal to actual url, want: %q, got: %q", tt.desc, tt.wURL, url)
				}
			}
		}
	}

	// Test real methods can be called.
	wantFakeCalled = false
	wantRealCalled = true
	runTests()

	// Test fake methods can be called.
	c.RetryFn = func(_ func(_ ...googleapi.CallOption) (*compute.Operation, error), _ ...googleapi.CallOption) (op *compute.Operation, err error) {
		fakeCalled = true
		return nil, nil
	}
	c.AttachDiskFn = func(_, _, _ string, _ *compute.AttachedDisk) error { fakeCalled = true; return nil }
	c.DetachDiskFn = func(_, _, _, _ string) error { fakeCalled = true; return nil }
	c.ResizeDiskFn = func(_, _, _ string, _ *compute.DisksResizeRequest) error { fakeCalled = true; return nil }
	c.CreateDiskFn = func(_, _ string, _ *compute.Disk) error { fakeCalled = true; return nil }
	c.CreateFirewallRuleFn = func(_ string, _ *compute.Firewall) error { fakeCalled = true; return nil }
	c.CreateImageFn = func(_ string, _ *compute.Image) error { fakeCalled = true; return nil }
	c.CreateInstanceFn = func(_, _ string, _ *compute.Instance) error { fakeCalled = true; return nil }
	c.CreateNetworkFn = func(_ string, _ *compute.Network) error { fakeCalled = true; return nil }
	c.CreateSubnetworkFn = func(_, _ string, _ *compute.Subnetwork) error { fakeCalled = true; return nil }
	c.StartInstanceFn = func(_, _, _ string) error { fakeCalled = true; return nil }
	c.StopInstanceFn = func(_, _, _ string) error { fakeCalled = true; return nil }
	c.DeleteDiskFn = func(_, _, _ string) error { fakeCalled = true; return nil }
	c.DeleteFirewallRuleFn = func(_, _ string) error { fakeCalled = true; return nil }
	c.DeleteImageFn = func(_, _ string) error { fakeCalled = true; return nil }
	c.DeleteInstanceFn = func(_, _, _ string) error { fakeCalled = true; return nil }
	c.DeleteNetworkFn = func(_, _ string) error { fakeCalled = true; return nil }
	c.DeleteSubnetworkFn = func(_, _, _ string) error { fakeCalled = true; return nil }
	c.DeprecateImageFn = func(_, _ string, _ *compute.DeprecationStatus) error { fakeCalled = true; return nil }
	c.GetSerialPortOutputFn = func(_, _, _ string, _, _ int64) (*compute.SerialPortOutput, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetProjectFn = func(_ string) (*compute.Project, error) { fakeCalled = true; return nil, nil }
	c.GetZoneFn = func(_, _ string) (*compute.Zone, error) { fakeCalled = true; return nil, nil }
	c.ListZonesFn = func(_ string, _ ...ListCallOption) ([]*compute.Zone, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetFirewallRuleFn = func(_, _ string) (*compute.Firewall, error) { fakeCalled = true; return nil, nil }
	c.ListFirewallRulesFn = func(_ string, _ ...ListCallOption) ([]*compute.Firewall, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetInstanceFn = func(_, _, _ string) (*compute.Instance, error) { fakeCalled = true; return nil, nil }
	c.AggregatedListInstancesFn = func(_ string, _ ...ListCallOption) ([]*compute.Instance, error) {
		fakeCalled = true
		return nil, nil
	}
	c.ListInstancesFn = func(_, _ string, _ ...ListCallOption) ([]*compute.Instance, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetDiskFn = func(_, _, _ string) (*compute.Disk, error) { fakeCalled = true; return nil, nil }
	c.AggregatedListDisksFn = func(_ string, _ ...ListCallOption) ([]*compute.Disk, error) {
		fakeCalled = true
		return nil, nil
	}
	c.ListDisksFn = func(_, _ string, _ ...ListCallOption) ([]*compute.Disk, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetImageFromFamilyFn = func(_, _ string) (*compute.Image, error) { fakeCalled = true; return nil, nil }
	c.GetImageFn = func(_, _ string) (*compute.Image, error) { fakeCalled = true; return nil, nil }
	c.ListImagesFn = func(_ string, _ ...ListCallOption) ([]*compute.Image, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetLicenseFn = func(_, _ string) (*compute.License, error) { fakeCalled = true; return nil, nil }
	c.GetNetworkFn = func(_, _ string) (*compute.Network, error) { fakeCalled = true; return nil, nil }
	c.ListNetworksFn = func(_ string, _ ...ListCallOption) ([]*compute.Network, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetSubnetworkFn = func(_, _, _ string) (*compute.Subnetwork, error) { fakeCalled = true; return nil, nil }
	c.AggregatedListSubnetworksFn = func(_ string, _ ...ListCallOption) ([]*compute.Subnetwork, error) {
		fakeCalled = true
		return nil, nil
	}
	c.ListSubnetworksFn = func(_, _ string, _ ...ListCallOption) ([]*compute.Subnetwork, error) {
		fakeCalled = true
		return nil, nil
	}
	c.GetMachineTypeFn = func(_, _, _ string) (*compute.MachineType, error) { fakeCalled = true; return nil, nil }
	c.ListMachineTypesFn = func(_, _ string, _ ...ListCallOption) ([]*compute.MachineType, error) {
		fakeCalled = true
		return nil, nil
	}
	c.InstanceStatusFn = func(_, _, _ string) (string, error) { fakeCalled = true; return "", nil }
	c.InstanceStoppedFn = func(_, _, _ string) (bool, error) { fakeCalled = true; return false, nil }
	c.SetInstanceMetadataFn = func(_, _, _ string, _ *compute.Metadata) error { fakeCalled = true; return nil }
	c.SetCommonInstanceMetadataFn = func(_ string, _ *compute.Metadata) error { fakeCalled = true; return nil }
	c.zoneOperationsWaitFn = func(_, _, _ string) error { fakeCalled = true; return nil }
	c.regionOperationsWaitFn = func(_, _, _ string) error { fakeCalled = true; return nil }
	c.globalOperationsWaitFn = func(_, _ string) error { fakeCalled = true; return nil }
	c.GetGuestAttributesFn = func(_, _, _, _, _ string) (*compute.GuestAttributes, error) { fakeCalled = true; return nil, nil }
	c.CreateMachineImageFn = func(_ string, _ *compute.MachineImage) error { fakeCalled = true; return nil }
	c.GetMachineImageFn = func(_, _ string) (*compute.MachineImage, error) { fakeCalled = true; return nil, nil }
	c.ListMachineImagesFn = func(_ string, _ ...ListCallOption) ([]*compute.MachineImage, error) {
		fakeCalled = true
		return nil, nil
	}
	c.DeleteMachineImageFn = func(_, _ string) error { fakeCalled = true; return nil }
	wantFakeCalled = true
	wantRealCalled = false
	runTests()
}
