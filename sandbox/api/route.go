package api

type Route struct {
	Name            string
	AcceptMethods   []string
	Priority        int
	Category        string
	Help            string
	LongDescription string
	Examples        []string
}
