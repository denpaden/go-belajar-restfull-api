package helper

func PanicIfError(err error) {
	if err != nil {
		LoggerError(err.Error())
		panic(err)
	}
}
