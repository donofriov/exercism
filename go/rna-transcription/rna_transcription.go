package strand

// ToRNA takes a dna string and returns the rna string equivalent
func ToRNA(dna string) string {
	rna := ""
	dnaToRna := map[string]string{
		"G": "C",
		"C": "G",
		"T": "A",
		"A": "U",
		"":  "",
	}

	for _, value := range dna {
		rna += dnaToRna[string(value)]
	}

	return rna
}
