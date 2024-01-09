package object

import (
	"fmt"
	"testing"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
)

func Test_BucketCreate(t *testing.T) {
	bucketName := randomNameWithPrefix("cli-test-bucket-create")

	t.Run("Simple", core.Test(&core.TestConfig{
		Commands: GetCommands(),
		//UseE2EClient:    false,
		Cmd:       fmt.Sprintf("scw object bucket create %s", bucketName),
		Check:     core.TestCheckGolden(),
		AfterFunc: deleteBucket(bucketName),
	}))
}

func Test_BucketDelete(t *testing.T) {
	bucketName := randomNameWithPrefix("cli-test-bucket-delete")

	t.Run("Simple", core.Test(&core.TestConfig{
		Commands:   GetCommands(),
		BeforeFunc: createBucket(bucketName),
		Cmd:        fmt.Sprintf("scw object bucket delete %s", bucketName),
		Check:      core.TestCheckGolden(),
	}))
}
