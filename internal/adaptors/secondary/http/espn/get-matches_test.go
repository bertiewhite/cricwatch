package espn

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_getMatches(t *testing.T) {
	client := NewClient(http.DefaultClient)

	matches, err := client.GetMatches()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("%+v", matches))
}
