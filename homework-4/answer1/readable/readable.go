package readable

type Readable interface {
	Open() bool
	Read() bool
	Close()
	State() string
}