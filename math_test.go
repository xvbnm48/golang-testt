package go_testing_baru

import "testing"

func TestMath(t *testing.T) {
	c:= Math(-5)
	 if c != 5 {
		 t.Logf("expect 5 ,but got %d" ,c)
		 t.Fail()
	 }

}
