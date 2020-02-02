package util
 
import (
	"testing"
)

func TestGenHashPassowrd(t *testing.T){
	inp := "abc"
	_, err := GenHashPassword(inp)
	if err != nil {
		t.Errorf("Test failed, string '%s' supplied but failed", inp)
	}
}