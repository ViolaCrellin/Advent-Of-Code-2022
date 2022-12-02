package util

type Bitmask uint32

func (f Bitmask) HasFlag(flag Bitmask) bool { return f&flag != 0 }
func (f Bitmask) IsNthBitSet(number, n int) bool {
	value := number & (1 << (n - 1))
	return value != 0
}
func (f *Bitmask) AddFlag(flag Bitmask)    { *f |= flag }
func (f *Bitmask) ClearFlag(flag Bitmask)  { *f &= ^flag }
func (f *Bitmask) ToggleFlag(flag Bitmask) { *f ^= flag }

func SplitIntArrayByBitmask(values []int, bit int) map[string][]int {
	var result = make(map[string][]int, 2)

	for _, value := range values {
		bitvalue := Bitmask(value)
		if bitvalue.IsNthBitSet(value, bit) {
			result["HasFlag"] = append(result["HasFlag"], value)
		} else {
			result["NoFlag"] = append(result["NoFlag"], value)
		}
	}

	return result
}
