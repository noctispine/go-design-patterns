package builder

type marchingDrum struct {
	Drum
}

func newMarchingDrum(name string) IDrum {
	return &marchingDrum{
		Drum: Drum{
			name:     name,
			drumType: DrumTypes.MARCHING,
			material: DrumMaterials.MAPLE,
		},
	}
}
