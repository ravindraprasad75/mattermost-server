// Copyright (c) 2017-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package api4

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"

	"github.com/mattermost/mattermost-server/utils"
)

func (api *API) InitDataRetention() {
	l4g.Debug(utils.T("api.data_retention.init.debug"))

	api.BaseRoutes.DataRetention.Handle("/policy", api.ApiSessionRequired(getPolicy)).Methods("GET")
}

func getPolicy(c *Context, w http.ResponseWriter, r *http.Request) {

	// No permission check required.

	if policy, err := c.App.GetDataRetentionPolicy(); err != nil {
		c.Err = err
		return
	} else {
		w.Write([]byte(policy.ToJson()))
	}
}
