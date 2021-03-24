package util

import "github.com/google/uuid"

func UUID() uuid.UUID {
	return uuid.New()
}
