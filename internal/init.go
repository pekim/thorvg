package internal

func Init() error {
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
