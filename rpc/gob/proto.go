package gob

type Request struct {
	Name  string
	Age   int
	Phone []string
	Data  map[string]string
}

type Reply struct {
	Name  string
	Age   int
	Phone []string
	Data  map[string]string
}
