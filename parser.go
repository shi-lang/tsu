package main

import (
	"strconv"
	"strings"
)

const EOF = rune(0)

type ParserState struct {
	input string
	pos   int
}

func (s *ParserState) Peek() rune {
	if s.pos == len(s.input) {
		return EOF
	}
	return []rune(s.input)[s.pos]
}

func (s *ParserState) Read() rune {
	c := s.Peek()
	if c != EOF {
		s.pos++
	}
	return c
}

func (ParserState) ValidNumChar(c rune) bool {
	return c >= '0' && c <= '9'
}

func (s *ParserState) ParseNum() *Obj {
	var num int64 = int64(s.Peek() - '0')
	for s.ValidNumChar(s.Peek()) {
		num = (num * 10) + int64(s.Peek()-'0')
		s.Read()
	}
	return NewInt(num)
}

func (s *ParserState) ValidSymFirstChar(c rune) bool {
	return strconv.IsPrint(c) && !strings.ContainsRune(" :[](){}#`'\"\\", c) && !s.ValidNumChar(c)
}

func (ParserState) ValidSymChar(c rune) bool {
	return strconv.IsPrint(c) && !strings.ContainsRune(" :[](){}#`'\"\\", c)
}

func (s *ParserState) ParseSym() *Obj {
	var sym string = ""
	for s.ValidSymChar(s.Peek()) {
		sym += string(s.Peek())
		s.Read()
	}

	// Parse function application
	if s.Peek() == '(' {
		callArgs := s.ParseArgs()
		return NewCall(InternSym(sym), callArgs)
	}

	return InternSym(sym)
}

func (ParserState) ValidKeyChar(c rune) bool {
	return strconv.IsPrint(c) && !strings.ContainsRune(" :[](){}#`'\"\\", c)
}

func (s *ParserState) ParseKey() *Obj {
	var key string = ""
	for s.ValidSymChar(s.Peek()) {
		key += string(s.Peek())
		s.Read()
	}
	return InternSym(key)
}

func (s *ParserState) ParseArgs() *Obj {
	items := []*Obj{}
	s.Read() // Skip opening (
	o := s.Parse()
	for o != RParen {
		if o == nil {
			panic("parser: unterminated parenthesis '('")
		}
		items = append(items, o)
		o = s.Parse()
	}
	return NewVec(items)
}

func (s *ParserState) Parse() *Obj {
	c := s.Peek()

	if c == EOF {
		return nil
	} else if strings.ContainsRune(" \n\r\t", c) {
		s.Read()
		return s.Parse()
	} else if c == '#' {
		for ; !strings.ContainsRune("\r\n", c); c = s.Read() {
			if c != EOF {
				return nil
			}
		}
		return s.Parse()
	} else if c == ')' {
		s.Read()
		return RParen
	} else if c == ']' {
		s.Read()
		return RSqaBr
	} else if c == '}' {
		s.Read()
		return RCurBr
	} else if c == ':' {
		return s.ParseKey()
	} else if s.ValidNumChar(c) {
		return s.ParseNum()
	} else if s.ValidSymFirstChar(c) {
		return s.ParseSym()
	} else {
		panic("parser: unhandled char: " + string(c))
	}
}

func Parse(input string) *Obj {
	ps := &ParserState{input, 0}
	return ps.Parse()
	/*
		for node := ps.Parse(); node != nil; node = ps.Parse() {
			if node == RParen {
				panic("parser: extra closing parens ')' found")
			}
			if node == RSqaBr {
				panic("parser: extra closing square bracket ']' found")
			}
			if node == RCurBr {
				panic("parser: extra closing curly brace '}' found")
			}
		}
		return ast
	*/
}
