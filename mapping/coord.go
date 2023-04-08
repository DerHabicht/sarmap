package mapping

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type Direction int

const (
	InvalidDirection = -1 + iota
	North
	South
	East
	West
)

func ParseDirection(dir string) (Direction, error) {
	switch strings.ToLower(dir) {
	case "n":
	case "north":
		return North, nil
	case "s":
	case "south":
		return South, nil
	case "e":
	case "east":
		return East, nil
	case "w":
	case "west":
		return West, nil
	}

	return InvalidDirection, fmt.Errorf("invalid coordinate direction: %s", dir)
}

func (d Direction) String() string {
	switch d {
	case North:
		return "N"
	case South:
		return "S"
	case East:
		return "E"
	case West:
		return "W"
	default:
		return ""
	}
}

type Coord struct {
	Degrees   int
	Minutes   int
	Seconds   float64
	Direction Direction
}

func ParseCoord(coord string) (Coord, error) {
	var deg int
	var min int
	var sec float64
	var dir string

	_, err := fmt.Scanf("%d-%d-%f%s", &deg, &min, &sec, &dir)
	if err != nil {
		return Coord{}, errors.WithStack(err)
	}
}
