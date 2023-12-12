package server

import "context"

type IServer interface {
	Start(pctx context.Context)
}
