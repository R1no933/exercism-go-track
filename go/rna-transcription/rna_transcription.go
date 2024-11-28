package strand

func ToRNA(dna string) string {
	resDna := ""
	for _, ch := range dna {
		switch ch {
		case 'A':
			resDna += "U"
		case 'T':
			resDna += "A"
		case 'C':
			resDna += "G"
		case 'G':
			resDna += "C"
		default:
			resDna = ""
		}
	}

	return resDna
}
