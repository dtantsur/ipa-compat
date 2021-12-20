package ipa

import (
	"github.com/gophercloud/gophercloud"
	ironicNoauth "github.com/gophercloud/gophercloud/openstack/baremetal/noauth"
	inspectorNoauth "github.com/gophercloud/gophercloud/openstack/baremetalintrospection/noauth"
	"github.com/pkg/errors"
)

// An Ironic client required for the agent to operate.
type Client struct {
	ironic    *gophercloud.ServiceClient
	inspector *gophercloud.ServiceClient
}

// Create a new client, wrapping Ironic and Inspector clients.
func NewClient(ironicEndpoint string, inspectorEndpoint string) (Client, error) {
	ironic, err := ironicNoauth.NewBareMetalNoAuth(ironicNoauth.EndpointOpts{
		IronicEndpoint: ironicEndpoint,
	})
	if err != nil {
		return Client{}, errors.Wrap(err, "failed to create an Ironic client")
	}
	ironic.Microversion = "1.62"

	inspector, err := inspectorNoauth.NewBareMetalIntrospectionNoAuth(inspectorNoauth.EndpointOpts{
		IronicInspectorEndpoint: inspectorEndpoint,
	})
	if err != nil {
		return Client{}, errors.Wrap(err, "failed to create an Inspector client")
	}

	return Client{
		ironic:    ironic,
		inspector: inspector,
	}, nil
}
