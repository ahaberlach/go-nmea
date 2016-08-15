package nmea

import (
	"fmt"
	"log"
	"strconv"
	)

const (
	PrefixPUBX = "PUBX"
)

type PUBX0 struct {
	Sentence

	//Subclass of this PUBX message	(always 0 for PUBX0)
	ID          int64
	Time string
	Latitude    LatLong
	Longitude   LatLong
	Altitude    float64  // m
	NavStat	    string
	HorizAcc    float64  // m
	VertAcc     float64  // m
	Speed	    float64  // km/h
	Course      float64  // degrees (bearing)
}

func NewPUBX0(sentence Sentence) PUBX0 {
	s := new(PUBX0)
	s.Sentence = sentence
	return *s
}

func (s *PUBX0) parse() error {
	var err error
	if s.Type != PrefixPUBX {
		return fmt.Errorf("%s is not a %s", s.Type, PrefixPUBX)
	}
	log.Printf("%v", s.Fields)
	s.ID, err = strconv.ParseInt(s.Fields[0], 10, 8)
	if err != nil {
		return err
	}
	s.Time = s.Fields[1]
	s.Latitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[2], s.Fields[3]))
	if err != nil {
		return fmt.Errorf("PUBX decode error: %s", err)
	}
	s.Longitude, err = NewLatLong(fmt.Sprintf("%s %s", s.Fields[4], s.Fields[5]))
	if err != nil {
		return fmt.Errorf("PUBX decode error: %s", err)
	}
	s.Altitude, err = strconv.ParseFloat(s.Fields[6], 64)
	if err != nil {
		return err
	}
	s.NavStat = s.Fields[7]
	s.HorizAcc, err = strconv.ParseFloat(s.Fields[8], 64)
	if err != nil {
		return err
	}
	s.VertAcc, err = strconv.ParseFloat(s.Fields[9], 64)
	if err != nil {
		return err
	}
	s.Speed, err = strconv.ParseFloat(s.Fields[10], 64)
	if err != nil {
		return err
	}
	s.Course, err = strconv.ParseFloat(s.Fields[11], 64)
	if err != nil {
		return err
	}
	return nil
}
