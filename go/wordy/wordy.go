package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	//Check input string for valid format. If not valid => 0, false
	if mathc, _ := regexp.MatchString(`What is -?\d+(?: (?:plus|minus|divided by|multiplied by) -?\d+)*\?`, question); !mathc {
		return 0, false
	}

	//Regexp for math operators and add it in operators[]
	operationReg := regexp.MustCompile(`(plus|minus|divided|multiplied)`)
	operators := operationReg.FindAllString(question, -1)

	//Regexp for numbers and add it in numbers[]
	numberReg := regexp.MustCompile(`-?\d+`)
	numbers := numberReg.FindAllString(question, -1)

	//Check for valid format parse numbers and operators
	if len(operators) != len(numbers)-1 {
		return 0, false
	}

	summ, _ := strconv.Atoi(numbers[0])

	//Loop for operators and calculate result with correct math operator
	for idx, opr := range operators {
		nmb, _ := strconv.Atoi(numbers[idx+1])
		switch opr {
		case "plus":
			summ += nmb
		case "minus":
			summ -= nmb
		case "divided":
			summ /= nmb
		case "multiplied":
			summ *= nmb
		}
	}

	return summ, true
}
