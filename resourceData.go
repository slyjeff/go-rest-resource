package restresource

type ResourceData interface {
	AsString() (string, bool)
	AsSlice() ([]ResourceData, bool)
	AsMap() (map[string]ResourceData, bool)
}

type resourceValue struct {
	value string
}

func (rv *resourceValue) AsString() (string, bool) {
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

func (rs *resourceSlice) AsString() (string, bool) {
	return "", false
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

func (rm *resourceMap) AsString() (string, bool) {
	return "", false
}

func (rm *resourceMap) AsSlice() ([]ResourceData, bool) {
	return nil, false
}

func (rm *resourceMap) AsMap() (map[string]ResourceData, bool) {
	return rm.Values, true
}
