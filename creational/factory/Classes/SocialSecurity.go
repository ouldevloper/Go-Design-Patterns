package classes

import "time"

type SocialSecurity struct {
	Id        int
	Number    int
	StartedAt time.Time
	EndAt     time.Time
}
