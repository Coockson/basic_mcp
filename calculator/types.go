package calculator

const Description = "Calculates basic artithmatic math operations such as addition (+) substraction (-) multiplication (*) and division (/)"

type Input struct {
	Operation string  `json:"operator" jsonschema:"the type of the operator such as +, -, *, / for addition, substraction, multiplication and division respectively"`
	A         float64 `json:"A" jsonschema:"first operator"`
	B         float64 `json:"B" jsonschema:"second operator"`
}

type Output struct {
	Result float64 `json:"result" jsonschema:"result of the math equation"`
}
