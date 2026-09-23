package api

type TriggerType int

const (
	EqualTrigger TriggerType = iota
	PrefixTrigger
	SuffixTrigger
	RegexTrigger
)

type Trigger struct {
	Exist bool
	Type  TriggerType
	Value string
}
type Path struct {
	Star    int
	End     int
	Name    string
	Trigger Trigger
}

type Route struct {
	Name            string
	AcceptMethods   []string
	Priority        int
	Category        string
	Help            string
	LongDescription string

	Paths    []Path
	Examples []string
}
