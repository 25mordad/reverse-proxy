package main

import(
        "net/url"
        "net/http"
        "net/http/httputil"
)

func main() {
        ////change the URL to what you need
        remote, err := url.Parse("http://app.devel/")
        if err != nil {
                panic(err)
        }

        handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
                return func(w http.ResponseWriter, r *http.Request) {
                        r.Host = remote.Host
                        p.ServeHTTP(w, r)
                }
        }

        proxy := httputil.NewSingleHostReverseProxy(remote)
        http.HandleFunc("/", handler(proxy))
        //set a new port
        err = http.ListenAndServe(":8082", nil)
        if err != nil {
                panic(err)
        }
}
