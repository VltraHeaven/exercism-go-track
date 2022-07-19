package erratum

import (
	"reflect"
)

func Use(opener ResourceOpener, input string) (err error) {
	var o Resource
	var ok bool

	for o, err = opener(); err != nil; {
		if _, ok = err.(TransientError); !ok {
			return err
		}
		o, err = opener()
	}

	defer func() {
		if r := recover(); r != nil {
			if _, ok = r.(FrobError); ok {
				err = r.(FrobError)
				o.Defrob(reflect.ValueOf(err).Field(0).String())
			} else {
				err = r.(error)
			}
		}
		o.Close()
	}()
	o.Frob(input)
	return
}
