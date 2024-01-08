package contextlearn

import (
	"context"
	"fmt"
	"testing"
)

func TestCtx(t *testing.T) {
	ctx := context.Background()
	fmt.Println(ctx)
}
