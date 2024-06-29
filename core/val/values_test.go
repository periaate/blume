package val

import (
	"fmt"
	"testing"
)

func TestTimeDate(t *testing.T) {
	res, _ := TimeDate("now")
	fmt.Println(res, "now")
	res, _ = TimeDate("now s")
	fmt.Println(res, "with sec")
	res, _ = TimeDate("M now")
	fmt.Println(res, "with date")
	res, _ = TimeDate("M")
	fmt.Println(res, "with date")
	res, _ = TimeDate("-1M")
	fmt.Println(res, "last month")
}
