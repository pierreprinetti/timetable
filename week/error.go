package week

const (
	ErrNotString   weekError = "week database value not of type string"
	ErrEmptyString weekError = "week database value empty"
)

type weekError string

func (wErr weekError) Error() string {
	return string(wErr)
}
