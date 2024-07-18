package osquery

type ScoreModes string

const (
	ScoreModeMultiply ScoreModes = "multiply"
	ScoreModeSum      ScoreModes = "sum"
	ScoreModeAvg      ScoreModes = "avg"
	ScoreModeFirst    ScoreModes = "first"
	ScoreModeMax      ScoreModes = "max"
	ScoreModeMin      ScoreModes = "min"
)

type BoostModes string

const (
	BoostModeMultiply BoostModes = "multiply"
	BoostModeReplace  BoostModes = "replace"
	BoostModeSum      BoostModes = "sum"
	BoostModeAvg      BoostModes = "avg"
	BoostModeMax      BoostModes = "max"
	BoostModeMin      BoostModes = "min"
)

type FunctionScore struct {
	query     Mappable
	boost     *float64
	maxBoost  *float64
	scoreMode ScoreModes
	boostMode BoostModes
	functions []Mappable
}

func NewFunctionScore() *FunctionScore {
	return &FunctionScore{
		functions: []Mappable{},
	}
}

func (fs *FunctionScore) Query(q Mappable) *FunctionScore {
	fs.query = q
	return fs
}

func (fs *FunctionScore) Boost(boost float64) *FunctionScore {
	fs.boost = &boost
	return fs
}

func (fs *FunctionScore) MaxBoost(maxBoost float64) *FunctionScore {
	fs.maxBoost = &maxBoost
	return fs
}

func (fs *FunctionScore) ScoreMode(mode ScoreModes) *FunctionScore {
	fs.scoreMode = mode
	return fs
}

func (fs *FunctionScore) BoostMode(mode BoostModes) *FunctionScore {
	fs.boostMode = mode
	return fs
}

func (fs *FunctionScore) Functions(functions ...Mappable) *FunctionScore {
	fs.functions = append(fs.functions, functions...)
	return fs
}

func (fs *FunctionScore) Map() map[string]interface{} {
	m := make(map[string]interface{})

	if fs.query != nil {
		m["query"] = fs.query.Map()
	}

	if fs.boost != nil {
		m["boost"] = *fs.boost
	}

	if fs.maxBoost != nil {
		m["max_boost"] = *fs.maxBoost
	}

	if fs.scoreMode != "" {
		m["score_mode"] = fs.scoreMode
	}

	if fs.boostMode != "" {
		m["boost_mode"] = fs.boostMode
	}

	if len(fs.functions) > 0 {
		funcs := []map[string]interface{}{}
		for _, f := range fs.functions {
			funcs = append(funcs, f.Map())
		}
		m["functions"] = funcs
	}

	return map[string]interface{}{
		"function_score": m,
	}
}
