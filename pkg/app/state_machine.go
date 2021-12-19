package app

const (
	stateInit int = iota
	stateStarted
	stateRunning
	stateStopped
	stateFatal
)

type StateResponse struct {
	Err  error
	Log  string
	Data interface{}
}

var stateName = [...]string{
	stateInit:    "Init",
	stateStarted: "Started",
	stateRunning: "Running",
	stateStopped: "Stopped",
	stateFatal:   "Fatal",
}

type stateMachine struct {
	state   string
	actionc func(chan StateResponse)
	logs    []string
	quitc   chan StateResponse
}

func NewStateMachine(f func(sm chan StateResponse)) (*stateMachine, chan StateResponse) {
	sm := &stateMachine{
		state:   stateName[0],
		actionc: f,
		logs:    []string{"State Machine Initiated"},
		quitc:   make(chan StateResponse, 1),
	}
	return sm, sm.quitc
}

func (sm *stateMachine) GetState() string {
	return sm.state
}

func (sm *stateMachine) GetLogs() []string {
	return sm.logs
}

func (sm *stateMachine) Run() {

	runner := make(chan StateResponse, 1)

	sm.state = stateName[1]
	go sm.actionc(runner)
	err := <-runner
	sm.quitc <- err
	sm.state = stateName[4]
}
