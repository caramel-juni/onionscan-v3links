package utils

import (
	"testing"
)

func TestIsOnion(t *testing.T) {

	// Valid v3 onion
	v3Onion := "dy3vbeecmljb2476ua26vasan24qezytsqb7rhiokqnw3jxzjsdse6ad.onion"
	if !IsOnion(v3Onion) {
		t.Errorf("%s should have matched as a valid v3 onion!", v3Onion)
	}

	// Invalid v3 onion (wrong length)
	invalidV3 := "dy3vbeecmljb2476ua26vasan24qezytsqb7rhiokqnw3jxzjsdse6a.onion" // 55 chars
	if IsOnion(invalidV3) {
		t.Errorf("%s should NOT have matched as a valid v3 onion!", invalidV3)
	}
}
