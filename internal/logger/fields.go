package logger

/*
type Field struct {
	Key   string
	Value interface{}
}

type OFields []Field

func (s OFields) toFlatIfaces() []interface{} {
	if s == nil {
		return nil
	}

	ifaces := make([]interface{}, len(s)*2)

	for i, f := range s {
		ifaces[i] = f.Key
		ifaces[i+1] = f.Value
	}

	return ifaces
}
*/

type Fields map[string]interface{}

type FieldSet []Fields

func (s FieldSet) Merge() Fields {
	if s == nil {
		return nil
	}

	fields := make(Fields)

	for _, fs := range s {
		for k, v := range fs {
			fields[k] = v
		}
	}

	return fields
}
