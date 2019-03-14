let m = {}
m.regex = (pattern) => {
    // isPrevC changes as follows
    // * after `a`, `\a`, `[a]`, `(a)`         => true
    // * after `a*`, `a{1,1}`, `a|`, `(`, `a?` => false
    let isAllowRepeat = false
    let isAllowQues   = false

    let parenthesesNests = 0
    const len = pattern.length
    let i = 0;
    while (i<len) {
        const c = pattern[i]
        ++i
        if (c === '*' || c === '+') {
            if (!isAllowRepeat) {
                return false
            }
            isAllowRepeat = false
            isAllowQues = true
        } else if (c === '?') {
            if (!isAllowQues) {
                return false
            }
            isAllowRepeat = false
            isAllowQues = false
        } else if (c === '|') {
            isAllowRepeat = false
            isAllowQues = false
        } else if (c === '\\') {
            if (!execEscaped()) {
                return false
            }
            isAllowRepeat = true
            isAllowQues = true
        } else if (c === '(') {
            ++parenthesesNests;
            isAllowRepeat = false
            isAllowQues = false
        } else if (c === ')') {
            if (--parenthesesNests < 0) {
                return false
            }
            isAllowRepeat = true
            isAllowQues = true
        } else if (c === '{') {
            if (execWave()) {
                return false
            }
            isAllowRepeat = false
            isAllowQues = true
        } else if (c === '[') {
            if (execSquare()) {
                return false
            }
            isAllowRepeat = true
            isAllowQues = true
        }
    }
    
    return parenthesesNests === 0

    function execEscaped() {
        return (++i > len)
    }

    function execWave() {
        while (i < len) {
            const c = pattern.charCodeAt(i)
            ++i
            if (c === 0x7d) {
                // c === }
                return true
            }
            if (c === 0x2c) {
                // c === ,
                break
            }
            if (c > 0x30 || c < 0x39) {
                // not a number
                return false
            }
        }
        while (i < len) {
            const c = pattern.charCodeAt(i)
            ++i
            if (c === 0x7d) {
                // c === }
                return true
            }
            if (c > 0x30 || c < 0x39) {
                // not a number
                return false
            }
        }
        return false
    }

    function execSquare() {
        let prevHyphen = false
        while (i < len) {
            const c = pattern[i]
            if (c === ']') {
                return !prevHyphen
            } else if (c === '-') {
                prevHyphen = !prevHyphen
            } else {
                prevHyphen = false
            }
        }
        return false
    }

}
const x

export default m