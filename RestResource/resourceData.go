package restresource

type ResourceData interface {
	AsValue() (interface{}, bool)
	AsSlice() ([]ResourceData, bool)
	AsMap() (map[string]ResourceData, bool)
}

type resourceValue struct {
	value interface{}
}

func (rv *resourceValue) AsValue() (interface{}, bool) {
	return rv.value, true
}

func (rv *resourceValue) AsSlice() ([]ResourceData, bool) {
	return nil, false
}

func (rv *resourceValue) AsMap() (map[string]ResourceData, bool) {
	return nil, false
}

type resourceSlice struct {
	Values []ResourceData
}

func (rs *resourceSlice) AsValue() (interface{}, bool) {

	return nil, false
}

func (rs *resourceSlice) AsSlice() ([]ResourceData, bool) {

	return rs.Values, true
}

func (rs *resourceSlice) AsMap() (map[string]ResourceData, bool) {
	return nil, false
}

type resourceMap struct {
	Values map[string]ResourceData
}

func (rm *resourceMap) AsValue() (interface{}, bool) {
	return nil, false
}

func (rm *resourceMap) AsSlice() ([]ResourceData, bool) {
	return nil, false
}

func (rm *resourceMap) AsMap() (map[string]ResourceData, bool) {
	return rm.Values, true
}
