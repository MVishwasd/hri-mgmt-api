/**
 * (C) Copyright IBM Corp. 2020
 *
 * SPDX-License-Identifier: Apache-2.0
 */
package auth

import (
	"strings"

	"gopkg.in/square/go-jose.v2/jwt"
)

// ClaimsHolder is an interface to support testing
type ClaimsHolder interface {
	Claims(claims interface{}) error
}

type HriClaims struct {
	// jwt.Audience can marshal string or []string json types
	Audience jwt.Audience `json:"aud"`
	Subject  string       `json:"sub"`

	// Some OAuth services use `scopes` (IBM AppID) and other use `roles` (Azure AD).
	// This extracts both if present and HasScope() searches both
	Scope string   `json:"scope"`
	Roles []string `json:"roles"`
}

func (c HriClaims) HasScope(claim string) bool {
	// split space-delimited scope string into an array
	scopes := strings.Fields(c.Scope)

	for _, val := range scopes {
		if val == claim {
			// token contains claim for this scope
			return true
		}
	}

	return false
}
