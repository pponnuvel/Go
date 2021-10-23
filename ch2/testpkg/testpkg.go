package test

const (
	Pi = 3.14
)

type Index int

func (t Index) str() string {
	return fmt.Sprintf("str(%d)", t)
}
