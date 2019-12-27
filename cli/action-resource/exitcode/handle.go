package exitcode

type Handle struct {
	ExitCode int
}

func NewHandle() *Handle {
	handle := &Handle{}
	handle.ExitCode = 1
	return handle
}

func (handle *Handle) ToErr() {
	handle.ExitCode = 1
}

func (handle *Handle) ToSuccess() {
	handle.ExitCode = 0
}

func (handle *Handle) DefaultErr() *Handle {
	handle.ExitCode = 1
	return handle
}

func (handle *Handle) DefaultSuccess() *Handle {
	handle.ExitCode = 0
	return handle
}
