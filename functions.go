package osquery

type DecayFunctions string

const (
	DecayFunctionGauss  DecayFunctions = "gauss"
	DecayFunctionExp    DecayFunctions = "exp"
	DecayFunctionLinear DecayFunctions = "linear"
)

type MultiValueModes string

const (
	MultiValueModeMin MultiValueModes = "min"
	MultiValueModeMax MultiValueModes = "max"
	MultiValueModeAvg MultiValueModes = "avg"
	MultiValueModeSum MultiValueModes = "sum"
)

type DecayFunction struct {
	field          string
	function       DecayFunctions
	params         DecayFunctionParams
	weight         *float64
	filter         Mappable
	multiValueMode MultiValueModes
}

type DecayFunctionParams struct {
	Origin interface{} `json:"origin,omitempty"`
	Scale  *float64    `json:"scale,omitempty"`
	Offset *float64    `json:"offset,omitempty"`
	Decay  *float64    `json:"decay,omitempty"`
}

func Decay(function DecayFunctions, field string) *DecayFunction {
	return &DecayFunction{
		field:    field,
		function: function,
		params:   DecayFunctionParams{},
	}
}

func (d *DecayFunction) Weight(weight float64) *DecayFunction {
	d.weight = &weight
	return d
}

func (d *DecayFunction) Filter(query Mappable) *DecayFunction {
	d.filter = query
	return d
}

func (d *DecayFunction) MultiValueMode(mode MultiValueModes) *DecayFunction {
	d.multiValueMode = mode
	return d
}

func (d *DecayFunction) Origin(origin any) *DecayFunction {
	d.params.Origin = origin
	return d
}

func (d *DecayFunction) Scale(scale float64) *DecayFunction {
	d.params.Scale = &scale
	return d
}

func (d *DecayFunction) Offset(offset float64) *DecayFunction {
	d.params.Offset = &offset
	return d
}

func (d *DecayFunction) Decay(decay float64) *DecayFunction {
	d.params.Decay = &decay
	return d
}

func (d *DecayFunction) Map() map[string]interface{} {
	res := map[string]interface{}{
		string(d.function): map[string]interface{}{
			d.field: map[string]interface{}{},
		},
	}

	pm := map[string]interface{}{}
	if d.params.Origin != nil {
		pm["origin"] = d.params.Origin
	}
	if d.params.Scale != nil {
		pm["scale"] = *d.params.Scale
	}
	if d.params.Offset != nil {
		pm["offset"] = *d.params.Offset
	}
	if d.params.Decay != nil {
		pm["decay"] = *d.params.Decay
	}
	res[string(d.function)].(map[string]interface{})[d.field] = pm
	if d.weight != nil {
		res["weight"] = *d.weight
	}
	if d.filter != nil {
		res["filter"] = d.filter.Map()
	}
	if d.multiValueMode != "" {
		res[string(d.function)].(map[string]interface{})["multi_value_mode"] = d.multiValueMode
	}
	return res
}

type WeightFunction struct {
	weight float64
	filter Mappable
}

func Weight(weight float64) *WeightFunction {
	return &WeightFunction{weight: weight}
}

func (w *WeightFunction) Filter(query Mappable) *WeightFunction {
	w.filter = query
	return w
}

func (w *WeightFunction) Map() map[string]interface{} {
	res := map[string]interface{}{
		"weight": w.weight,
	}
	if w.filter != nil {
		res["filter"] = w.filter.Map()
	}
	return res
}
