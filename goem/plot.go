package goem

import (
	"image/color"
	"strconv"

	"github.com/gonum/plot"
	"github.com/gonum/plot/plotter"
	"github.com/gonum/plot/vg"
)

func (em EM) Plot(fileid int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "EM Algorithm Plot"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	bs, err := plotter.NewBubbles(em.clusterTriples(), vg.Points(30), vg.Points(70))
	if err != nil {
		panic(err)
	}
	bs.Color = color.RGBA{R: 255, B: 255, A: 255}
	p.Add(bs)

	ss, err := plotter.NewScatter(em.dataTriples())
	if err != nil {
		panic(err)
	}
	ss.Color = color.Black
	p.Add(ss)

	filename := "pic/" + strconv.Itoa(fileid) + ".png"
	if err := p.Save(10*vg.Inch, 10*vg.Inch, filename); err != nil {
		panic(err)
	}
}

func (em EM) dataTriples() plotter.XYs {
	data := make(plotter.XYs, len(em.data))
	for i, v := range em.data {
		data[i].X = v[0]
		data[i].Y = v[1]
	}
	return data
}

func (em EM) clusterTriples() plotter.XYZs {
	data := make(plotter.XYZs, em.k)
	for i, v := range em.mu {
		data[i].X = v[0]
		data[i].Y = v[1]
		data[i].Z = em.pi[i] * 10
	}

	return data
}
