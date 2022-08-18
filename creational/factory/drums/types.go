package drums

type DrumType string

var DrumTypes = struct {
	HAND     DrumType
	MARCHING DrumType
}{
	HAND:     "hand",
	MARCHING: "marching",
}
