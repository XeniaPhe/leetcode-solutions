package repeated_dna_sequences

func findRepeatedDnaSequences(s string) []string {
    var nucToQuat = ['T' + 1]int{}
    nucToQuat['C'], nucToQuat['G'], nucToQuat['T'] = 1, 2, 3
	i, quat, tenth, seqs, repeated := 0, 0, 1 << 18, [1 << 20]byte{}, []string{}
	for ; i < min(10, len(s)); quat, i = 4 * quat + nucToQuat[s[i]], i + 1 { }
	
	for seqs[quat] = 1; i < len(s); i += 1 {
		switch quat = 4 * (quat - tenth * nucToQuat[s[i - 10]]) + nucToQuat[s[i]]; seqs[quat] {
		case 0: seqs[quat] = 1
		case 1: seqs[quat], repeated = 2, append(repeated, s[i - 9:i + 1])
		}
	}

	return repeated
}