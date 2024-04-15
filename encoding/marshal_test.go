package encoding

type testItem struct {
	Name        string
	Quantity    int
	IsAvailable bool
	Price       float64
}

func newTestItem1() testItem {
	return testItem{"widget", 15, true, 45.2531}
}
func newTestItem2() testItem {
	return testItem{"thingy", 7, false, 13.84}
}
