package main

import "fmt"

// CPU is a struct that represents a CPU
type CPU struct {
	// Model represents the model of the CPU
	Model string

	// Cores represents the number of cores of the CPU
	Cores int

	// Threads represents the number of threads of each core of the CPU
	Threads int

	// BaseClock represents the base clock speed of the CPU
	BaseClock float64

	// MaxClock represents the max clock speed of the CPU
	MaxClock float64
}

// PerformanceScore returns the performance score of the CPU
func (c CPU) PerformanceScore() (p float64) {
	avgClockSpeed := (c.BaseClock + c.MaxClock) / 2
	p = float64(c.Cores * c.Threads) * avgClockSpeed
	return
}

// Laptop is a struct that represents a laptop
type Laptop struct {
	// Model represents the model of the laptop
	Model string
	// Manufacturer represents the manufacturer of the laptop
	Manufacturer string
	// CPU represents the CPU of the laptop (embedded struct)
	CPU
}

// ShowInfo returns the info of the laptop
func (l Laptop) ShowInfo() (info string) {
	info = fmt.Sprintf(
		"Model: %s\nManufacturer: %s\nCPU Performance Score: %.2f\n",
		l.Model,
		l.Manufacturer,
		l.PerformanceScore(),
	)

	return
}

// PC is a struct that represents a PC
type PC struct {
	// Height represents the height of the PC
	Height int

	// Width represents the width of the PC
	Width int
	
	// Manufacturer represents the manufacturer of the PC
	Manufacturer string

	// CPU represents the CPU of the PC (embedded struct)
	CPU
}

// ShowInfo returns the info of the PC
func (p PC) ShowInfo() (info string) {
	info = fmt.Sprintf(
		"Height: %d\nWidth: %d\nManufacturer: %s\nCPU Performance Score: %.2f\n",
		p.Height,
		p.Width,
		p.Manufacturer,
		p.PerformanceScore(),
	)

	return
}

func main() {
	laptop := Laptop{
		Model:        "MacBook Pro 16",
		Manufacturer: "Apple",
		CPU: CPU{
			Model:     "Intel Core i9-9980HK",
			Cores:     8,
			Threads:   16,
			BaseClock: 2.4,
			MaxClock:  5.0,
		},
	}

	pc := PC{
		Height:       42,
		Width:        20,
		Manufacturer: "Dell",
		CPU: CPU{
			Model:     "Intel Core i9-9900K",
			Cores:     8,
			Threads:   16,
			BaseClock: 3.6,
			MaxClock:  5.0,
		},
	}

	fmt.Println(laptop.ShowInfo())
	fmt.Println(pc.ShowInfo())

	laptopPerformance := laptop.PerformanceScore()
	pcPerformance := pc.PerformanceScore()

	if laptopPerformance == pcPerformance {
		fmt.Println("The laptop and the PC have the same performance score")
	} else if laptopPerformance > pcPerformance {
		fmt.Println("The laptop has a better performance score than the PC")
	} else {
		fmt.Println("The PC has a better performance score than the laptop")
	}
}
	
	