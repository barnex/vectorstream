package main

import (
	"math"
	"path"
)

type Model2 struct {
	w        [10]Mat
	b        [10]float64
	training [10][]Mat
}

var digits = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

func (m *Model2) Train(dir string) {
	m.training = loadAllDigits(path.Join(dir, "training"))

	for i := range digits {
		w := &m.w[i]
		*w = NewMat(28, 28)
		for _, img := range m.training[i] {
			Add(m.w[i].List, img.List)
		}
		stdnorm(m.w[i].List, m.w[i].List)
	}

	//for _, w := range m.w {
	//	w.Print()
	//}
}

func (m *Model2) Infer(img Mat) []float32 {
	prob := make([]float32, 10)
	for i, w := range m.w {
		prob[i] = Dot(img.List, w.List)
	}
	//SoftMax(prob, prob)
	return prob
}

func (m *Model2) Loss() float64 {
	loss := 0.0
	for i := range m.training {
		var real [10]float32
		real[i] = 1
		for _, img := range m.training[i] {
			pred := m.Infer(img)
			loss += float64(XEntropy(real[:], pred))
		}
	}
	return loss
}

func findMax(x []float32) int {
	maxX := float32(math.Inf(-1))
	maxI := 0
	for i, x := range x {
		if x > maxX {
			maxI = i
			maxX = x
		}
	}
	return maxI
}

//func (m *Model2) load(dir string) {
//	log.Println("loading", dir)
//	m.training = loadAllDigits(path.Join(dir, "training"))
//	log.Println("done")
//	for i := range m.w {
//		m.w[i] = NewMat(28, 28)
//		w := m.w[i]
//		for _, img := range m.training[i] {
//			Add(w.List, img.List)
//		}
//		stdnorm(w.List, w.List)
//	}
//}
