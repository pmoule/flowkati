package flowkati

import "github.com/pmoule/flowkati/status"

type WorkflowAction[T StepData] func(*Workflowstep[T], T)

type Workflowstep[T StepData] struct {
	index       int
	action      WorkflowAction[T]
	Title       string
	Description string
	Info        string
	Status      status.WorkflowStatus
	InputData   InputData
	Data        T
	StepData    any
}

func NewWorkflowstep[T StepData](title string, description string, action func(*Workflowstep[T], T)) Workflowstep[T] {
	return Workflowstep[T]{Title: title, Description: description, action: action, Status: status.Running}
}

func (ws *Workflowstep[StepData]) Execute() {
	ws.action(ws, ws.Data)
}

func (ws *Workflowstep[StepData]) IsInputRequired() bool {
	if ws.InputData == nil {
		return false
	}

	for _, item := range ws.InputData {
		if item.Prompt != "" {
			return true
		}
	}

	return false
}

func (ws *Workflowstep[StepData]) Finish() {
	for key := range ws.InputData {
		ws.InputData[key] = &InputItem{}
	}

	ws.Status = status.Finished
}
