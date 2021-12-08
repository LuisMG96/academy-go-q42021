package common

type Filter struct {
	TypeFilter     string
	Items          int64
	ItemsPerWorker int64
}

func NewFilter(typeFilter string, items int64, itemsPerWorker int64) *Filter {
	return &Filter{
		typeFilter,
		items,
		itemsPerWorker,
	}
}
