package tests

import (
	"green-api-test/client"
	"testing"

	"github.com/stretchr/testify/require"
)

type StateInstance string

var IsAuthorized bool

const (
	NotAuthorized StateInstance = "notAuthorized"
	Authorized    StateInstance = "authorized"
	Blocked       StateInstance = "blocked"
	SleepMode     StateInstance = "sleepMode"
	Starting      StateInstance = "starting"
	YellowCard    StateInstance = "yellowCard"
)

func requireAuthorized(t *testing.T, c *client.Client) {
	t.Helper()

	if IsAuthorized {
		return
	}

	resp, err := c.GetStateInstance()
	require.NoError(t, err)
	require.NotNil(t, resp)

	require.Equal(t, 200, resp.StatusCode)
	require.Equal(t, Authorized, StateInstance(resp.StateInstance), "Instance is not authorized, state: %s", resp.StateInstance)

	IsAuthorized = true
}
