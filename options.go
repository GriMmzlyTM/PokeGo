package pokego

type Options struct {
	URL             string
	Cache           bool
	CacheDir        string
	CacheExpiryDays uint8
}

var ops Options

//SetOptions sets the options variable to your Options. You must make a new variable of type Options.
func SetOptions(options Options) {
	ops = options
}
