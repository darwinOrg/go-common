package result

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestVoidJson(t *testing.T) {

	ret := Success(VoidValue)

	j, err := json.Marshal(ret)
	fmt.Println(string(j), err)
	//t.Log(string(j), err)

	v := &Result[Void]{}

	err = json.Unmarshal(j, v)
	fmt.Println(err)
}
