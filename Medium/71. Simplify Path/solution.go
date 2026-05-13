package simplify_path

import "unsafe"

func simplifyPath(path string) string {
    dotCount, buf, dirs := 0, make([]byte, 1, len(path)), make([]int, 1, 16)
    buf[0], dirs[0] = '/', 1

    handleDots := func() {
        switch dotCount {
        case 1:
            dotCount = 0
        case 2:
            lastDir := 1
            if lastDirIdx := len(dirs) - 1; lastDirIdx >= 1 {
                lastDir, dirs = dirs[lastDirIdx - 1], dirs[:lastDirIdx]
            }

            dotCount, buf = 0, buf[:lastDir]
        default:
            for j := 0; j < dotCount; j += 1 {
                buf = append(buf, '.')
            }
        }
    }

    for i := 0; i < len(path); i += 1 {
        if chr := path[i]; chr == '/' {
            handleDots()
            if dotCount = 0; buf[len(buf) - 1] != '/' {
                buf = append(buf, '/')
                dirs = append(dirs, len(buf))
            }
        } else if chr == '.' && buf[len(buf) - 1] == '/' {
            dotCount += 1
        } else {
            for j := 0; j < dotCount; j += 1 {
                buf = append(buf, '.')
            }

            dotCount, buf = 0, append(buf, chr)
        }
    }

    handleDots()
    if buf[len(buf) - 1] == '/' && len(buf) != 1 {
        buf = buf[:len(buf) - 1]
    }

    return unsafe.String(unsafe.SliceData(buf), len(buf))
}