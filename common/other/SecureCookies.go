package other

import (
	"encoding/hex"

	"github.com/chmike/securecookie"
)

// generate a random key
var SecureKey []byte = securecookie.MustGenerateRandomKey()

// have key on a string
var KeyString string = hex.EncodeToString(SecureKey)

var SecureCookieObj = securecookie.MustNew("Auth", SecureKey, securecookie.Params{
	Path:     "/",              // cookie received only when URL starts with this path
	MaxAge:   48 * 60 * 60,     // cookie becomes invalid 3600 seconds after it is set
	HTTPOnly: true,             // disallow access by remote javascript code
	SameSite: securecookie.Lax, // cookie received with same or sub-domain names
	// ------------------------------- these will leave unset, but you can set them if you want on live server
	// Domain:   "example.com",    // cookie received only when URL domain matches this one
	// Secure: true, // cookie received only with HTTPS, never with HTTP
})
