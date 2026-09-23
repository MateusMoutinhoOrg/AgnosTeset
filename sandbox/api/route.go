package api

type Path struct {
	Star int
	End  int
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
