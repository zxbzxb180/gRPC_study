package sample

import (
	"github.com/google/uuid"
	"math/rand"
)
import "gRPC_study/pb"

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet(
			"Core i5-100",
			"Core i7-100",
			"Core i9-100",
			"Core i3-100",
		)
	} else if brand == "AMD" {
		return randomStringFromSet(
			"Ryzen 3",
			"Ryzen 5",
			"Ryzen 7",
		)
	} else {
		return ""
	}
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet(
			"GTX 1060",
			"GTX 2060",
			"GTX 3060",
			"GTX 4060",
		)
	} else if brand == "AMD" {
		return randomStringFromSet(
			"RX 390",
			"RX 590",
			"RX 790",
		)
	} else {
		return ""
	}
}

func randomStringFromSet(s ...string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	return s[rand.Intn(l)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1000, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Width:  uint32(width),
		Height: uint32(height),
	}
	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	} else {
		return pb.Screen_OLED
	}
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Lenovo", "Dell", "Apple")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Lenovo":
		return randomStringFromSet(
			"Xiao Xin 10",
			"Xiao Xin 11",
			"Xiao Xin 12",
			"Xiao Xin 13",
		)
	case "Dell":
		return randomStringFromSet(
			"Night elves 7",
			"Night elves 8",
			"Night elves 9",
		)
	case "Apple":
		return randomStringFromSet(
			"Mac 12",
			"Mac 13",
			"Mac 14",
		)
	default:
		return ""
	}
}
