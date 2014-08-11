package vmClient

import (
	"encoding/xml"
)

type VMDeployment struct {
	XMLName   xml.Name `xml:"Deployment"`
	Xmlns	  	string `xml:"xmlns,attr"`
	Name	  	string
	DeploymentSlot string
	Label 		string
	RoleList RoleList
}

type HostedServiceDeployment struct {
	XMLName   xml.Name `xml:"CreateHostedService"`
	Xmlns	  	string `xml:"xmlns,attr"`
	ServiceName		string
	Label			string
	Description		string
	Location		string
}

type RoleList struct {
	Role	[]Role
}

type Role struct {
	RoleName      string
	RoleType	  string
	ConfigurationSets	ConfigurationSets
	ResourceExtensionReferences ResourceExtensionReferences `xml:",omitempty"`
	OSVirtualHardDisk OSVirtualHardDisk
	RoleSize	  string
	ProvisionGuestAgent bool
}

type ConfigurationSets struct {
	ConfigurationSet []ConfigurationSet
}

type ResourceExtensionReferences struct {
	ResourceExtensionReference	[]ResourceExtensionReference
}

type InputEndpoints struct {
	InputEndpoint []InputEndpoint
}

type SubnetNames struct {
	SubnetName []SubnetName
}

type ResourceExtensionReference struct {
	ReferenceName		string
	Publisher			string
	Name				string
	Version				string
	ResourceExtensionParameterValues ResourceExtensionParameterValues `xml:",omitempty"`
	State				string
}

type ResourceExtensionParameterValues struct {
	ResourceExtensionParameterValue	[]ResourceExtensionParameter
}

type ResourceExtensionParameter struct {
	Key		    string
	Value		string
	Type		string
}

type OSVirtualHardDisk struct {
	MediaLink			string
	SourceImageName		string
}

type ConfigurationSet struct {
	ConfigurationSetType	string
	HostName				string	`xml:",omitempty"`
	UserName				string	`xml:",omitempty"`
	UserPassword			string	`xml:",omitempty"`
	DisableSshPasswordAuthentication bool
	InputEndpoints			InputEndpoints	`xml:",omitempty"`
	SubnetNames				SubnetNames		`xml:",omitempty"`
}

type InputEndpoint struct {
	LocalPort	int
	Name			string
	Port	int
	Protocol		string
}

//!TODO
type SubnetName struct {
}
