package data

import "fmt"

type DataError struct {
	Err    string
	Code   int
	Action string
}

func (de DataError) Error() string {
	errPlc := 0

	if de.Err != "" {
		errPlc++
	}
	if de.Code != 0 {
		errPlc += 2
	}
	if de.Action != "" {
		errPlc += 3
	}

	switch errPlc {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%s", de.Err)
	case 2:
		return fmt.Sprintf("%d", de.Code)
	case 3:
		if de.Action == "" {
			return fmt.Sprintf("%d:%s", de.Code, de.Err)
		}
		return fmt.Sprintf("%s", de.Action)
	case 4:
		return fmt.Sprintf("%s: CORRECTIVE ACTION %s", de.Err, de.Action)
	case 5: //aka case 5
		return fmt.Sprintf("%d: CORRECTIVE ACTION %s", de.Code, de.Action)
	case 6: //aka case 5
		return fmt.Sprintf("%d:%s: CORRECTIVE ACTION %s", de.Code, de.Err, de.Action)
	}
	return ""
}
