package main

import "fmt"

type Bus struct {
	ram [4096]byte
	cpu CPU
}

type DimensionError struct {
	Details string
}

func (d *DimensionError) Error() string {
	return fmt.Sprintf("dimension error encountered: %s", d.Details)

}

func (bus *Bus) Write(addr int, value byte) error {
	if addr < 0 || addr >= 4096 {
		return &DimensionError{Details: "expected address to be between 0 and 4096"}
	}
	bus.ram[addr] = value
	return nil
}

func (bus *Bus) Read(addr int) (byte, error) {
	if addr < 0 || addr >= 4096 {
		return 0, &DimensionError{Details: "expected address to be between 0 and 4096"}
	}
	return bus.ram[addr], nil
}
