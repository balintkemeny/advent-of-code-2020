package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
)

/*
HEADING: N = 0, E = 90, S = 180, W = 270
LATITUDE: N > 0, S < 0
LONGITUDE: E > 0, W < 0
*/
type ship struct {
	heading   int
	latitude  int
	longitude int
	wp        waypoint
}

type waypoint struct {
	latitude  int
	longitude int
}

func executeShipCommand(s *ship, command string) error {
	if len(command) > 4 || len(command) < 2 {
		return errors.New("COMMAND IS NOT OF APPROPRIATE LENGTH")
	}

	var commandType rune = rune(command[0])
	commandSize, err := strconv.Atoi(command[1:])
	if err != nil {
		return err
	}

	switch commandType {
	case 'N':
		s.latitude += commandSize
	case 'E':
		s.longitude += commandSize
	case 'S':
		s.latitude -= commandSize
	case 'W':
		s.longitude -= commandSize
	case 'R':
		s.heading += commandSize
		if s.heading >= 360 {
			s.heading -= 360
		}
	case 'L':
		s.heading -= commandSize
		if s.heading < 0 {
			s.heading += 360
		}
	case 'F':
		switch s.heading {
		case 0:
			s.latitude += commandSize
		case 90:
			s.longitude += commandSize
		case 180:
			s.latitude -= commandSize
		case 270:
			s.longitude -= commandSize
		default:
			return errors.New("WE'RE HEADING IN A STRANGE DIRECTION")
		}
	default:
		return errors.New("UNRECOGNIZED COMMAND TYPE")
	}
	return nil
}

func executeWaypointCommand(s *ship, command string) error {
	if len(command) > 4 || len(command) < 2 {
		return errors.New("COMMAND IS NOT OF APPROPRIATE LENGTH")
	}

	var commandType rune = rune(command[0])
	commandSize, err := strconv.Atoi(command[1:])
	if err != nil {
		return err
	}

	switch commandType {
	case 'N':
		s.wp.latitude += commandSize
	case 'E':
		s.wp.longitude += commandSize
	case 'S':
		s.wp.latitude -= commandSize
	case 'W':
		s.wp.longitude -= commandSize
	case 'R':
		switch commandSize {
		case 90:
			s.wp.latitude, s.wp.longitude = -s.wp.longitude, s.wp.latitude
		case 180:
			s.wp.latitude, s.wp.longitude = -s.wp.latitude, -s.wp.longitude
		case 270:
			s.wp.latitude, s.wp.longitude = s.wp.longitude, -s.wp.latitude
		default:
			return errors.New("INAPPROPRIATE WAYPOINT SWITCHING DEGREE")
		}
	case 'L':
		switch commandSize {
		case 90:
			s.wp.latitude, s.wp.longitude = s.wp.longitude, -s.wp.latitude
		case 180:
			s.wp.latitude, s.wp.longitude = -s.wp.latitude, -s.wp.longitude
		case 270:
			s.wp.latitude, s.wp.longitude = -s.wp.longitude, s.wp.latitude
		default:
			return errors.New("INAPPROPRIATE WAYPOINT SWITCHING DEGREE")
		}
	case 'F':
		s.latitude += commandSize * s.wp.latitude
		s.longitude -= commandSize * s.wp.longitude
	default:
		return errors.New("UNRECOGNIZED COMMAND TYPE")
	}
	return nil
}

func manhattan(lat, lon int) int {
	return int(math.Abs(float64(lat)) + math.Abs(float64(lon)))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rawData, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var commands []string
	for _, line := range bytes.Split(rawData, []byte{'\n'}) {
		command := string(line)
		commands = append(commands, command)

	}

	s := ship{90, 0, 0, waypoint{0, 0}}
	for _, command := range commands {
		fmt.Printf("COMMAND: %s, CURRENT HEADING: %d, LAT: %d, LON: %d\n", command, s.heading, s.latitude, s.longitude)
		err := executeShipCommand(&s, command)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("EX01:", s)
	fmt.Println("EX01:", manhattan(s.latitude, s.longitude))

	s2 := ship{90, 0, 0, waypoint{1, 10}}
	for _, command := range commands {
		fmt.Printf("COMMAND: %s, CURRENT HEADING: %d, LAT: %d, LON: %d\n", command, s2.heading, s2.latitude, s2.longitude)
		err := executeWaypointCommand(&s2, command)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("EX01:", s2)
	fmt.Println("EX01:", manhattan(s2.latitude, s2.longitude))
}
