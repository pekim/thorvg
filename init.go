package thorvg

var initialised = false

func initLibraries() error {
	if initialised {
		return nil
	}
	defer func() { initialised = true }()

	if err := initLibc(); err != nil {
		return err
	}
	if err := initLibThorvg(); err != nil {
		return err
	}

	return nil
}
