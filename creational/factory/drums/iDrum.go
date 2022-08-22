package factory

type IDrum interface {
	setName(name DrumType)
	setAge(age int)
	GetName() DrumType
	GetAge() int
}
