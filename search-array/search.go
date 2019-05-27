package search

import (
	"reflect"
)

// SearchInArray func
func SearchInArray(val interface{}, array interface{}) float64 {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			// _, div := math.Modf(val.(float64))
			// _, sdiv := math.Modf(s.Index(i).Float())

			if val.(float64) == s.Index(i).Float() {
				return s.Index(i).Float()
			} // end

			if val.(float64) < s.Index(i).Float() {
				if val.(float64)-s.Index(i-1).Float() < s.Index(i).Float()-val.(float64) {
					return s.Index(i - 1).Float()
				}

				if val.(float64)-s.Index(i-1).Float() > s.Index(i).Float()-val.(float64) {
					return s.Index(i + 1).Float()
				}

			} // end

		}
	}

	return 0
}
