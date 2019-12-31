package state

import (
	"log"
	"net"
	"sync"
	"time"
)

// State struct that holds information about current state.
type State struct {
	// Should te bell be banged.
	needsBanging bool

	connections map[net.Conn]bool

	// Mux to make sure that state can be edited from multiple routines.
	mux *sync.Mutex
}

// Current (shared) state.
var Current *State

// Initializes Current state.
func CreateState() {
	Current = &State{
		needsBanging: false,
		connections:  make(map[net.Conn]bool),
		mux:          &sync.Mutex{},
	}
}

// Sets needs banging to the specified value (thread safe).
func (st *State) SetNeedsBanging(needsBanging bool) {
	st.mux.Lock()
	st.needsBanging = needsBanging
	st.mux.Unlock()
}

// Returns current needs banging state (thread safe).
func (st *State) NeedsBanging() bool {
	st.mux.Lock()
	defer st.mux.Unlock()

	return st.needsBanging
}

func (st *State) AddConnection(conn net.Conn) {
	st.mux.Lock()
	st.connections[conn] = true
	st.mux.Unlock()
}

func (st *State) RemoveConnection(conn net.Conn) {
	st.mux.Lock()
	delete(st.connections, conn)
	st.mux.Unlock()
}

func (st *State) Bang() {
	st.mux.Lock()

	connections := make([]net.Conn, 0)
	for c := range st.connections {
		connections = append(connections, c)
	}

	for _, c := range connections {
		go func(c net.Conn) {
			err := c.SetWriteDeadline(time.Now().Add(30 * time.Second))
			if err != nil {
				log.Println(err)

				_ = c.Close()
				st.RemoveConnection(c)
				return
			}

			_, err = c.Write([]byte("bang\n"))
			if err != nil {
				log.Println(err)

				_ = c.Close()
				st.RemoveConnection(c)
				return
			}
		}(c)
	}

	st.mux.Unlock()
}
