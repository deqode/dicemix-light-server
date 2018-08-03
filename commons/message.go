package commons

// type Message struct {
// 	Header  header
// 	Payload payload
// }

// type Header struct {
//   string timestamp  = 1;
//   uint32 id  = 2;
//   uint32 state  = 3;
// }

// // const  Payload {
// //   JoinTx joinTx = 1;
// // }

// // message PeerInfo {
// //   uint32 id = 1;
// //   bytes publicKey = 2;
// //   bytes dcVector = 3;
// // }

// // message JoinTx {
// //   uint32 id = 1;
// // }

const (
	C_JOIN_REQUEST     = 1
	C_KEY_EXCHANGE     = 2
	C_EXP_DC_VECTOR    = 3
	C_SIMPLE_DC_VECTOR = 4
	C_TX_SIGNATURE     = 5
	ERROR              = 6
)

const (
	S_JOIN_RESPONSE = 100
	S_KEY_EXCHANGE  = 101
	S_HASHES_VECTOR = 102
	S_MIXED_TX      = 103
	S_JOINED_TX     = 104
)

type JoinTx struct {
	Status  uint32
	ID      int32
	Message string
	Err     string
}
