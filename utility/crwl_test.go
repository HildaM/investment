package utility

import (
	"testing"
)

func TestCrwl_Run(t *testing.T) {
	t.Log(NewDefaultCrwl().Run("https://www.cls.cn"))
}
