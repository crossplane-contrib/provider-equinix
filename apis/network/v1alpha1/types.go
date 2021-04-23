/*
Copyright 2021 The Crossplane Authors.

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

package v1alpha1

import (
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	ne "github.com/equinix/ne-go"
)

var _ ne.Device = ne.Device{}

// DeviceSpec defines the desired state of Device
type DeviceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DeviceParameters `json:"forProvider"`
}

// An DeviceStatus represents the observed state of an Device manager.
type DeviceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DeviceExternalStatus `json:"atProvider"`
}

// DeviceParameters defines the desired state of an AWS Device.
type DeviceParameters struct {
	UUID                *string              `json:"uuid,omitempty"`
	Name                *string              `json:"name,omitempty"`
	TypeCode            *string              `json:"typeCode,omitempty"`
	Status              *string              `json:"status,omitempty"`
	LicenseStatus       *string              `json:"licenseStatus,omitempty"`
	MetroCode           *string              `json:"metroCode,omitempty"`
	IBX                 *string              `json:"ibx,omitempty"`
	Region              *string              `json:"region,omitempty"`
	Throughput          *int                 `json:"throughput,omitempty"`
	ThroughputUnit      *string              `json:"throughputUnit,omitempty"`
	HostName            *string              `json:"hostName,omitempty"`
	PackageCode         *string              `json:"packageCode,omitempty"`
	Version             *string              `json:"version,omitempty"`
	IsBYOL              *bool                `json:"isBYOL,omitempty"`
	LicenseToken        *string              `json:"licenseToken,omitempty"`
	LicenseFile         *string              `json:"licenseFile,omitempty"`
	LicenseFileID       *string              `json:"licenseFileID,omitempty"`
	ACLTemplateUUID     *string              `json:"aclTemplateUUID,omitempty"`
	SSHIPAddress        *string              `json:"sshIPAddress,omitempty"`
	SSHIPFqdn           *string              `json:"sshIPFqdn,omitempty"`
	AccountNumber       *string              `json:"accountNumber,omitempty"`
	Notifications       []string             `json:"notifications,omitempty"`
	PurchaseOrderNumber *string              `json:"purchaseOrderNumber,omitempty"`
	RedundancyType      *string              `json:"redundancyType,omitempty"`
	RedundantUUID       *string              `json:"redundantUUID,omitempty"`
	TermLength          *int                 `json:"termLength,omitempty"`
	AdditionalBandwidth *int                 `json:"additionalBandwidth,omitempty"`
	OrderReference      *string              `json:"orderReference,omitempty"`
	InterfaceCount      *int                 `json:"interfaceCount,omitempty"`
	CoreCount           *int                 `json:"coreCount,omitempty"`
	IsSelfManaged       *bool                `json:"isSelfManaged,omitempty"`
	Interfaces          []DeviceInterface    `json:"interfaces,omitempty"`
	VendorConfiguration map[string]string    `json:"vendorConfiguration,omitempty"`
	UserPublicKey       *DeviceUserPublicKey `json:"userPublicKey,omitempty"`
	ASN                 *int                 `json:"asn,omitempty"`
	ZoneCode            *string              `json:"zoneCode,omitempty"`
}

// DeviceInterface is a Crossplane representation of ne.DeviceInterface
type DeviceInterface struct {
	ID                *int    `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	Status            *string `json:"status,omitempty"`
	OperationalStatus *string `json:"operationalStatus,omitempty"`
	MACAddress        *string `json:"macAddress,omitempty"`
	IPAddress         *string `json:"ipAddress,omitempty"`
	AssignedType      *string `json:"assignedType,omitempty"`
	Type              *string `json:"type,omitempty"`
}

// DeviceUserPublicKey is a Crossplane representation of ne.DeviceUserPublicKey
type DeviceUserPublicKey struct {
	Username *string `json:"username,omitempty"`
	KeyName  *string `json:"keyName,omitempty"`
}

// DeviceExternalStatus is the external status of a Device.
type DeviceExternalStatus struct {
}

// +kubebuilder:object:root=true

// Device is a managed resource that represents an AWS Device Manager.
// +kubebuilder:printcolumn:name="STATUS",type="string",JSONPath=".status.atProvider.status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,equinix}
type Device struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeviceSpec   `json:"spec,omitempty"`
	Status DeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceList contains a list of Device
type DeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Device `json:"items"`
}
