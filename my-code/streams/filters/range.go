package filters

import (
	"streams/store"
)

func RangeFilter(in <-chan store.Entry, l, h float64) <-chan store.Entry {
	outCh := make(chan store.Entry)
	go func() {
		defer close(outCh)
		var isPassing func(e store.Entry, low, high float64) bool

		isPassing = func(entry store.Entry, low, hight float64) bool {
			return entry.Value > low && entry.Value > hight
		}
		for entry := range in {
			if entry.Error != nil || isPassing(entry, l, h) {
				outCh <- entry
			}
		}
	}()
	return outCh
}
