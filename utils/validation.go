/*
	package utils

import (

	"regexp"
	"strings"

)

	func IsOnion(identifier string) bool {
		// TODO: At some point we will want to support i2p
		if len(identifier) >= 22 && strings.HasSuffix(identifier, ".onion") {
			matched, _ := regexp.MatchString(`(^|\.)[a-z2-7]{16}\.onion$`, identifier)
			return matched
		}
		return false
	}
*/
package utils

import (
	"regexp"
	"strings"
)

func IsOnion(identifier string) bool {
	if !strings.HasSuffix(identifier, ".onion") {
		return false
	}

	// Remove the .onion suffix
	noSuffix := strings.TrimSuffix(identifier, ".onion")

	// Get last label (the onion address part), ignoring any subdomains
	parts := strings.Split(noSuffix, ".")
	onionAddress := parts[len(parts)-1]

	switch len(onionAddress) {
	case 16:
		return isValidOnionV2(onionAddress)
	case 56:
		return isValidOnionV3(onionAddress)
	default:
		return false
	}
}

func IsOnionV2(identifier string) bool {
	if !strings.HasSuffix(identifier, ".onion") {
		return false
	}
	core := strings.TrimSuffix(identifier, ".onion")
	return len(core) == 16 && isValidOnionV2(core)
}

func IsOnionV3(identifier string) bool {
	if !strings.HasSuffix(identifier, ".onion") {
		return false
	}
	core := strings.TrimSuffix(identifier, ".onion")
	return len(core) == 56 && isValidOnionV3(core)
}

func isValidOnionV2(core string) bool {
	matched, _ := regexp.MatchString(`^[a-z2-7]{16}$`, core)
	return matched
}

func isValidOnionV3(core string) bool {
	matched, _ := regexp.MatchString(`^[a-z2-7]{56}$`, core)
	return matched
}
