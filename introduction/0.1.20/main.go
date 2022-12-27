package main

import (
	"encoding/json"
	"fmt"
	"github.com/hacker65536/0.1.20/pkg/vault"
	"os"
)

func main() {
	s := vault.NewSecret()

	err := json.NewEncoder(os.Stdout).Encode(s)
	if err != nil {
		fmt.Println("failed to json encode, error =", err)
	}
	// tokenフィールドは、JSONに含まれていないことに注目
	// Output: {"ID":"dummy_id","CreateTime":"2009-11-10T23:00:00Z"}
}
