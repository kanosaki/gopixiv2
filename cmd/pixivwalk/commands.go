package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/k0kubun/pp"
)

type command struct {
	fn func(*app, []string) error
}

func cmdGet(a *app, blocks []string) error {
	if len(blocks) == 0 {
		return nil
	}
	param := map[string]string{}
	for _, block := range blocks[1:] {
		kv := strings.SplitN(block, "=", 2)
		if len(kv) != 2 {
			return fmt.Errorf("invalid parameter %s", block)
		}
		param[kv[0]] = kv[1]
	}
	d, err := a.api.DoGet(context.TODO(), "/"+strings.Replace(blocks[0], ".", "/", -1), param)
	if err != nil {
		return err
	}
	pp.Println(d)
	return nil
}
