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
	Star        int
	End         int
	Description string
	Name        string
	Trigger     Trigger
}
type ParamenterFont int

const (
	HeaderParam ParamenterFont = iota
	QueryParam
	PathParam
	BodyParam
)

type ParamenterType int

const (
	StringType ParamenterType = iota
	NumberType
	BooleanType
	DateTimeType
	StringArrayType
)

type Parameter struct {
	Name     string
	Font     ParamenterFont
	Required bool
	Type     ParamenterType
}
type Route struct {
	Name            string
	AcceptMethods   []string
	Priority        int
	Category        string
	Help            string
	LongDescription string
	Parameters      []Parameter
	Paths           []Path
	Examples        []string
}
