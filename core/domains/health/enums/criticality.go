package enums

type Criticality int

const (
	SuperCritical Criticality = iota + 1
	HighCritical
	MediumCritical
	LowCritical
	NotCritical
)

func (c Criticality) Int() int {
	return int(c)
}
