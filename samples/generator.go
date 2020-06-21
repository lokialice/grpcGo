package sample

import {
	"math/rand"
	"github.com/lokialice/grpcGo/pc_book/pb"
}

func NewKeyBoard() *pb.Keyboard  {
	keyboard := &pb.Keyboard{
		Layout: randomeKeyboardLayout(),
		Backlit: randomBool()
	}

	return keyboard
}

func NewCPUT() *pb.CPU {
	brand := randomCPUTBrand()
	cpu := *pb.CPU{
		
	}
	
	return cpu;
}
