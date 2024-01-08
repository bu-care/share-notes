package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func main() {
	new_uuid := strings.Replace(uuid.NewString(), "-", "", -1)
	fmt.Println(new_uuid)

	// pwd = str(uuid.uuid4())
	pwd := uuid.NewString()
	fmt.Println(pwd)
}
