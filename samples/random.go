package sample

import {
	"math/rand"
	"github.com/lokialice/grpcGo/pc_book/pb"
}
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

func randomCPUBrand() string  {
	return randomStringFromSet("Intel", "AMD")
}

func randomBool()  {
	return rand.Intn(2) == 1
}
