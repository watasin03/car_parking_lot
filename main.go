package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type car_park struct {
	Slot        int    `json:slot`
	Regisnumber string `json:regisnumber`
	Colour      string `json:colour`
}

func main() {
	parkingLot := 0
	var car []car_park
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Parking lot management cli\n")
	for {
		fmt.Print("command: ")
		text, _ := reader.ReadString('\n')
		text = strings.ToLower(strings.Trim(text, " \r\n"))
		textArr := strings.Split(text, " ")
		if strings.Compare(textArr[0], "create_parking_lot") == 0 {
			s, err := strconv.Atoi(textArr[1])
			if err != nil {
				fmt.Printf("Please insert number\n")
			} else {
				parkingLot = CheckSlot(s, parkingLot, car)
				car = CreateParkingLot(parkingLot, car)
			}

		} else if strings.Compare(textArr[0], "park") == 0 {
			if textArr[1] == "" || textArr[2] == "" {
				if textArr[1] == "" {
					println("Please enter Registration number")
				} else {
					println("Please enter Registration color")
				}
			} else {
				car = Parking(car, parkingLot, textArr[2], textArr[1])
			}
		} else if strings.Compare(textArr[0], "leave") == 0 {
			if textArr[1] == "" {
				println("Please enter slot")
			} else {
				s, err := strconv.Atoi(textArr[1])
				if err != nil {
					fmt.Printf("Please insert number\n")
				} else {
					car = Leave(s, car)
				}
			}

		} else if strings.Compare(textArr[0], "status") == 0 {
			Status(car)
		} else if strings.Compare(textArr[0], "slot_number_for_registration_number") == 0 {
			if textArr[1] == "" {
				println("Please enter registration number")
			} else {
				result := CarSlotByRegisnumber(car, textArr[1])
				if result == 0 {
					fmt.Println("Not Found")
				}
			}
		} else if strings.Compare(textArr[0], "slot_numbers_for_cars_with_colour") == 0 {
			if textArr[1] == "" {
				println("Please enter color")
			} else {
				CarRegisSlotByColour(car, textArr[1])
			}
		} else if strings.Compare(textArr[0], "registration_numbers_for_cars_with_colour") == 0 {
			if textArr[1] == "" {
				println("Please enter color")
			} else {
				CarRegisColor(car, textArr[1])
			}
		} else if strings.Compare(textArr[0], "--help") == 0 {
			fmt.Print("create_parking_lot <number> \n")
			fmt.Print("park <string> <string> \n")
			fmt.Print("leave <number> \n")
			fmt.Print("status \n")
			fmt.Print("slot_number_for_registration_number <string> \n")
			fmt.Print("slot_numbers_for_cars_with_colour <string> \n")
			fmt.Print("registration_numbers_for_cars_with_colour <string> \n")

		} else if strings.Compare(textArr[0], "exit") == 0 {
			fmt.Print("Bye\n")
			break
		} else {
			fmt.Print("--help to see all command\n")
		}
	}

}

//create parking lot
func CreateParkingLot(slot int, car []car_park) []car_park {
	var myCarArr []car_park
	if len(car) == 0 {
		myCarArr = make([]car_park, slot)
		for i := 0; i < slot; i++ {
			myCarArr[i] = car_park{
				Slot:        i + 1,
				Colour:      "",
				Regisnumber: "",
			}
		}
		fmt.Printf("Created a parking lot with %d slots\n", slot)
		return myCarArr
	} else if slot-len(car) > 0 {
		newNum := len(car)
		for i := len(car); i < slot; i++ {
			car = append(car, car_park{
				Slot:        newNum + 1,
				Colour:      "",
				Regisnumber: "",
			})
		}
		fmt.Printf("Created a parking lot with %d slots\n", slot)
		return car
	} else {
		return car
	}
}

//CheckSlot
func CheckSlot(new int, slot int, car []car_park) int {
	if slot >= new {
		fmt.Printf("now slot has more your request \n")
		return slot
	} else {
		return new
	}
}

func (f *car_park) GetRegisnumber() string {
	return f.Regisnumber
}

func (f *car_park) GetColor() string {
	return f.Colour
}

func (f *car_park) GetSlotNumber() int {
	return f.Slot
}

//Parking
func Parking(car []car_park, slot int, color string, name string) []car_park {

	for i := range car {
		carArr := car[i]
		regisName := carArr.GetRegisnumber()
		regisColour := carArr.GetColor()
		if regisName == "" {
			car[i].Slot = i + 1
			car[i].Colour = color
			car[i].Regisnumber = name
			fmt.Printf("Allocated slot number: %d \n", i+1)
			return car
		} else if regisName == name && regisColour == color {
			fmt.Printf("Allocated slot already exits\n")
			return car
		} else {
			continue
		}
	}
	fmt.Printf("Sorry, parking lot is full\n")
	return car
}

//
func Leave(slot int, car []car_park) []car_park {

	if len(car) < slot {
		fmt.Printf("Slot number %d is not create\n", slot)
		return car
	} else {
		car[slot-1].Colour = ""
		car[slot-1].Regisnumber = ""
		fmt.Printf("Slot number %d is free\n", slot)
		return car
	}
}

//
func Status(car []car_park) {
	fmt.Println("Slot No.   Registration   No Colour")
	for i := range car {
		slot := car[i].GetSlotNumber()
		name := car[i].GetRegisnumber()
		colour := car[i].GetColor()
		if name == "" {
			continue
		} else {
			fmt.Printf("%d          %s  %s\n", slot, name, colour)
		}
	}
}

//
func CarRegisColor(car []car_park, color string) {
	var text string = ""
	for i := range car {
		if car[i].GetColor() == color {
			text += car[i].GetRegisnumber() + ","
		}
	}
	text = strings.TrimRight(text, ",")
	fmt.Printf("%s\n", text)
}

//
func CarRegisSlotByColour(car []car_park, color string) {
	var text string = ""
	for i := range car {
		if car[i].GetColor() == color {
			text += strconv.Itoa(car[i].GetSlotNumber()) + ","
		}
	}
	text = strings.TrimRight(text, ",")
	fmt.Printf("%s\n", text)
}

//
func CarSlotByRegisnumber(car []car_park, regisnum string) int {
	var num int = 0
	for i := range car {
		if car[i].GetRegisnumber() == regisnum {
			fmt.Printf("%d\n", car[i].GetSlotNumber())
			num = num + 1
			break
		}
	}
	return num
}
