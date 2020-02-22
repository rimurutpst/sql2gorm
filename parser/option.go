package parser

type NullStyle int

const (
	NullInSql NullStyle = iota + 1
	NullInPointer
)

type Option func(*options)

type options struct {
	Charset      string
	Collation    string
	JsonTag      bool
	TablePrefix  string
	ColumnPrefix string
	NoNullType   bool
	NullStyle    NullStyle
	Package      string
}

var defaultOptions = options{
	NullStyle: NullInSql,
	Package:   "model",
}

func WithCharset(charset string) Option {
	return func(o *options) {
		o.Charset = charset
	}
}

func WithCollation(collation string) Option {
	return func(o *options) {
		o.Collation = collation
	}
}

func WithTablePrefix(p string) Option {
	return func(o *options) {
		o.TablePrefix = p
	}
}

func WithColumnPrefix(p string) Option {
	return func(o *options) {
		o.ColumnPrefix = p
	}
}

func WithJsonTag() Option {
	return func(o *options) {
		o.JsonTag = true
	}
}

func WithNoNullType() Option {
	return func(o *options) {
		o.NoNullType = true
	}
}

func WithNullStyle(s NullStyle) Option {
	return func(o *options) {
		o.NullStyle = s
	}
}

func WithPackage(pkg string) Option {
	return func(o *options) {
		o.Package = pkg
	}
}

func parseOption(options []Option) options {
	o := defaultOptions
	for _, f := range options {
		f(&o)
	}
	return o
}
