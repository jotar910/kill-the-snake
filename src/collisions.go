package main

import (
	"math"
)

type Body struct {
	position *Vector
	radius   float64
}

func circleCollides(c1 *Body, c2 *Body) bool {
	p1 := c1.position
	p2 := c2.position
	dist := math.Sqrt(math.Pow(p1.X-p2.X, 2.0) + math.Pow(p1.Y-p2.Y, 2.0))
	return dist <= c1.radius+c2.radius
}

/* func CheckCollisions(elements []*Element) error {
	lenElements := len(elements)
	for i := 0; i < lenElements; i++ {
		for j := 0; j < lenElements; j++ {
			if i == j || !elements[i].Active || !elements[j].Active {
				continue
			}

			b1 := elements[i].BodyHit
			for _, c := range elements[j].Collisions {
				if c.DoCollides(&b1) {
					c.Collision(&b1)
				}
			}
		}
	}
	return nil
} */
