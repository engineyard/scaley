package workflows

// Workflow is an interaface that describes procedures for getting work done. An
// object that implements this interface must provide a private perform()
// method.
type Workflow interface {
	perform() error
}

// Perform, given a workflow, will perform the work defined by said workflow.
func Perform(wf Workflow) error {
	return wf.perform()
}
