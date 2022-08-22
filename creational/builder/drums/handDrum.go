package builder

type handDrum struct {
	Drum
}

func newHandDrum(name string) IDrum {
	return &handDrum{
		Drum: Drum{
			name:     name,
			drumType: DrumTypes.HAND,
			material: DrumMaterials.OAK,
		},
	}
}
