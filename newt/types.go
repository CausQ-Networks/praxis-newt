package newt

import (
	wgclients "github.com/CausQ-Networks/praxis-newt/clients"
	"github.com/CausQ-Networks/praxis-newt/exitnode"
	"github.com/CausQ-Networks/praxis-newt/healthcheck"
	"github.com/CausQ-Networks/praxis-newt/netstack2"
)

type BrowserGatewayTarget struct {
	ID              int    `json:"id"`
	Type            string `json:"type"`
	Destination     string `json:"destination"`
	DestinationPort int    `json:"destinationPort"`
	AuthToken       string `json:"authToken"`
}

type WgData struct {
	Endpoint              string                 `json:"endpoint"`
	RelayPort             uint16                 `json:"relayPort"`
	PublicKey             string                 `json:"publicKey"`
	ServerIP              string                 `json:"serverIP"`
	TunnelIP              string                 `json:"tunnelIP"`
	Targets               TargetsByType          `json:"targets"`
	HealthCheckTargets    []healthcheck.Config   `json:"healthCheckTargets"`
	BrowserGatewayTargets []BrowserGatewayTarget `json:"browserGatewayTargets"`
	RemoteExitNodeSubnets []string               `json:"remoteExitNodeSubnets"`
	ChainId               string                 `json:"chainId"`
}

type TargetsByType struct {
	UDP []string `json:"udp"`
	TCP []string `json:"tcp"`
}

type TargetData struct {
	Targets []string `json:"targets"`
}

// ExitNodeData, ExitNode and ExitNodePingResult are aliases for the shared
// exit-node ping dance types in package exitnode, kept here so existing code
// in this package can keep referring to them unqualified.
type (
	ExitNodeData       = exitnode.ExitNodeData
	ExitNode           = exitnode.ExitNode
	ExitNodePingResult = exitnode.ExitNodePingResult
)

type BlueprintResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// Define the sync data structure
type SyncData struct {
	Targets               TargetsByType              `json:"proxyTargets"`
	HealthCheckTargets    []healthcheck.Config       `json:"healthCheckTargets"`
	RemoteExitNodeSubnets []string                   `json:"remoteExitNodeSubnets"`
	Peers                 []wgclients.Peer           `json:"peers"`
	ClientTargets         []wgclients.Target         `json:"clientTargets"`
	Certs                 []wgclients.CertData       `json:"certs"`
	ClientIdentities      []netstack2.ClientIdentity `json:"clientIdentities"`
	BrowserGatewayTargets []BrowserGatewayTarget     `json:"browserGatewayTargets"`
}
