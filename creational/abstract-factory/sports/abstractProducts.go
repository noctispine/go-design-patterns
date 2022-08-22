package sports

type IShoe interface {
	setLogo(string)
	setSize(int)
	getLogo() string
	getSize() int
}

type Shoe struct {
	logo string
	size int
}

func (s *Shoe) setLogo(logo string) {
	s.logo = logo
}

func (s *Shoe) setSize(size int) {
	s.size = size
}

func (s *Shoe) getLogo() string {
	return s.logo
}

func (s *Shoe) getSize() int {
	return s.size
}

type IShirt interface {
	setLogo(string)
	setSize(int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	logo string
	size int
}

func (s *Shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *Shirt) setSize(size int) {
	s.size = size
}

func (s *Shirt) getLogo() string {
	return s.logo
}

func (s *Shirt) getSize() int {
	return s.size
}
