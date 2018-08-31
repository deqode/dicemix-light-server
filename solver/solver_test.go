package solver

import (
	"testing"

	"github.com/manjeet-thadani/dicemix-server/utils"
)

type testpair struct {
	data []uint64
	res  []uint64
}

var solverTests = []testpair{
	{
		[]uint64{1859546079985200847, 1646884441642370562, 1945157946220288822, 2071666930927106951, 1683255082316998317},
		[]uint64{338987782431557515, 760646884788788847, 805715802280412061, 855681932209541597, 1404356687488594778},
	},

	{
		[]uint64{775687816567226590, 1737237369846472686, 370018584425987655, 1458106296827716655, 1451264941915419093, 1707673020921186118},
		[]uint64{574560148427821549, 808691242602581210, 1062313112672455589, 1465054401282329629, 1477299493679429209, 2305298445543691257},
	},

	{
		[]uint64{1890557469049449562, 1784695988170892157, 2043447899551859310, 1647265920402198959, 1817920860757799870},
		[]uint64{106674738265709861, 1441016479585612707, 1516345228348743180, 1674760640954454528, 1763446400322317188},
	},
}

func TestFlintSolver(t *testing.T) {
	for _, pair := range solverTests {
		output := Solve(pair.data, len(pair.data))

		if !utils.CheckEqualUint64(pair.res, output) {
			t.Error(
				"For", pair.data,
				"expected", pair.res,
				"got", output,
			)
		}
	}
}
