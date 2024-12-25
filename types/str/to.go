package str

import "strconv"

func ToInt(s string) (int, error) { return strconv.Atoi(s) }
func ToInt8(s string) (int8, error) {
	i, err := strconv.ParseInt(s, 10, 8)
	return int8(i), err
}
func ToInt16(s string) (int16, error) {
	i, err := strconv.ParseInt(s, 10, 16)
	return int16(i), err
}
func ToInt32(s string) (int32, error) {
	i, err := strconv.ParseInt(s, 10, 32)
	return int32(i), err
}
func ToInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return int64(i), err
}
func ToUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 0)
	return uint(i), err
}
func ToUint8(s string) (uint8, error) {
	i, err := strconv.ParseUint(s, 10, 8)
	return uint8(i), err
}
func ToUint16(s string) (uint16, error) {
	i, err := strconv.ParseUint(s, 10, 16)
	return uint16(i), err
}
func ToUint32(s string) (uint32, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	return uint32(i), err
}
func ToUint64(s string) (uint64, error) {
	i, err := strconv.ParseUint(s, 10, 64)
	return uint64(i), err
}
func ToFloat32(s string) (float32, error) {
	i, err := strconv.ParseFloat(s, 32)
	return float32(i), err
}
func ToFloat64(s string) (float64, error) {
	i, err := strconv.ParseFloat(s, 64)
	return float64(i), err
}
