package sample

import (
	"gRPC_study/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}

	return keyboard
}

func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat64(2.0, 3.5)
	maxGhz := randomFloat64(minGhz, 5.0)

	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}

	return cpu
}

func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat64(1.0, 1.5)
	maxGhz := randomFloat64(minGhz, 2.0)
	memory := &pb.Memory{
		Value: uint64(randomInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}
	cpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}

	return cpu
}

func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randomInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

func NewSSD() *pb.Storage {
	memory := &pb.Memory{
		Value: uint64(randomInt(128, 1024)),
		Unit:  pb.Memory_GIGABYTE,
	}
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: memory,
	}
	return ssd
}

func NewHHD() *pb.Storage {
	memory := &pb.Memory{
		Value: uint64(randomInt(1, 6)),
		Unit:  pb.Memory_TERABYTE,
	}
	hhd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: memory,
	}
	return hhd
}

func NewScreen() *pb.Screen {
	resolution := randomScreenResolution()
	panel := randomScreenPanel()
	screen := &pb.Screen{
		SizeInch:   float32(randomFloat64(13, 17)),
		Resolution: resolution,
		Panel:      panel,
		Multitouch: randomBool(),
	}

	return screen
}

func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:          randomID(),
		Brand:       brand,
		Name:        name,
		Cpu:         NewCPU(),
		Memory:      NewRAM(),
		Gpus:        []*pb.GPU{NewGPU(), NewGPU()},
		Storage:     []*pb.Storage{NewHHD(), NewSSD()},
		Screen:      NewScreen(),
		Keyboard:    NewKeyboard(),
		Weight:      &pb.Laptop_WeightKg{WeightKg: randomFloat64(1.0, 3.0)},
		PriceUsd:    randomFloat64(3500, 6500),
		ReleaseYear: uint32(randomInt(2015, 2023)),
		UpdatedAt:   timestamppb.Now(),
	}
	return laptop
}
