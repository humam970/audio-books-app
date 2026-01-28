package types

var GoTypes = Types{
	Int:      Type("int"),
	I8:       Type("int8"),
	I16:      Type("int16"),
	I32:      Type("int32"),
	I64:      Type("int64"),
	Uint:     Type("uint"),
	U8:       Type("uint8"),
	U16:      Type("uint16"),
	U32:      Type("uint32"),
	U64:      Type("uint64"),
	F32:      Type("float32"),
	F64:      Type("float64"),
	Uuid:     Type("uuid.UUID"),
	String:   Type("string"),
	Bool:     Type("bool"),
	Datetime: Type("time.Time"),
	Any:      Type("any"),
	Array:    Type("[]"),
}
