package frum

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// TimestampToTime ...
func TimestampToTime(ts *timestamp.Timestamp) (time.Time, error) {
	return ptypes.Timestamp(ts)
}

// TripTimestampToTime ...
func TripTimestampToTime(ts *timestamp.Timestamp, err error) (time.Time, error) {
	if err != nil {
		return time.Time{}, err
	}

	return TimestampToTime(ts)
}

// TimestampToNullTime ...
func TimestampToNullTime(ts *timestamp.Timestamp) (mysql.NullTime, error) {
	t, err := TimestampToTime(ts)
	nt := mysql.NullTime{Time: t}
	if err != nil {
		return nt, err
	}

	if !t.IsZero() {
		nt.Valid = true
	}

	return nt, nil
}

// TripTimestampToNullTime ...
func TripTimestampToNullTime(ts *timestamp.Timestamp, err error) (mysql.NullTime, error) {
	if err != nil {
		return mysql.NullTime{Time: time.Time{}}, err
	}

	return TimestampToNullTime(ts)
}

// TimeToTimestamp ...
func TimeToTimestamp(t time.Time) (*timestamp.Timestamp, error) {
	return ptypes.TimestampProto(t)
}

// TripTimeToTimestamp ...
func TripTimeToTimestamp(t time.Time, err error) (*timestamp.Timestamp, error) {
	if err != nil {
		return nil, err
	}

	return TimeToTimestamp(t)
}

// NullTimeToTimestamp ...
func NullTimeToTimestamp(nt mysql.NullTime) (*timestamp.Timestamp, error) {
	if !nt.Valid {
		return nil, nil
	}

	return TimeToTimestamp(nt.Time)
}

// TripNullTimeToTimestamp ...
func TripNullTimeToTimestamp(nt mysql.NullTime, err error) (*timestamp.Timestamp, error) {
	if err != nil {
		return nil, err
	}

	return NullTimeToTimestamp(nt)
}
