// This file is part of DiceDB.
// Copyright (C) 2024 DiceDB (dicedb.io).
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package geo

import (
	"math"

	"github.com/dicedb/dice/internal/errors"
	"github.com/mmcloughlin/geohash"
)

// Earth's radius in meters
const earthRadius float64 = 6372797.560856

// Bit precision for geohash - picked up to match redis
const bitPrecision = 52

// Bit precision for geohash string - picked up to match redis
const bitPrecisionString = 10

func DegToRad(deg float64) float64 {
	return math.Pi * deg / 180.0
}

func RadToDeg(rad float64) float64 {
	return 180.0 * rad / math.Pi
}

func GetDistance(
	lon1,
	lat1,
	lon2,
	lat2 float64,
) float64 {
	lon1r := DegToRad(lon1)
	lon2r := DegToRad(lon2)
	v := math.Sin((lon2r - lon1r) / 2)
	// if v == 0 we can avoid doing expensive math when lons are practically the same
	if v == 0.0 {
		return GetLatDistance(lat1, lat2)
	}

	lat1r := DegToRad(lat1)
	lat2r := DegToRad(lat2)
	u := math.Sin((lat2r - lat1r) / 2)

	a := u*u + math.Cos(lat1r)*math.Cos(lat2r)*v*v

	return 2.0 * earthRadius * math.Asin(math.Sqrt(a))
}

func GetLatDistance(lat1, lat2 float64) float64 {
	return earthRadius * math.Abs(DegToRad(lat2)-DegToRad(lat1))
}

// EncodeInt returns a geo hash for a given coordinate, and returns it in float64 so it can be used as score in a zset
func EncodeInt(lat, lon float64) float64 {
	h := geohash.EncodeIntWithPrecision(lat, lon, bitPrecision)

	return float64(h)
}

// DecodeInt returns the latitude and longitude from a geo hash
// The hash should be a float64, as it is used as score in a zset
func DecodeInt(hash float64) (lat, lon float64) {
	lat, lon = geohash.DecodeIntWithPrecision(uint64(hash), bitPrecision)

	return lat, lon
}

func EncodeString(lat, lon float64) string {
	return geohash.EncodeWithPrecision(lat, lon, bitPrecisionString)
}

// ConvertDistance converts a distance from meters to the desired unit
func ConvertDistance(
	distance float64,
	unit string,
) (converted float64, err []byte) {
	switch unit {
	case "m":
		return distance, nil
	case "km":
		return distance / 1000, nil
	case "mi":
		return distance / 1609.34, nil
	case "ft":
		return distance / 0.3048, nil
	default:
		return 0, errors.NewErrWithMessage("ERR unsupported unit provided. please use m, km, ft, mi")
	}
}