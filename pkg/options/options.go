package opt

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	JSONOmitEmpty bool
	Panic         bool
}

// JSONOmitEmptyEnabled ...
func JSONOmitEmptyEnabled(f *Options) {
	f.JSONOmitEmpty = true
	return
}

// PanicEnabled allow Check to panic if an empty value is detected
func PanicEnabled(f *Options) {
	f.Panic = true
	return
}
