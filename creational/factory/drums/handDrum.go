package drums

type handDrum struct {
	Drum
}

func newHandDrum(age int) IDrum {
	return &handDrum{
		Drum: Drum{
			name: DrumTypes.HAND,
			age:  age,
		},
	}
}
