package match

type Action int

const (
	Nothing Action = iota
	Replaces
	Deletes
	Stops
)

type Match[I, O any] interface {
	All(input I) [][]int
	Next(input I) (from int, to int, found bool)
	Found(input I) (value O, action Action)
}

func StartsWith[I, O any]() (res Match[I, O]) { return }
func EndsWith  [I, O any]() (res Match[I, O]) { return }
func Contains  [I, O any]() (res Match[I, O]) { return }
func IsExactly [I, O any]() (res Match[I, O]) { return }
func PrecededBy[I, O any]() (res Match[I, O]) { return }
func FollowedBy[I, O any]() (res Match[I, O]) { return }
func PatternOf [I, O any]() (res Match[I, O]) { return }

func Split[Arr, El any](src Iter[Arr, El], matcher Match[I, O]) [][]T

