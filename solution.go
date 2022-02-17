package square

import "math"

type myInt int

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
)

func CalcSquare(sideLen float64, sidesNum myInt) float64 {
	switch sidesNum {
	case SidesTriangle:
		return (sideLen * math.Sqrt(3)) / 4
	case SidesSquare:
		return math.Pow(sideLen, 2)
	case SidesCircle:
		return (math.Pi * math.Pow(sideLen, 2)) / 2
	default:
		return 0
	}
}
