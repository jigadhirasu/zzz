package zdi

import (
	"github.com/jigadhirasu/zzz/z"
	"github.com/kubemq-io/kubemq-go"
)

type KubemqDI func(client *kubemq.Client) z.Bytes
