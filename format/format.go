package format

// Format specifies how structured logs should be formatted
type Format interface {
	AddField(name string, value interface{})
	GetFields() string
	IsEmpty() bool
}

// Factory is a function that returns a format specifier
type Factory func() Format

// FactoryOf creates a factory method for a format specifier
func FactoryOf(product Format) Factory {
	return func() Format {
		return product
	}
}
