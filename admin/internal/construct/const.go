package construct

import "time"

type ClaimsUIDContextKey struct{}

type ClaimsUserContextKey struct{}

type Claims struct {
	UID       int64
	Sub       string
	Expire    time.Duration
	JwtSecret string
}
