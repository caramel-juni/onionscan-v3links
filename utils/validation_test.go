package utils

import (
	"testing"
)

func TestIsOnion(t *testing.T) {
	// Valid v2 onion
	test := IsOnion("maschqwtr4lqt2pj.onion")
	if !test {
		t.Errorf("%s should have matched!", "maschqwtr4lqt2pj.onion")
	}

	// Valid v2 onion with subdomains
	test = IsOnion("subdomain.lots.of.maschqwtr4lqt2pj.onion")
	if !test {
		t.Errorf("%s should have matched!", "subdomain.lots.of.maschqwtr4lqt2pj.onion")
	}

	// Invalid onions (bad chars, wrong TLD)
	test = IsOnion("thisis9notavalidonion.onion")
	if test {
		t.Errorf("%s should not have matched!", "thisis9notavalidonion.onion")
	}

	test = IsOnion("maschqwtr4lqt2pj.com")
	if test {
		t.Errorf("%s should not have matched!", "maschqwtr4lqt2pj.com")
	}

	test = IsOnion("www.thisisnotanonionitistoolong.onion")
	if test {
		t.Errorf("%s should not have matched!", "www.thisisnotanonionitistoolong.onion")
	}

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
