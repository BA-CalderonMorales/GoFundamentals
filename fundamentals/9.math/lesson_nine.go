package lesson_nine

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var log = fmt.Println
var logF = fmt.Printf

func Start() string {
	log("Lesson Nine Started.")
	log("Math")

	basicArithmetic()
	shortHandOps()
	precision()
	randomValues()
	mathLibraryUsefulFunctions()
	conversions()
	trigFunctions()
	return "Lesson Nine Complete!"
}

func basicArithmetic() {
	log("5 + 4=", 5+4)
	log("5 - 4=", 5-4)
	log("5 * 4=", 5*4)
	log("5 / 4=", 5/4)
	log("5 % 4=", 5%4)
}

func shortHandOps() {
	mInt := 1
	mInt += 1
	log(mInt) // can use -= /= *= %=...
}

func precision() {
	log("Float Precision =", 0.1111111111111+0.111111111111111)
}

func randomValues() {
	// seconds since the date 1/1/1970
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1 // generate rand values from 0-49
	log("Random :", randNum)
}

func mathLibraryUsefulFunctions() {
	log("Abs(-10) =", math.Abs(-10))
	log("Pow(4, 2) =", math.Pow(4, 2))
	log("Sqrt(16) =", math.Sqrt(16))
	log("Cbrt(8) =", math.Cbrt(8))
	log("Ceil(4.4) =", math.Ceil(4.4))
	log("Floor(4.4) =", math.Floor(4.4))
	log("Round(4.4) =", math.Round(4.4))
	log("Log2(8) =", math.Log2(8))
	log("Log10(100 =", math.Log10(100))
	// Get the log of e to the power of 2
	log("Log(7.389) =", math.Log(math.Exp(2)))
	log("Max(5, 4) =", math.Max(5, 4))
	log("Min(5, 4) =", math.Min(5, 4))
}

func conversions() {
	r90 := 90 * math.Pi / 180 // Convert 90 degrees to radians
	log(r90)
	d90 := r90 * (180 / math.Pi)
	log(d90)
	logF("%.2f radians = %.2f degrees\n", r90, d90)
}

func trigFunctions() {
	r90 := 90 * math.Pi / 180 // Convert 90 degrees to radians
	log("Sin(90) =", math.Sin(r90))
	// There are also functions for Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
	// Htpot (hypotenuse)
}
