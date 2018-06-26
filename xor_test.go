package vs

import (
	"testing"

	"github.com/barnex/vectorstream/test"
)

// Test feedforward of an XOR network.
// deeplearningbook.org p. 171, adapted to use softmax.
func TestXOR(t *testing.T) {
	l2 := LU(2, 2)
	l1 := NewActivation(2, Re)
	l0 := LU(2, 2)
	net := NewNet(l2, l1, l0)

	// pre-trained weights from deeplearningbook.org p.171
	copyv(l0.Weights(net.wl[0]).List, V{1, 1, 1, 1})
	copyv(l0.Biases(net.wl[0]), V{0, -1})
	copyv(l2.Weights(net.wl[2]).List, V{1, -2, -1, 2})
	copyv(l2.Biases(net.wl[2]), V{0, 1})

	cases := []struct {
		x, want V
	}{
		{V{0, 0}, V{0, 1}},
		{V{0, 1}, V{1, 0}},
		{V{1, 0}, V{1, 0}},
		{V{1, 1}, V{0, 1}},
	}

	y := MakeV(net.NumOut())
	for _, c := range cases {
		net.Eval(y, c.x)
		want := MakeV(c.want.Len())
		softmax(want, c.want)
		test.Eqv(t, y, want)
	}

	testGrad(t, net)
}