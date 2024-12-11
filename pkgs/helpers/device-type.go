package helpers

import "github.com/avct/uasurfer"

func GetDeviceType(userAgent string) string {
	ua := uasurfer.Parse(userAgent)

	var device string

	switch deviceType := ua.DeviceType; deviceType {
	case uasurfer.DeviceComputer:
		device = "DESKTOP"
	case uasurfer.DeviceTablet:
		device = "TABLET"
	case uasurfer.DevicePhone:
		device = "PHONE"
	default:
		device = "MISSING"
	}

	return device
}
