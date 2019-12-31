package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"zvon/state"
)

// Client should call this route to know, weather it should bang.
// Route accepts optional GET parameter 'will_bang', which specifies if server should
// change the 'needs_banging' state. If 'will_bang' is true, then 'needs_banging' will be set
// to false (after notifying the client of current needs_banging state). If 'will_bang' is false,
// then server will only display the current 'needs_banging' state and not change anything.
func BangerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Get will_bang parameter and convert it to bool.
	willBangRaw := r.URL.Query().Get("will_bang")
	willBang, err := strconv.ParseBool(willBangRaw)
	if err != nil {
		willBang = false
	}

	// Get current 'needs banging' state and change it to false if will bang is true.
	needsBanging := state.Current.NeedsBanging()
	if willBang {
		state.Current.SetNeedsBanging(false)
	}

	// Write response to client.
	res, _ := json.Marshal(map[string]interface{}{
		"needs_banging": needsBanging,
	})
	_, _ = w.Write(res)
}

func Bang(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	state.Current.SetNeedsBanging(true)
	state.Current.Bang()

	w.WriteHeader(http.StatusOK)
}
