package node

import (
	"fmt"
	"log"
	"net"

	"crypto/tls"

	"github.com/numbleroot/pluto/config"
	"github.com/numbleroot/pluto/imap"
)

// Constants

// Integer counter for defining node types.
const (
	DISTRIBUTOR Type = iota
	WORKER
	STORAGE
)

// Structs

// Type declares what role a node takes in the system.
type Type int

// Node struct bundles information of one node instance.
type Node struct {
	Type   Type
	Config *config.Config
	Socket net.Listener
}

// Functions

// InitNode listens for TLS connections on a TCP socket
// opened up on supplied IP address and port. It returns
// those information bundeled in above Node struct.
func InitNode(config *config.Config, distributor bool, worker string, storage bool) (*Node, error) {

	var err error
	node := new(Node)

	// Check if no type indicator was supplied, not possible.
	if !distributor && worker == "" && !storage {
		return nil, fmt.Errorf("[node.InitNode] Node must be of one type, either '-distributor' or '-worker WORKER-ID' or '-storage'.\n")
	}

	// Check if multiple type indicators were supplied, not possible.
	if (distributor && worker != "" && storage) || (distributor && worker != "") || (distributor && storage) || (worker != "" && storage) {
		return nil, fmt.Errorf("[node.InitNode] One node can not be of multiple types, please provide exclusively '-distributor' or '-worker WORKER-ID' or '-storage'.\n")
	}

	if distributor {

		// Set struct type to distributor.
		node.Type = DISTRIBUTOR

		// TLS config is taken from the excellent blog post
		// "Achieving a Perfect SSL Labs Score with Go":
		// https://blog.bracelab.com/achieving-perfect-ssl-labs-score-with-go
		tlsConfig := &tls.Config{
			Certificates:             make([]tls.Certificate, 1),
			MinVersion:               tls.VersionTLS12,
			CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
			PreferServerCipherSuites: true,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			},
		}

		// Put in supplied TLS cert and key.
		tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(config.Distributor.TLS.CertLoc, config.Distributor.TLS.KeyLoc)
		if err != nil {
			return nil, fmt.Errorf("[node.InitNode] Failed to load distributor TLS cert and key: %s\n", err.Error())
		}

		// Build Common Name (CN) and Subject Alternate
		// Name (SAN) from tlsConfig.Certificates.
		tlsConfig.BuildNameToCertificate()

		// Start to listen on defined IP and port.
		node.Socket, err = tls.Listen("tcp", fmt.Sprintf("%s:%s", config.Distributor.IP, config.Distributor.Port), tlsConfig)
		if err != nil {
			return nil, fmt.Errorf("[node.InitNode] Listening for TLS connections on distributor port failed with: %s\n", err.Error())
		}

		log.Printf("[node.InitNode] Listening as distributor node for incoming IMAP requests on %s.\n", node.Socket.Addr())

	} else if worker != "" {

		// Check if supplied worker ID actually is configured.
		if _, ok := config.Workers[worker]; !ok {

			var workerID string

			// Retrieve first valid worker ID to provide feedback.
			for workerID = range config.Workers {
				break
			}

			return nil, fmt.Errorf("[node.InitNode] Specified worker ID does not exist in config file. Please provide a valid one, for example '%s'.\n", workerID)
		}

		// Set struct type to worker.
		node.Type = WORKER

		// Start to listen on defined IP and port.
		node.Socket, err = net.Listen("tcp", fmt.Sprintf("%s:%s", config.Workers[worker].IP, config.Workers[worker].Port))
		if err != nil {
			return nil, fmt.Errorf("[node.InitNode] Listening for TCP connections on worker '%s' failed with: %s\n", worker, err.Error())
		}

		log.Printf("[node.InitNode] Listening as worker node '%s' for incoming IMAP requests on %s.\n", worker, node.Socket.Addr())

	} else if storage {

		// Set struct type to storage.
		node.Type = STORAGE

		// Start to listen on defined IP and port.
		node.Socket, err = net.Listen("tcp", fmt.Sprintf("%s:%s", config.Storage.IP, config.Storage.Port))
		if err != nil {
			return nil, fmt.Errorf("[node.InitNode] Listening for TCP connections on storage failed with: %s\n", err.Error())
		}

		log.Printf("[node.InitNode] Listening as storage node for incoming IMAP requests on %s.\n", node.Socket.Addr())

	}

	// Set remaining general elements.
	node.Config = config

	return node, nil
}

// HandleRequest acts as the jump start for any new
// incoming connection to pluto. It creates the needed
// control structure, sends out the initial server
// greeting and after that hands over control to the
// IMAP state machine.
func (node *Node) HandleRequest(conn net.Conn, greeting string) {

	// Create a new connection struct for incoming request.
	c := imap.NewConnection(conn)

	// Send initial server greeting.
	err := c.Send("* OK IMAP4rev1 " + greeting)
	if err != nil {
		c.Error("Encountered send error", err)
		return
	}

	// Dispatch to not-authenticated state.
	c.Transition(imap.NOT_AUTHENTICATED)
}

// RunNode loops over incoming requests and
// dispatches each one to a goroutine taking
// care of the commands supplied.
func (node *Node) RunNode(greeting string) error {

	for {

		// Accept request or fail on error.
		conn, err := node.Socket.Accept()
		if err != nil {
			return fmt.Errorf("[node.RunNode] Accepting incoming request failed with: %s\n", err.Error())
		}

		// Dispatch to goroutine.
		go node.HandleRequest(conn, greeting)
	}
}
