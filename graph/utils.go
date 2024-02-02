package graph

import (
	"context"
	"encoding/base64"
	"math"
	"strconv"

	"github.com/lakshyab1995/tiger-kittens/db"
	"github.com/lakshyab1995/tiger-kittens/graph/model"
	"github.com/lakshyab1995/tiger-kittens/jwt"
)

const earthRadiusKm = 6371.0

func RefreshToken(ctx context.Context, token string) (*model.TokenMeta, error) {
	username, err := jwt.ParseToken(token)
	if err != nil {
		return nil, err
	}
	jwtToken, err := jwt.GenerateToken(username)
	if err != nil {
		return nil, err
	}
	return &model.TokenMeta{
		Token:  jwtToken.Token,
		Expiry: jwtToken.Expiry,
	}, nil
}

func EncodeCursor(id int) string {
	return base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(id)))
}
func DecodeCursor(encoded string) (int, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(string(decoded))
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	dLat := degreesToRadians(lat2 - lat1)
	dLon := degreesToRadians(lon2 - lon1)

	lat1 = degreesToRadians(lat1)
	lat2 = degreesToRadians(lat2)

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Sin(dLon/2)*math.Sin(dLon/2)*math.Cos(lat1)*math.Cos(lat2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadiusKm * c
}

func isWithinRange(newSighting model.CoordinatesInput, lastSighting db.Coordinates) bool {
	// Calculate the distance between the last sighting and the new sighting
	distance := haversine(lastSighting.Lat, lastSighting.Lon, newSighting.Lat, newSighting.Lon)

	// Return true if the distance is less than or equal to 5 kilometers
	return distance <= 5.0
}
