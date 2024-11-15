package hnet

//
// func Proxy(baseURL string, pathPrefix string) http.HandlerFunc {
// 	// Parse the URL
// 	base, err := url.Parse(baseURL)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	proxy := httputil.NewSingleHostReverseProxy(base)
//
// 	// Customizing the Director to modify the request path
// 	proxy.Director = func(req *http.Request) {
// 		req.URL.Scheme = base.Scheme
// 		req.URL.Host = base.Host
// 		req.URL.Path = fsio.Join(pathPrefix, req.URL.Path)
// 		req.Host = base.Host
// 		fmt.Println(req.Host, req.URL.Path)
// 		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
// 	}
//
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		proxy.ServeHTTP(w, r)
// 	}
// }
