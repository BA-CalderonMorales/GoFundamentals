package lesson_three

var greeting = "This was suppose to be hidden!"

func GetGreeting() string {
	return greeting
}

func SetGreeting(newGreeting string) {
	greeting = newGreeting
}
