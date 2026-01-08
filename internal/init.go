package internal

var initialised = false

func Init() error {
	if initialised {
		return nil
	}
	defer func() { initialised = true }()

	var err error

	err = initLibc()
	if err != nil {
		return err
	}

	err = initLibThorvg()
	if err != nil {
		return err
	}

	return nil
}
