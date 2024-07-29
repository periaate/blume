package str

//
// func TestCapture(t *testing.T) {
// 	tst := `Hello, "World"! This is {my test} for .Delimiter,`
// 	delims := []string{`"`, `"`, "{", "}", ".", ","}
// 	res, err := CaptureDelims(tst, delims...)
// 	if err != nil {
// 		t.Fatalf("error: %v", err)
// 	}
// 	for i, r := range res {
// 		fmt.Println(i+1, r)
// 	}
//
// 	if len(res) != 6 {
// 		t.Fatalf("expected 6, got %d", len(res))
// 	}
//
// }
