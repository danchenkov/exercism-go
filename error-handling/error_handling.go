package erratum

const testVersion = 2

func Use(o ResourceOpener, input string) (err error) {
	var resource Resource

FOR:
	for {
		resource, err = o()
		switch err.(type) {
		case TransientError:
			continue
		case error:
			return err
		default:
			break FOR
		}
	}

	defer func() {
		if r := recover(); r != nil {
			switch r.(type) {
			case FrobError:
				resource.Defrob(r.(FrobError).defrobTag)
				err = r.(FrobError)
			default:
				err = r.(error)
			}
		}
		resource.Close()
	}()

	resource.Frob(input)
	return
}
