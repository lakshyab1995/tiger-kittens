package graph

import "fmt"

type WrongUsernameOrPasswordError struct{}

func (m *WrongUsernameOrPasswordError) Error() string {
	return "wrong username or password"
}

type TigerWithinRangeError struct {
	TigerID int
}

func (e *TigerWithinRangeError) Error() string {
	return fmt.Sprintf("Tiger with ID %d is within 5 kilometers of its previous sighting", e.TigerID)
}
