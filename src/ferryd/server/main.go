//
// Copyright © 2017 Ikey Doherty <ikey@solus-project.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package server

import (
	"net"
	"net/http"
	"os"
)

const (
	// UnixSocketPath is the unique socket path on the system for the ferry daemon
	UnixSocketPath = "./ferryd.sock"
)

// Server sits on a unix socket accepting connections from authenticated
// client, i.e. root or those in the "ferry" group
type Server struct {
	socket net.Listener
}

// New will return a newly initialised Server which is currently unbound
func New() *Server {
	return &Server{}
}

// Serve will continuously serve on the unix socket until dead
func (s *Server) Serve() error {
	l, e := net.Listen("unix", UnixSocketPath)
	if e != nil {
		return e
	}
	s.socket = l
	return http.Serve(l, nil)
}

// Close will shut down and cleanup the socket
func (s *Server) Close() {
	s.socket.Close()
	os.Remove(UnixSocketPath)
}
