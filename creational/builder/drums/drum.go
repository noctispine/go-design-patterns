package builder

type IDrum interface {
	GetName() string
	GetType() DrumType
	GetMaterial() DrumMaterial
}

type Drum struct {
	name     string
	drumType DrumType
	material DrumMaterial
}

func (d *Drum) GetName() string {
	return d.name
}

func (d *Drum) GetType() DrumType {
	return d.drumType
}

func (d *Drum) GetMaterial() DrumMaterial {
	return d.material
}
