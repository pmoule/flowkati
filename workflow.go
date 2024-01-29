package flowkati

import (
	"fmt"
	"sort"

	"github.com/pmoule/flowkati/status"
)

type Dto interface {
	ToDto() any
}

type StepData interface {
	fmt.Stringer
	Dto
}

type Workflow[T StepData] struct {
	id         string
	steps      []*Workflowstep[T]
	activeStep *Workflowstep[T]
	data       T
}

func NewWorkflow[T StepData](id string, data T) *Workflow[T] {
	return &Workflow[T]{id: id, steps: []*Workflowstep[T]{}, data: data}
}

func (w *Workflow[T]) ID() string {
	return w.id
}

func (w *Workflow[T]) AddStep(step Workflowstep[T]) {
	step.index = len(w.steps) + 1
	w.steps = append(w.steps, &step)
}

func (w *Workflow[T]) ActiveStep() *Workflowstep[T] {
	return w.activeStep
}

func (w *Workflow[T]) Steps() []Workflowstep[T] {
	steps := []Workflowstep[T]{}

	for _, step := range w.steps {
		steps = append(steps, *step)
	}

	return steps
}

func (w *Workflow[T]) Status() status.WorkflowStatus {
	workflowStatus := status.Finished

	for _, step := range w.steps {
		if step.Status == status.Finished {
			continue
		}

		workflowStatus = step.Status

		break
	}

	return workflowStatus
}

func (w *Workflow[T]) Run() bool {
	if len(w.steps) == 0 {
		return false
	}

	sort.Slice(w.steps, func(i int, j int) bool {
		return w.steps[i].index < w.steps[j].index
	})

	var activeStep *Workflowstep[T]

	if w.ActiveStep() == nil {
		activeStep = w.steps[0]
	} else if w.ActiveStep().index < len(w.steps) {
		nextIndex := w.ActiveStep().index + 1
		activeStep = w.steps[nextIndex-1]
	}

	if activeStep == nil {
		return false
	}

	w.activeStep = activeStep
	w.activeStep.Data = w.data

	return true
}
