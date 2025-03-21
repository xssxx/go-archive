package gun

import (
	"fmt"
)

func GetGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newMusket(), nil
	}

	return nil, fmt.Errorf("wrong gun type passed")
}
