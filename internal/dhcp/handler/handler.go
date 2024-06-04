// Package handler holds the interface that backends implement, handlers take in, and the top level dhcp package passes to handlers.
package handler

import (
	"context"
	"net"

	"github.com/tinkerbell/smee/internal/dhcp/data"
)

// BackendReader is the interface for getting data from a backend.
//
// Backends implement this interface to provide DHCP and Netboot data to the handlers.
type BackendReader interface {
	// Read data (from a backend) based on a mac address
	// and return DHCP headers and options, including netboot info.
	GetByMac(context.Context, net.HardwareAddr) (*data.DHCP, *data.Netboot, error)
	GetByIP(context.Context, net.IP) (*data.DHCP, *data.Netboot, error)
}

// BackendWriter is the interface for writing data to a backend.
type BackendWriter interface {
	// Write data (to a backend) based on a mac address for creating a new hardware.
	CreateByMac(context.Context, net.HardwareAddr) error
}

// BackendReadWriter is the interface for getting and setting data from a backend.
type BackendReadWriter interface {
	BackendReader
	BackendWriter
}
