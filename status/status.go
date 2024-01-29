package status

const (
	Running  WorkflowStatus = "Running"
	Finished WorkflowStatus = "Finished"
)

type WorkflowStatus string

func (ws WorkflowStatus) String() string {
	return string(ws)
}
