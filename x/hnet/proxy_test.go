package hnet

//
// func TestProxy(t *testing.T) {
// 	var hit bool
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("GET test/", func(w http.ResponseWriter, r *http.Request) {
// 		hit = true
// 		w.WriteHeader(http.StatusOK)
// 	})
//
// 	// Create a test server
// 	ts := httptest.NewServer(mux)
// 	defer ts.Close()
// 	uri := ts.URL
//
// 	p := Proxy(uri, "test")
//
// 	ps := httptest.NewServer(p)
// 	defer ps.Close()
//
// 	tar := fsio.Join(ps.URL)
// 	clog.Info("target", "url", tar)
//
// 	// Create a request to pass to our handler
// 	resp, err := http.Get(tar)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
//
// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("expected status OK; got %v", resp.Status)
// 	}
//
// 	if !hit {
// 		t.Fatal("expected handler to be hit")
// 	}
// }
