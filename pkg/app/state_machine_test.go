package app_test

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/ahmed-saleh/playbook/pkg/app"
)

func TestStateMachine(t *testing.T) {

	//action function
	actionFunc := func(v chan app.StateResponse) {
		time.Sleep(time.Second * 2)
		response := app.StateResponse{
			Err:  errors.New("new error"),
			Log:  "machine was terminated",
			Data: "stack trace",
		}
		v <- response
	}

	sm, ch := app.NewStateMachine(actionFunc)
	assert.Equal(t, "Init", sm.GetState(), "they should be equal")

	assert.Equal(t, 1, len(sm.GetLogs()), "they should be equal")
	go sm.Run()
	time.Sleep(time.Millisecond * 100)
	assert.Equal(t, "Started", sm.GetState(), "they should be equal")

	res := <-ch
	assert.NotNil(t, res.Err)
	assert.Equal(t, "machine was terminated", res.Log, "they should be equal")
	assert.Equal(t, "Fatal", sm.GetState(), "they should be equal")

}
