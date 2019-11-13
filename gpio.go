package main

import "os"
import "io/ioutil"

type GPIO struct{}

func (r GPIO) Pin(name string) (GPIO_Pin, error) {
	pin := GPIO_Pin{name}
	filename := pin.Filename()
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// export gpio pin
		if err = ioutil.WriteFile("/sys/class/gpio/export", []byte(pin.Name), 0666); err != nil {
			return pin, err
		}
	} else {
		return pin, err
	}
	return pin, nil
}

type GPIO_Pin struct {
	Name string
}

func (r GPIO_Pin) Filename() string {
	return "/sys/class/gpio/gpio" + r.Name
}
func (r GPIO_Pin) write(where, what string) error {
	filename := r.Filename() + "/" + where
	return ioutil.WriteFile(filename, []byte(what), 0666)
}
func (r GPIO_Pin) Output() error {
	return r.write("direction", "out")
}
func (r GPIO_Pin) Input() error {
	return r.write("direction", "in")
}
func (r GPIO_Pin) High() error {
	return r.write("value", "1")
}
func (r GPIO_Pin) Low() error {
	return r.write("value", "0")
}

