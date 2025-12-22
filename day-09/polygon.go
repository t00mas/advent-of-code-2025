package main

import "sort"

type Point = [2]int

type vertEdge struct {
	x, y1, y2 int
}

type band struct {
	yStart, yEnd int
	minX, maxX   int
}

type Polygon struct {
	points    []Point
	vertEdges []vertEdge
	bands     []band
	horizYs   map[int]bool
}

func NewPolygon(points []Point) *Polygon {
	p := &Polygon{
		points:  points,
		horizYs: make(map[int]bool),
	}
	p.buildEdges()
	p.buildBands()
	return p
}

func (p *Polygon) buildEdges() {
	for i := range p.points {
		p1 := p.points[i]
		p2 := p.points[(i+1)%len(p.points)]

		if p1[0] == p2[0] {
			p.vertEdges = append(p.vertEdges, vertEdge{
				x:  p1[0],
				y1: min(p1[1], p2[1]),
				y2: max(p1[1], p2[1]),
			})
		} else {
			p.horizYs[p1[1]] = true
		}
	}
}

func (p *Polygon) buildBands() {
	// bounds
	var ys []int
	for y := range p.horizYs {
		ys = append(ys, y)
	}
	sort.Ints(ys)

	if len(ys) == 0 {
		return
	}

	// bands
	for i := 0; i < len(ys); i++ {
		y := ys[i]
		minX, maxX, valid := p.computeRowBounds(y)
		if valid {
			p.bands = append(p.bands, band{y, y, minX, maxX})
		}

		// band (interior)
		if i+1 < len(ys) && ys[i+1] > y+1 {
			interiorY := y + 1
			minX, maxX, valid := p.computeRowBounds(interiorY)
			if valid {
				p.bands = append(p.bands, band{y + 1, ys[i+1] - 1, minX, maxX})
			}
		}
	}

	sort.Slice(p.bands, func(i, j int) bool {
		return p.bands[i].yStart < p.bands[j].yStart
	})
}

func (p *Polygon) computeRowBounds(y int) (int, int, bool) {
	var xs []int
	for _, e := range p.vertEdges {
		if e.y1 < y && y <= e.y2 {
			xs = append(xs, e.x)
		}
	}

	rangeMin, rangeMax := -1, -1

	if len(xs) >= 2 {
		sort.Ints(xs)
		rangeMin, rangeMax = xs[0], xs[len(xs)-1]
	}

	for i := range p.points {
		p1 := p.points[i]
		p2 := p.points[(i+1)%len(p.points)]
		if p1[1] == y && p2[1] == y {
			edgeMin, edgeMax := min(p1[0], p2[0]), max(p1[0], p2[0])
			if rangeMin == -1 {
				rangeMin, rangeMax = edgeMin, edgeMax
			} else {
				rangeMin = min(rangeMin, edgeMin)
				rangeMax = max(rangeMax, edgeMax)
			}
		}
	}

	if rangeMin == -1 {
		return 0, 0, false
	}
	return rangeMin, rangeMax, true
}

func (p *Polygon) ContainsRect(x1, y1, x2, y2 int) bool {
	// binary search
	lo := sort.Search(len(p.bands), func(i int) bool {
		return p.bands[i].yEnd >= y1
	})

	for i := range p.bands {
		b := p.bands[i]
		if b.yStart > y2 {
			break
		}
		if b.yEnd < y1 {
			continue
		}
		if b.minX > x1 || b.maxX < x2 {
			return false
		}
	}

	covered := y1
	for i := lo; i < len(p.bands) && covered <= y2; i++ {
		b := p.bands[i]
		if b.yStart > covered {
			return false // gap
		}
		if b.yStart <= covered && b.yEnd >= covered {
			covered = b.yEnd + 1
		}
	}

	return covered > y2
}
