package retry

type ContinueRetry struct {
	Err error
}

func (e *ContinueRetry) Error() string {
	if e.Err == nil {
		return "unknown error"
	}
	return e.Err.Error()
}

type ExitRetry struct {
	Err error
}

func (e *ExitRetry) Error() string {
	if e.Err == nil {
		return "unknown error"
	}
	return e.Err.Error()
}
