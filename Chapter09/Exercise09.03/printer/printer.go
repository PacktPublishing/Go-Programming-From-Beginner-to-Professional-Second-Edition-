package printer

import (
	"fmt"

	"github.com/google/uuid"
)

func PrintNewUUID() string {
	id := uuid.New()
	return fmt.Sprintf("Generated UUID: %s\n", id)
}
