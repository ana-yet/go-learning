package main

func main() {
	day := "no more lal day"

	switch day { // tag switch statement
	case "saturday", "sunday":
		println("weekend")

	case "monday", "tuesday", "wednesday", "thursday", "friday":
		println("weekday")

	default:
		println("invalid day")
	}

	// normal switch statement
	switch {
	case day == "saturday" || day == "sunday":
		println("weekend")
	}

}