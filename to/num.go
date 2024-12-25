package to

import "strconv"

func Int(s string) (int, error) { return strconv.Atoi(s) }
func Int8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	return int8(i), err
}
func Int16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	return int16(i), err
}
func Int32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}
func Int64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int64(i), err
}
func Uint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	return uint(i), err
}
func Uint8(s string) (uint8, error) {
	i, err := strconv.ParseUint(s, 10, 8)
	return uint8(i), err
}
func Uint16(s string) (uint16, error) {
	i, err := strconv.ParseUint(s, 10, 16)
	return uint16(i), err
}
func Uint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	return uint32(i), err
}
func Uint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	return uint64(i), err
}
func Float32(s string) (float32, error) {
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}
func Float64(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	return float64(i), err
}
