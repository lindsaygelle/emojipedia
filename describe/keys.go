package describe

import "math"

type Keys []Key

func (keys Keys) Append(key ...Key) Keys {
	keys = append(keys, key...)
	return keys
}

func (keys Keys) Longest() (number int) {
	for _, key := range keys {
		number = int(math.Max(float64(number), float64(key.Literal.Len())))
	}
	return number
}

func (keys Keys) Prepend(key ...Key) Keys {
	keys = append((Keys{}).Append(key...), keys...)
	return keys
}
