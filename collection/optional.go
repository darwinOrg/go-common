package dgcoll

type Optional struct {
	value any
}

var emptyOptional = &Optional{value: nil}

func NewOptional(value any) *Optional {
	if value == nil {
		return emptyOptional
	}

	return &Optional{value: value}
}

func EmptyOptional() *Optional {
	return emptyOptional
}

func (o *Optional) Get() any {
	return o.value
}

func (o *Optional) IsPresent() bool {
	return o.value != nil
}

func (o *Optional) IsEmpty() bool {
	return o.value == nil
}

func (o *Optional) IfPresent(consumer Consumer[any]) {
	if o.value != nil {
		consumer(o.value)
	}
}

func (o *Optional) IfPresentOrElse(consumer Consumer[any], emptyAction Runnable) {
	if o.value != nil {
		consumer(o.value)
	} else {
		emptyAction()
	}
}

func (o *Optional) Filter(filterFunc Predicate[any]) *Optional {
	if o.value == nil {
		return emptyOptional
	}

	if filterFunc(o.value) {
		return o
	}

	return emptyOptional
}

func (o *Optional) Map(mapFunc Function[any, any]) *Optional {
	if o.value == nil {
		return emptyOptional
	}

	return NewOptional(mapFunc(o.value))
}

func (o *Optional) FlatMap(mapFunc Function[any, *Optional]) *Optional {
	if o.value == nil {
		return emptyOptional
	}

	return mapFunc(o.value)
}

func (o *Optional) Or(supplier Supplier[*Optional]) *Optional {
	if o.value != nil {
		return o
	}

	return supplier()
}

func (o *Optional) OrElse(value any) any {
	if o.value != nil {
		return o.value
	} else {
		return value
	}
}

func (o *Optional) OrElseGet(supplier Supplier[any]) any {
	if o.value != nil {
		return o.value
	} else {
		return supplier()
	}
}

func (o *Optional) OrElseError(err error) (any, error) {
	if o.value != nil {
		return o.value, nil
	} else {
		return nil, err
	}
}

func (o *Optional) OrElseSupplyError(supplier Supplier[error]) (any, error) {
	if o.value != nil {
		return o.value, nil
	} else {
		return nil, supplier()
	}
}
