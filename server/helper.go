package server

import (
	"../commons"
)

func registerDelayHandler(h *Hub, code uint32, message string) {
	switch code {
	case commons.C_KEY_EXCHANGE:
		// TODO: if function not called yet then call it
	}
}
