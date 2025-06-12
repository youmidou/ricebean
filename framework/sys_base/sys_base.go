package sys_base

import (
	"fmt"
	"strings"
)

// Nid
func GetGCidAndCCid(Cid string) (string, string, bool) {
	ks := strings.Split(Cid, "->")
	if len(ks) != 2 {
		return "", "", false
	}
	GCid := ks[0] //网关Id
	CCid := ks[1] //客户端id
	return GCid, CCid, true
}
func GetCid(GCid string, PCid string) string {
	return fmt.Sprintf("%s->%s", GCid, PCid)
}
