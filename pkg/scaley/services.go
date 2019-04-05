package scaley

// Services is a handy collection of the various services one uses to perform
// tasks within Scaley.
type Services struct {
	Groups       GroupService
	Servers      ServerService
	Environments EnvironmentService
	Scripts      ScalingScriptService
	Locker       LockService
	Runner       ExecService
	Log          LogService
}

// GroupService is an interface that describes an object that knows how to
// interact with a Group.
type GroupService interface {
	Get(string) (*Group, error)
}

// ServerService is an interface that describes an object that knows how to
// interact with a Server.
type ServerService interface {
	Get(string) (*Server, error)
	Start(*Server) error
	Stop(*Server) error
}

// EnvironmentService is an interface that describes an object that knows how
// to interact with an Environment.
type EnvironmentService interface {
	Get(string) (*Environment, error)
	Configure(*Environment) error
}

// LockService is an interface that describes an object that knows how to
// deal with Group locking.
type LockService interface {
	Lock(*Group) error
	Unlock(*Group) error
	Locked(*Group) bool
}

// ExecService is an interface that describes an object that knows how to
// execute external commands.
type ExecService interface {
	Run(string) int
}

// ScalingScriptService is an interface that describes an object that knows how
// to interct with a ScalingScript.
type ScalingScriptService interface {
	Exists(string) bool
}

// LogService is an interface that describes an object that provides logging
// capabilities.
type LogService interface {
	Info(*Group, string)
	Failure(*Group, string)
	Success(*Group, string)
}
