package util

func ErrString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
