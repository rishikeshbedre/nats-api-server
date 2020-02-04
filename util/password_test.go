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
	
	inp2 := "abcdefghijklmnopqrstuvwxyz!@#$%^&*()_+{};:.,~1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	_, err2 := GenHashPassword(inp2)
	if err2 != nil {
		t.Errorf("Test failed, string '%s' supplied but failed", inp2)
	}
}
