package models

import (
	"strconv"
	"strings"
)

type Version struct {
	Version string `json:"version" validate:"required"`
}

func (v *Version) Unwrap() UnwrappedVersion {
	if v != nil {
		arr := strings.Split(v.Version, ".")

		// Lint checks does not like when errors are overwritten. Skip Lint on following rows:
		major, err := strconv.Atoi(arr[0])                                 //nolint:all
		middle, err := strconv.Atoi(arr[1])                                //nolint:all
		minor, err := strconv.Atoi(arr[2])                                 //nolint:all
		full, err := strconv.Atoi(strings.Replace(v.Version, ".", "", -1)) //nolint:all

		if err != nil {
			return UnwrappedVersion{}
		}

		return UnwrappedVersion{
			Major:    major,
			Middle:   middle,
			Minor:    minor,
			AsString: v.Version,
			AsInt:    full,
		}
	}

	return UnwrappedVersion{}
}

// Non-standard object to make compare easier
type UnwrappedVersion struct {
	Major  int
	Middle int
	Minor  int

	AsString string // full version as string
	AsInt    int    // full version as int
}
