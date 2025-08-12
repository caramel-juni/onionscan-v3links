package utils

import (
	"strings"
	"testing"
)

func TestDebugOnionValidation(t *testing.T) {
	testOnions := []string{
		"maschqwtr4lqt2pj.onion",                                         // v2 valid
		"subdomain.lots.of.maschqwtr4lqt2pj.onion",                       // v2 valid with subdomain
		"thisis9notavalidonion.onion",                                    // invalid (bad chars)
		"maschqwtr4lqt2pj.com",                                           // invalid (wrong TLD)
		"www.thisisnotanonionitistoolong.onion",                          // invalid (too long)
		"dy3vbeecmljb2476ua26vasan24qezytsqb7rhiokqnw3jxzjsdse6ad.onion", // v3 valid
		"dy3vbeecmljb2476ua26vasan24qezytsqb7rhiokqnw3jxzjsdse6a.onion",  // v3 invalid (length)
	}

	for _, onion := range testOnions {
		t.Logf("Testing: %s", onion)

		if !strings.HasSuffix(onion, ".onion") {
			t.Log(" - FAIL: Does not have .onion suffix")
			continue
		}

		core := strings.TrimSuffix(onion, ".onion")
		t.Logf(" - Core onion host: %s (length %d)", core, len(core))

		isOnion := IsOnion(onion)
		t.Logf(" - IsOnion(): %v", isOnion)

		isV2 := IsOnionV2(onion)
		isV3 := IsOnionV3(onion)
		t.Logf(" - IsOnionV2(): %v", isV2)
		t.Logf(" - IsOnionV3(): %v", isV3)

		t.Log("")
	}
}
