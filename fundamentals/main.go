package main

import (
	"fmt"
	lesson_one "fundamentals/1.exporting_variables"
	lesson_ten "fundamentals/10.format_txt_output"
	lesson_eleven "fundamentals/11.loops"
	lesson_twelve "fundamentals/12.using_range"
	lesson_thirteen "fundamentals/13.more_on_arrays"
	lesson_fourteen "fundamentals/14.functions"
	lesson_fifteen "fundamentals/15.file_io"
	lesson_sixteen "fundamentals/16.command_line"
	lesson_seventeen "fundamentals/17.packages_and_modules"
	lesson_eighteen "fundamentals/18.maps"
	lesson_nineteen "fundamentals/19.generics"
	lesson_two "fundamentals/2.exporting_variables_example_two"
	lesson_twenty "fundamentals/20.structs"
	lesson_twenty_one "fundamentals/21.defined_types"
	lesson_twenty_two "fundamentals/22.setters_and_getters"
	lesson_twenty_three "fundamentals/23.interfaces"
	lesson_twenty_four "fundamentals/24.concurrency"
	lesson_twenty_five "fundamentals/25.bank_acct_example"
	lesson_twenty_six "fundamentals/26.closures"
	lesson_twenty_seven "fundamentals/27.recursion"
	lesson_twenty_eight "fundamentals/28.regular_expression"
	lesson_twenty_nine "fundamentals/29.automated_testing"
	lesson_three "fundamentals/3.exported_functions"
	lesson_four "fundamentals/4.variables"
	lesson_five "fundamentals/5.operators"
	lesson_six "fundamentals/6.more_string_operations"
	lesson_seven "fundamentals/7.working_with_runes"
	lesson_eight "fundamentals/8.working_with_time"
	lesson_nine "fundamentals/9.math"
)

var pl = fmt.Println // print line short hand
var logicalBreak = "================================================================================"

// to set the namespace for the project, use: go mod init "package-name"
// see go.mod to remember what it means to set the root namespace...

func lessonZero() {
	pl("Lesson Zero Started.")
	pl("Hello World")
	pl("Lesson Zero Complete!")
}

func lessonOne() {
	pl("Lesson One Started.")
	pl("Hello person from lesson_one:", lesson_one.Name)
	pl("Lesson One Complete!")
}

func lessonTwo() {
	pl("Lesson Two Started.")
	pl("Article from lesson_two:", lesson_two.Article)
	pl("Lesson Two Complete!")
}

func lessonThree() {
	pl("Lesson Three Started.")
	pl("Using a getter to get a private greeting variable:", lesson_three.GetGreeting())
	lesson_three.SetGreeting("This is a new greeting!")
	pl("An updated greeting:", lesson_three.GetGreeting())
	pl("Lesson Three Complete!")
}

func main() {
	lessonZero() // namespace is important...
	pl(logicalBreak)
	lessonOne() // don't place comments after the name of the package...
	pl(logicalBreak)
	lessonTwo() // use of a variable or function can be done via importing...
	pl(logicalBreak)
	lessonThree() // getters and setters over use within the same file for readability...
	pl(logicalBreak)
	pl(lesson_four.Start()) // variables
	pl(logicalBreak)
	pl(lesson_five.Start()) // operators
	pl(logicalBreak)
	pl(lesson_six.Start()) // more string operations
	pl(logicalBreak)
	pl(lesson_seven.Start()) // working with runes
	pl(logicalBreak)
	pl(lesson_eight.Start()) // working with time
	pl(logicalBreak)
	pl(lesson_nine.Start()) // math
	pl(logicalBreak)
	pl(lesson_ten.Start()) // formatting text/string output
	pl(logicalBreak)
	pl(lesson_eleven.Start()) // loops
	pl(logicalBreak)
	pl(lesson_twelve.Start()) // using range
	pl(logicalBreak)
	pl(lesson_thirteen.Start()) // more on arrays
	pl(logicalBreak)
	pl(lesson_fourteen.Start()) // functions
	pl(logicalBreak)
	pl(lesson_fifteen.Start()) // file i/o
	pl(logicalBreak)
	pl(lesson_sixteen.Start()) // command line
	pl(logicalBreak)
	pl(lesson_seventeen.Start()) // packages/modules
	pl(logicalBreak)
	pl(lesson_eighteen.Start()) // maps
	pl(logicalBreak)
	pl(lesson_nineteen.Start()) // generics
	pl(logicalBreak)
	pl(lesson_twenty.Start()) // structs
	pl(logicalBreak)
	pl(lesson_twenty_one.Start()) // defined types
	pl(logicalBreak)
	pl(lesson_twenty_two.Start()) // setters/getters
	pl(logicalBreak)
	pl(lesson_twenty_three.Start()) // interfaces
	pl(logicalBreak)
	pl(lesson_twenty_four.Start()) // concurrency
	pl(logicalBreak)
	pl(lesson_twenty_five.Start()) // bank account example
	pl(logicalBreak)
	pl(lesson_twenty_six.Start()) // closures
	pl(logicalBreak)
	pl(lesson_twenty_seven.Start()) // recursion
	pl(logicalBreak)
	pl(lesson_twenty_eight.Start()) // regular expressions
	pl(logicalBreak)
	pl(lesson_twenty_nine.Start()) // automated testing
}
