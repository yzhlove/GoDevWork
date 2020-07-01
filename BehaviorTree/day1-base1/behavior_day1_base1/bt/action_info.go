package bt

import (
	"errors"
	"strconv"
	"strings"
)

type CondHpInfo struct {
	Min, Max int
}

func (c *CondHpInfo) Parse(str string) (err error) {
	if params := strings.Split(str, ","); len(params) == 2 {
		if c.Min, err = strconv.Atoi(params[0]); err != nil {
			return err
		}
		if c.Max, err = strconv.Atoi(params[1]); err != nil {
			return err
		}
		return nil
	}
	return errors.New("cond hp params length error")
}

type SkillInfo struct {
	Sid int
	Mp  int
}

func (s *SkillInfo) Parse(str string) (err error) {
	if params := strings.Split(str, ","); len(params) == 2 {
		if s.Sid, err = strconv.Atoi(params[0]); err != nil {
			return err
		}
		if s.Mp, err = strconv.Atoi(params[1]); err != nil {
			return err
		}
		return nil
	}
	return errors.New("skill params length err")
}

type EscapeInfo struct{}

type EatInfo struct {
	Tid int
}

func (e *EatInfo) Parse(str string) (err error) {
	e.Tid, err = strconv.Atoi(str)
	return
}
