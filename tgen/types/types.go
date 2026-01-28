package types

type Type string

type Types struct {
	Int      Type
	I8       Type
	I16      Type
	I32      Type
	I64      Type
	Uint     Type
	U8       Type
	U16      Type
	U32      Type
	U64      Type
	F32      Type
	F64      Type
	Uuid     Type
	String   Type
	Bool     Type
	Datetime Type
	Any      Type
	Array    Type
}

var TypesLookup = Types{
	Int:      Type("int"),
	I8:       Type("i8"),
	I16:      Type("i16"),
	I32:      Type("i32"),
	I64:      Type("i64"),
	Uint:     Type("uint"),
	U8:       Type("u8"),
	U16:      Type("u16"),
	U32:      Type("u32"),
	U64:      Type("u64"),
	F32:      Type("f32"),
	F64:      Type("f64"),
	Uuid:     Type("uuid"),
	String:   Type("string"),
	Bool:     Type("bool"),
	Datetime: Type("datetime"),
	Any:      Type("any"),
	Array:    Type("[]"),
}

type Options struct {
	Type     *Type `yaml:"type,omitempty"`
	Public   *bool `yaml:"public,omitempty"`
	Nullable *bool `yaml:"nullable,omitempty"`
	JsonTag  *bool `yaml:"jsontag,omitempty"`
}

type Props struct {
	Options
	UseDefaults *bool `yaml:"useDefaults"`
}

type GlobalConfig struct {
	Defaults Options `yaml:"defaults"`
}

type Field struct {
	Name     string
	Type     Type
	Public   bool
	Nullable bool
	JsonTag  bool
}

type Model struct {
	Name   string
	Fields []Field
}

type Models map[string]map[string]Props

type Manifest struct {
	Config GlobalConfig `yaml:"config"`
	Models Models       `yaml:"models"`
}

func (m Models) ToArray(defaults Options) []Model {
	var result []Model

	for modelName, fieldsMap := range m {
		currentModel := Model{Name: modelName}

		for fieldName, props := range fieldsMap {
			field := Field{Name: fieldName}
			opt := props.Options

			useDef := props.UseDefaults != nil && *props.UseDefaults
			field.Type = resolveType(opt.Type, defaults.Type, useDef)
			field.Public = resolveBool(opt.Public, defaults.Public, useDef)
			field.Nullable = resolveBool(opt.Nullable, defaults.Nullable, useDef)
			field.JsonTag = resolveBool(opt.JsonTag, defaults.JsonTag, useDef)

			currentModel.Fields = append(currentModel.Fields, field)
		}
		result = append(result, currentModel)
	}
	return result
}

func resolveBool(local *bool, fallback *bool, useDefault bool) bool {
	if local != nil {
		return *local
	}
	if useDefault && fallback != nil {
		return *fallback
	}
	return false
}

func resolveType(local *Type, fallback *Type, useDefault bool) Type {
	if local != nil {
		return *local
	}
	if useDefault && fallback != nil {
		return *fallback
	}
	return "string"
}
