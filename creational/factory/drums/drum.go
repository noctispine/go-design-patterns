package factory

type Drum struct {
	name DrumType
	age  int
}

func (d *Drum) setName(name DrumType) {
	d.name = name
}

func (d *Drum) GetName() DrumType {
	return d.name
}

func (d *Drum) setAge(age int) {
	d.age = age
}

func (d *Drum) GetAge() int {
	return d.age
}
