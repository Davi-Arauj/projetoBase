package util

import "math"

// Max retorna o valor máximo de um slice
func Max(valores ...int) int {
	if len(valores) == 0 {
		return math.MaxInt32
	}

	max := valores[0]
	for _, v := range valores {
		if v > max {
			max = v
		}
	}

	return max
}

// Min retorna o valor mínimo de um slice
func Min(valores ...int) int {
	if len(valores) == 0 {
		return math.MinInt32
	}

	min := valores[0]
	for _, v := range valores {
		if v < min {
			min = v
		}
	}

	return min
}

// Max64 retorna o valor máximo de um slice
func Max64(valores ...int64) int64 {
	if len(valores) == 0 {
		return math.MaxInt64
	}

	max := valores[0]
	for _, v := range valores {
		if v > max {
			max = v
		}
	}

	return max
}

// Min64 retorna o valor mínimo de um slice
func Min64(valores ...int64) int64 {
	if len(valores) == 0 {
		return math.MinInt64
	}

	min := valores[0]
	for _, v := range valores {
		if v < min {
			min = v
		}
	}

	return min
}

// Contem retorna true se o valor "v" está contido
// dentre as opções especificadas
func Contem(v int, opcoes ...int) bool {
	for _, k := range opcoes {
		if v == k {
			return true
		}
	}

	return false
}

// Contem64 retorna true se o valor "v" está contido
// dentre as opções especificadas
func Contem64(v int64, opcoes ...int64) bool {
	for _, k := range opcoes {
		if v == k {
			return true
		}
	}

	return false
}

// AjustarCasasDecimais trunca para o maior valor a quantidade de casas definidas
//	ex.: o valor 176.37485 com ajuste em duas casa decimais se torna 176.38
func AjustarCasasDecimais(v *float64, casas float64) *float64 {
	*v *= math.Pow(10.0, casas)
	*v = math.Ceil(*v)
	*v /= math.Pow(10.0, casas)

	return v
}
