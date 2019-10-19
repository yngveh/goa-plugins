// +build debug

package logrus

import (
	"fmt"
)

func debugf(s string, p ...interface{}) {
	fmt.Printf("[Plugin] "+s+"\n", p)
}
