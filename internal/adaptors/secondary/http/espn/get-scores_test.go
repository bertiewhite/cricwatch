package espn

import (
	"fmt"
	"net/http"
	"testing"
)

func Test_GetScores(t *testing.T) {

	client := NewEspnClient(http.DefaultClient)
	score, err := client.GetScore(1336129)

	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("+%v", score))
}
