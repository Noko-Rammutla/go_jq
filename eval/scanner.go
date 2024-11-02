package eval

func Scan(source string) ([]Token, error) {
	s := scanner{
		source:  source,
		current: 0,
	}
	return s.scan()
}

type scanner struct {
	source  string
	current int
}

func (s *scanner) scan() ([]Token, error) {
	return nil, nil
}
