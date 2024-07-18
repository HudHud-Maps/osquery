package osquery

type Params map[string]interface{}

type ScriptField struct {
	name   string
	source string
	params Params
}

func (s *ScriptField) Name() string {
	return s.name
}

func (s *ScriptField) Map() map[string]interface{} {
	r := map[string]interface{}{}
	if s.source != "" {
		r["source"] = s.source
	}
	if len(s.params) > 0 {
		r["params"] = s.params
	}
	return map[string]interface{}{
		"script": r,
	}
}

func Script(name string) ScriptField {
	return ScriptField{
		name: name,
	}
}

func (s ScriptField) Source(source string) ScriptField {
	s.source = source
	return s
}

func (s ScriptField) Params(params Params) ScriptField {
	s.params = params
	return s
}
