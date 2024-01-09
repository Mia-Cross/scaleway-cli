package object

import (
	"fmt"
	"math/rand"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
)

func randomNameWithPrefix(prefix string) string {
	return fmt.Sprintf("%s-%d", prefix, rand.Int())
}

func createBucket(name string) core.BeforeFunc {
	return core.ExecStoreBeforeCmd("Bucket", fmt.Sprintf("scw object bucket create %s", name))
}

func deleteBucket(name string) core.AfterFunc {
	return core.ExecAfterCmd(fmt.Sprintf("scw object bucket delete %s", name))
}
