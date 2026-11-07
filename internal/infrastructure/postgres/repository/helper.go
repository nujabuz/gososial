package repository

import (
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func timeToTimestamptz(t time.Time) pgtype.Timestamptz {
	var tz pgtype.Timestamptz

	_ = tz.Scan(t)

	return tz
}

func timestamptzToTime(tz pgtype.Timestamptz) time.Time {
	if tz.Valid {
		return tz.Time
	}

	return time.Time{}
}
