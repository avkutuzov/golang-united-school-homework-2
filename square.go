package square

import(
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#
type intCustomType int
const (
	SidesCircle intCustomType = 0
	SidesTriangle intCustomType = 3
	SidesSquare intCustomType = 4
)

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	sn := sidesNum

	switch sn {
	case SidesTriangle:
		p := 3*sideLen/2
		return 3*p*(p-sideLen)
	case SidesSquare:
		return sideLen*sideLen
	case SidesCircle:
		pi := math.Pi
		return pi*sideLen*sideLen
	default:
		return 0
	}	
}
