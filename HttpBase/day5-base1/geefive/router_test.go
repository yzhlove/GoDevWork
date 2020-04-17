package geefive

import (
	"fmt"
	"testing"
)

func Test_Router(t *testing.T) {

	router := newRouter()
	router.addRouter("GET", "/v1/hello", nil)
	router.addRouter("GET", "/v2/hello", nil)

	fmt.Println(router.getRouter("GET", "/v1/hello"))
	fmt.Println(router.getRouter("GET", "/v2/hello"))

}
