package nativelib

import . "github.com/NjinN/RML-view/core"

func InitView(ctx *BindMap) {
	var viewToken = Token{
		NATIVE,
		Native{
			"view",
			2,
			View,
			nil,
		},
	}
	ctx.PutNow("view", &viewToken)

}
