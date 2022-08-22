package builder

type DrumType string

var DrumTypes = struct {
	HAND     DrumType
	MARCHING DrumType
}{
	HAND:     "hand",
	MARCHING: "marching",
}

type DrumMaterial string

var DrumMaterials = struct {
	OAK   DrumMaterial
	MAPLE DrumMaterial
	BIRCH DrumMaterial
}{
	OAK:   "oak",
	MAPLE: "maple",
	BIRCH: "birch",
}
