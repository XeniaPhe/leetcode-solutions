package compare_version_numbers

func compareVersion(version1 string, version2 string) int {
    for pos1, pos2, rev1, rev2 := 0, 0, 0, 0; pos1 < len(version1) || pos2 < len(version2); {
        rev1, pos1 = getRevision(version1, pos1)
		if rev2, pos2 = getRevision(version2, pos2); rev1 < rev2 {
			return -1
		} else if rev2 < rev1 {
            return 1
        }
    }
    
    return 0
}

func getRevision(ver string, pos int) (rev, nextPos int) {
    for ; pos < len(ver) && ver[pos] != '.'; pos, rev = pos + 1, 10 * rev + int(ver[pos] - '0') { }
    return rev, pos + 1
}