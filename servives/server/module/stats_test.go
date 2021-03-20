package module

import (
	"fmt"
	"testing"
	"time"
)

func TestPopular(t *testing.T) {
	c := time.Now().Unix()
	fmt.Println(GetStats(c-0))
	fmt.Println(GetStats(c-1))
	fmt.Println(GetStats(c-61))
	fmt.Println(GetStats(c-(60*60+1)))
	fmt.Println(GetStats(c-(60*60*24+1)))
}