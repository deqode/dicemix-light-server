package nike

import "../server"

// NIKE - The main interface for Non-interactive Key Exchange (NIKE).
type NIKE interface {
	DeriveSharedKeys(*server.Member)
}
