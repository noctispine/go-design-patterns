package factory

import "errors"

func DrumFactory(drumType DrumType, age int) (IDrum, error) {
	switch drumType {
	case DrumTypes.HAND:
		{
			return newHandDrum(age), nil
		}

	case DrumTypes.MARCHING:
		{
			return newMarchingDrum(age), nil
		}

	}

	return nil, errors.New("wrong type of drum")
}
