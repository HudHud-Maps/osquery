package osquery

type Unit string

const (
	Meters     Unit = "m"
	Kilometers Unit = "km"
	Miles      Unit = "mi"
	Feet       Unit = "ft"
)

type DistanceType string

const (
	ArcDistance   DistanceType = "arc"
	PlaneDistance DistanceType = "plane"
)

type GeoSortMode string

const (
	GeoSortModeMin    GeoSortMode = "min"
	GeoSortModeMax    GeoSortMode = "max"
	GeoSortModeAvg    GeoSortMode = "avg"
	GeoSortModeMedian GeoSortMode = "median"
)

type GeoDistanceSort struct {
	name           string
	points         [][2]float64
	order          Order
	unit           Unit
	distanceType   DistanceType
	mode           GeoSortMode
	ignoreUnmapped *bool
}

func GeoSort(name string) *GeoDistanceSort {
	return &GeoDistanceSort{
		name: name,
	}
}

func (s *GeoDistanceSort) Points(point ...[2]float64) *GeoDistanceSort {
	s.points = append(s.points, point...)
	return s
}

func (s *GeoDistanceSort) Order(order Order) *GeoDistanceSort {
	s.order = order
	return s
}

func (s *GeoDistanceSort) Unit(unit Unit) *GeoDistanceSort {
	s.unit = unit
	return s
}

func (s *GeoDistanceSort) DistanceType(distanceType DistanceType) *GeoDistanceSort {
	s.distanceType = distanceType
	return s
}

func (s *GeoDistanceSort) Mode(mode GeoSortMode) *GeoDistanceSort {
	s.mode = mode
	return s
}

func (s *GeoDistanceSort) IgnoreUnmapped(ignoreUnmapped bool) *GeoDistanceSort {
	s.ignoreUnmapped = &ignoreUnmapped
	return s
}

func (s *GeoDistanceSort) Map() map[string]interface{} {
	r := make(map[string]interface{})
	if len(s.points) > 0 {
		r[s.name] = s.points
	}
	if s.order != "" {
		r["order"] = s.order
	}
	if s.unit != "" {
		r["unit"] = s.unit
	}
	if s.distanceType != "" {
		r["distance_type"] = s.distanceType
	}
	if s.mode != "" {
		r["mode"] = s.mode
	}
	if s.ignoreUnmapped != nil {
		r["ignore_unmapped"] = *s.ignoreUnmapped
	}
	return map[string]interface{}{
		"_geo_distance": r,
	}
}

func Sort(name string, order Order) *FieldSort {
	return &FieldSort{
		name:  name,
		order: order,
	}
}

type FieldSort struct {
	name  string
	order Order
}

func (s FieldSort) Map() map[string]interface{} {
	r := map[string]interface{}{
		s.name: map[string]interface{}{
			"order": s.order,
		},
	}
	return r
}
