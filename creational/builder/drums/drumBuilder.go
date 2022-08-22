package builder

type DrumBuilder interface {
	Name(string) DrumBuilder
	DrumType(DrumType) DrumBuilder
	Build() IDrum
}

type drumBuilder struct {
	name     string
	material DrumMaterial
	drumType DrumType
}

func (d *drumBuilder) Name(name string) DrumBuilder {
	d.name = name
	return d
}

func (d *drumBuilder) DrumType(drumType DrumType) DrumBuilder {
	d.drumType = drumType
	return d
}

func (d *drumBuilder) Build() IDrum {
	var drum IDrum
	switch d.drumType {
	case DrumTypes.HAND:
		{
			return newHandDrum(string(DrumTypes.HAND))
		}
	case DrumTypes.MARCHING:
		{
			return newMarchingDrum(string(DrumTypes.MARCHING))
		}
	}

	return drum
}

func NewDrumBuilder() DrumBuilder {
	return &drumBuilder{}
}
