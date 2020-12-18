package opt

// Option ...
type Option func(f *Options)

const (
	flag = "emptyvalue"
)

// Tag ...
type Tag []string

// Options ...
type Options struct {
	Panic bool
	Tags  Tag
}

// OmitEmptyTag ...
func OmitEmptyTag(t string, v ...string) Option {
	return func(f *Options) {
		f.Tags = append(v, t)
	}
}

// PanicEnabled allow Check to panic if an empty value is detected
func PanicEnabled(f *Options) {
	f.Panic = true
	return
}

// ApplyTagEmptyValue apply emptyvalue as generic flag
func (f *Options) ApplyTagEmptyValue() {
	f.Tags = append(f.Tags, flag)
}
