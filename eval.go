package main

func EnvGet(env *Obj, key string) *Obj {
	if v, ok := env.mapv[key]; ok {
		return v
	}
	if env.objv != nil { // Check in parent env
		return EnvGet(env.objv, key)
	}
	return nil
}

func Eval(env *Obj, val *Obj) *Obj {
	switch val.typ {
	case OTObj:
		fallthrough
	case OTInt:
		fallthrough
	case OTStr:
		fallthrough
	case OTKey:
		return val
	case OTSym:
		return EvalSym(env, val)
	case OTVec:
		return EvalVec(env, val)
	case OTMap:
		return EvalMap(env, val)
	case OTCall:
		return EvalCall(env, val)
	default:
		panic("eval: unhandled value type given")
	}
}

func EvalSym(env *Obj, val *Obj) *Obj {
	return EnvGet(env, "nil")
}

func EvalVec(env *Obj, val *Obj) *Obj {
	ret := NewVec(val.vecv)
	for i, v := range ret.vecv {
		ret.vecv[i] = Eval(env, v)
	}
	return ret
}

func EvalMap(env *Obj, val *Obj) *Obj {
	ret := NewMap(val.mapv)
	for k, v := range ret.mapv {
		ret.mapv[k] = Eval(env, v)
	}
	return ret
}

func EvalCall(env *Obj, val *Obj) *Obj {
	callee := Eval(env, InternSym(val.strv))
	return nil
}
