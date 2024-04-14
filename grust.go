package grust

type Ok func() interface{}
type Err func() interface{}
type MapFn func(interface{}) interface{}
type AndThenFn func(interface{}) *Result

type Result struct {
	Ok    Ok
	IsOk  bool
	Err   Err
	IsErr bool
	Panic bool
}

func (r *Result) AndThen(fn AndThenFn) *Result {
	if r.IsErr {
		return r
	}

	return fn(r.Ok())
}

func (r *Result) Map(fn MapFn) *Result {
	if r.IsErr {
		return r
	}

	mappedValue := fn(r.Ok())
	return &Result{
		Ok:    func() interface{} { return mappedValue },
		IsOk:  true,
		Err:   r.Err,
		IsErr: false,
	}
}

func (r *Result) OrElse(defaultValue interface{}) interface{} {
	if r.IsErr {
		return defaultValue
	}
	return r.Ok()
}

func (r *Result) Unwrap() interface{} {
	if r.IsErr {
		if r.Panic {
			panic(r.Err())
		} else {
			return r.Err()
		}
	}
	return r.Ok()
}

func (r *Result) UnwrapOr(defaultValue interface{}) interface{} {
	if r.IsErr {
		return defaultValue
	}
	return r.Ok()
}

func (r *Result) UnwrapErr() interface{} {
	if r.IsErr {
		return r.Err()
	}
	if r.Panic {
		panic(r.Err())
	} else {
		return r.Err()
	}
}
