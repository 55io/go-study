package study_errors

import errors
// Similar to:
//   if e, ok := err.(*QueryError); ok { … }
var e *QueryError
if errors.As(err, &e) {
    // err is a *QueryError, and e is set to the error's value
}

    // The *os.PathError returned by os.Open is an internal detail.
    // To avoid exposing it to the caller, repackage it as a new
    // error with the same text. We use the %v formatting verb, since
    // %w would permit the caller to unwrap the original *os.PathError.
    return fmt.Errorf("%v", err)
	