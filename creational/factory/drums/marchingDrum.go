package drums

type marchingDrum struct {
	Drum
}

func newMarchingDrum(age int) IDrum {
	return &marchingDrum{
		Drum: Drum{
			name: DrumTypes.MARCHING,
			age:  age,
		},
	}
}
