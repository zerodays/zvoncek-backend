package state

import "sync"

// State struct that holds information about current state.
type State struct {
	// Should te bell be banged.
	needsBanging bool

	// Mux to make sure that state can be edited from multiple routines.
	mux *sync.Mutex
}

// Current (shared) state.
var Current *State

// Initializes Current state.
func CreateState() {
	Current = &State{
		needsBanging: false,
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
