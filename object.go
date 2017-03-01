package main

import (
	"fmt"
	"strconv"
)

type ObjType int

const (
	OTInt     ObjType = iota
	OTStr             = iota
	OTKey             = iota
	OTSym             = iota
	OTVec             = iota
	OTMap             = iota
	OTObj             = iota
	OTMessage         = iota
	OTRParen          = iota
	OTRSqaBr          = iota
	OTRCurBr          = iota
)

var RParen = &Obj{typ: OTRParen}
var RSqaBr = &Obj{typ: OTRSqaBr}
var RCurBr = &Obj{typ: OTRCurBr}

type Obj struct {
	typ  ObjType
	intv int64
	strv string
	vecv []*Obj
	mapv map[string]*Obj
	objv *Obj
}

func NewInt(value int64) *Obj {
	return &Obj{typ: OTInt, intv: value}
}

func NewStr(value string) *Obj {
	return &Obj{typ: OTStr, strv: value}
}

func NewKey(value string) *Obj {
	return &Obj{typ: OTKey, strv: value}
}

var Keywords = map[string]*Obj{}

func InternKey(value string) *Obj {
	if s, ok := Keywords[value]; ok {
		return s
	}
	Keywords[value] = NewKey(value)
	return Keywords[value]
}

func NewSym(value string) *Obj {
	return &Obj{typ: OTSym, strv: value}
}

var Symbols = map[string]*Obj{}

func InternSym(value string) *Obj {
	if s, ok := Symbols[value]; ok {
		return s
	}
	Symbols[value] = NewSym(value)
	return Symbols[value]
}

func NewVec(value []*Obj) *Obj {
	return &Obj{typ: OTVec, vecv: value}
}

func NewMap(value map[string]*Obj) *Obj {
	return &Obj{typ: OTVec, mapv: value}
}

func NewMessage(calleeSym *Obj, argsVec *Obj) *Obj {
	return &Obj{typ: OTMessage, objv: calleeSym, vecv: argsVec.vecv}
}

func (o *Obj) String() string {
	switch o.typ {
	case OTInt:
		return strconv.FormatInt(o.intv, 10)
	case OTStr:
		return fmt.Sprintf("\"%s\"", o.strv)
	case OTKey:
		return fmt.Sprintf(":%s", o.strv)
	case OTSym:
		return o.strv
	case OTVec:
		ret := "["
		for i, v := range o.vecv {
			if i != 0 {
				ret += " "
			}
			ret += v.String()
		}
		ret += "]"
		return ret
	case OTMap:
		ret := "map["
		i := 0
		for k, v := range o.mapv {
			if i != 0 {
				ret += " "
			}
			ret += ":" + k
			ret += " "
			ret += v.String()
			i++
		}
		ret += "]"
		return ret
	case OTObj:
		return "Obj(...)"
	case OTCall:
		ret := o.objv.strv + "("
		for i, v := range o.vecv {
			if i != 0 {
				ret += " "
			}
			ret += v.String()
		}
		ret += ")"
		return ret
	default:
		panic("obj: string: unhandled obj type")
	}
}
