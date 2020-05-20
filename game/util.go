package game

// Intersec2D Compute the intersection between 2 lines
/* ********************************************************************** */
/*                                                                        */
/*  Calcula a interseccao entre 2 retas (no plano "XY" Z = 0)             */
/*                                                                        */
/* k : ponto inicial da reta 1                                            */
/* l : ponto final da reta 1                                              */
/* m : ponto inicial da reta 2                                            */
/* n : ponto final da reta 2                                              */
/*                                                                        */
/* s: valor do parâmetro no ponto de interseção (sobre a reta KL)         */
/* t: valor do parâmetro no ponto de interseção (sobre a reta MN)         */
/*                                                                        */
/* ********************************************************************** */
func Intersec2D(k, l, m, n Point) (s, t float32, ok bool) {
	pk := k.Get()
	pl := l.Get()
	pm := m.Get()
	pn := n.Get()

	det := (pn.x-pm.x)*(pl.y-pk.y) - (pn.y-pm.y)*(pl.x-pk.x)

	if det == 0 {
		return 0, 0, false // no intersection
	}

	s = ((pn.x-pm.x)*(pm.y-pk.y) - (pn.y-pm.y)*(pm.x-pk.x)) / det
	t = ((pl.x-pk.x)*(pm.y-pk.y) - (pl.y-pk.y)*(pm.x-pk.x)) / det

	return s, t, true
}

// HasIntersection check if has intersection between two lines
func HasIntersection(k, l, m, n Point) bool {
	s, t, ok := Intersec2D(k, l, m, n)

	if !ok {
		return false
	}

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		return true
	}

	return false
}
