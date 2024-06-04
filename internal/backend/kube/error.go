package kube

type hardwareNotFoundError struct{}

func (hardwareNotFoundError) NotFound() bool { return true }

func (hardwareNotFoundError) Error() string { return "hardware not found" }

func IsHardwareNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	e, ok := err.(hardwareNotFoundError)
	return ok && e.NotFound()
}
