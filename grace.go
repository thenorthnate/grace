package grace

const (
	OpAdd = iota
	OpSub
	OpMul
	OpDiv

	gtU8  = "uint8"
	gtF64 = "float64"
	gtF32 = "float32"
)

// Grace is an interface that supports all operations for matrix manipulations
type Grace interface {
	Display(depth int)
	IsLeaf() bool
}

// Zeros creates a new matrix of zeros with the given dimensions and data type
func Zeros(dtype string, dims ...int) Grace {
	return nil
}

// Show replaces "head" and "tail" for other frameworks. It prints out the matrix to the depth specified.
// If depth == 0, it will print the entire matrix. If depth is negative, it acts like "tail". Positive
// values will print out the first rows of a matrix.
func Show(g Grace, depth int) {
	g.Display(depth)
}
