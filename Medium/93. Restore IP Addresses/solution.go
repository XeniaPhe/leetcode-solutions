package restore_ip_addresses

import "unsafe"

func restoreIpAddresses(s string) []string {
    return resolveAllIpAddresses(s, make([]string, 0, 4), []string{})
}

func resolveAllIpAddresses(s string, path []string, result []string) []string {
    if lp, ls := len(path), len(s); lp == 4 && ls == 0 {
        idx, bytes := 0, make([]byte, len(path[0]) + len(path[1]) + len(path[2]) + len(path[3]) + 4)
        for i := 0; i < 4; idx, i = idx + 1, i + 1 {
            idx += copy(bytes[idx:], path[i])
            bytes[idx] = '.'
        }

        return append(result, unsafe.String(unsafe.SliceData(bytes[:idx - 1]), idx - 1))
    } else if rem := 4 - lp; 3 * rem < ls || rem > ls {
        return result
    }

    if len(s) > 2 && (s[0] == '1' || (s[0] == '2' && (s[1] < '5' || (s[1] == '5' && s[2] < '6')))) {
        result = resolveAllIpAddresses(s[3:], append(path, s[:3]), result)
    }

    if len(s) > 1 && s[0] > '0' {
        result = resolveAllIpAddresses(s[2:], append(path, s[:2]), result)
    }

    return resolveAllIpAddresses(s[1:], append(path, s[:1]), result)
}