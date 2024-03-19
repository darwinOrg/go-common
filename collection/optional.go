package dgcoll

type Optional struct {
	value any
}

func NewOptional(value any) *Optional {
	return &Optional{value: value}
}

func (o *Optional) Map(mapFunc Function[any, any]) *Optional {
	o.value = mapFunc(o.value)
	return o
}

func (o *Optional) Filter(filterFunc Predicate[any]) *Optional {
	if o.value == nil || filterFunc(o.value) {
		return o
	}

	o.value = nil
	return o
}

func (o *Optional) OrElse(value any) any {
	if o.value != nil {
		return o.value
	} else {
		return value
	}
}

func (o *Optional) IsPresent() bool {
	return o.value != nil
}

func (o *Optional) IsEmpty() bool {
	return o.value == nil
}

func (o *Optional) Get() any {
	return o.value
}
