package tools

import "hello/local"

var (
	Lang map[string]map[string]string
)

func init(){
	Lang = local.En
}
