package helper

func PanicIfError(err error) {
	if err != nil {
		print("ERROR: ", string(err.Error()))
		panic(err)
	}
}
