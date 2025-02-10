package simplerat

import "time"

type Device struct {
	Id              int64      `json:"id" db:"id,omitempty"`
	Name            string     `json:"name" db:"name"`
	RegisterToken   string     `json:"register_token" db:"register_token"`
	Dtype           string     `json:"dtype" db:"dtype"`
	Status          string     `json:"status" db:"status"`
	LastIp          string     `json:"last_ip" db:"last_ip"`
	IsBeaconEnabled int        `json:"beacon_enabled" db:"beacon_enabled"`
	LastActive      *time.Time `json:"last_active" db:"last_active,omitempty"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at,omitempty"`
}

type DeviceBeacon struct {
	Id         int64      `json:"id" db:"id,omitempty"`
	Title      string     `json:"title" db:"title"`
	DeviceId   int64      `json:"device_id" db:"device_id"`
	Lat        float32    `json:"lat" db:"lat"`
	Lng        float32    `json:"lng" db:"lng"`
	BeaconType string     `json:"beacon_type" db:"beacon_type"`
	BeaconData string     `json:"beacon_data" db:"beacon_data"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at"`
}
