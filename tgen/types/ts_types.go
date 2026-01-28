package types

var TsTypes = Types{
	Int:      Type("number"),
	I8:       Type("number"),
	I16:      Type("number"),
	I32:      Type("number"),
	I64:      Type("number"),
	Uint:     Type("number"),
	U8:       Type("number"),
	U16:      Type("number"),
	U32:      Type("number"),
	U64:      Type("number"),
	F32:      Type("number"),
	F64:      Type("number"),
	Uuid:     Type("uuid.UUID"),
	String:   Type("string"),
	Bool:     Type("boolean"),
	Datetime: Type("string"),
	Any:      Type("any"),
	Array:    Type("[]"),
}
