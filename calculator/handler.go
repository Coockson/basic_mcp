package calculator

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func Calculate(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	operator := input.Operation

	switch operator {
	case "+":
		return nil, Output{Result: input.A + input.B}, nil
	case "-":
		return nil, Output{Result: input.A - input.B}, nil
	case "*":
		return nil, Output{Result: input.A * input.B}, nil
	case "/":
		if input.B == 0 {
			return nil, Output{}, fmt.Errorf("division by zero: cannot divide %v by 0", input.A)
		}
		return nil, Output{Result: input.A / input.B}, nil
	default:
		return nil, Output{}, fmt.Errorf("invalid operator: %q (supported: +, -, *, /)", operator)
	}

}
