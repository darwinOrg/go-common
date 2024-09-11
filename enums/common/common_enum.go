package comm_enum

const (
	No  int8 = 0
	Yes int8 = 1
)

var (
	Bool2IntMap = map[bool]int8{
		false: No,
		true:  Yes,
	}

	Int2BoolMap = map[int8]bool{
		No:  false,
		Yes: true,
	}
)
