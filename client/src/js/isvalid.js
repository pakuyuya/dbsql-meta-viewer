let m = {}
m.regex = (pattern) => {
    // isPrevC changes as follows
    // * after `a`, `\a`, `[a]`, `(a)`         => true
    // * after `a*`, `a{1,1}`, `a|`, `(`, `a?` => false
    let isAllowRepeat = false
    let isAllowOptQues   = false
    let isAllowPoorQues   = false

    let parenthesesNests = 0
    const len = pattern.length
    let i = 0;
    while (i < len) {
        const c = pattern[i++]
        if (c === '*' || c === '+') {
            if (!isAllowRepeat) {
                return false
            }
            isAllowRepeat = false
            isAllowOptQues = false
            isAllowPoorQues = true
        } else if (c === '?') {
            if (isAllowOptQues) {
                isAllowRepeat = false
                isAllowOptQues = false
                isAllowPoorQues = true
            } else if (isAllowPoorQues) {
                isAllowRepeat = false
                isAllowOptQues = false
                isAllowPoorQues = false
            } else {
                return false;
            }
        } else if (c === '|') {
            isAllowRepeat = false
            isAllowOptQues = false
            isAllowPoorQues = false
        } else if (c === '\\') {
            if (!execEscaped()) {
                return false
            }
            isAllowRepeat = true
            isAllowOptQues = true
            isAllowPoorQues = false
        } else if (c === '(') {
            ++parenthesesNests;
            isAllowRepeat = false
            isAllowOptQues = false
            isAllowPoorQues = false
        } else if (c === ')') {
            if (--parenthesesNests < 0) {
                return false
            }
            isAllowRepeat = true
            isAllowOptQues = true
            isAllowPoorQues = false
        } else if (c === '{') {
            if (!execWave()) {
                return false
            }
            isAllowRepeat = false
            isAllowOptQues = true
            isAllowPoorQues = false
        } else if (c === '}') {
            return false;
        } else if (c === '[') {
            if (!execSquare()) {
                return false
            }
            isAllowRepeat = true
            isAllowOptQues = true
            isAllowPoorQues = false
        } else {
            isAllowRepeat = true
            isAllowOptQues  = true
            isAllowPoorQues = false
        }
    }
    return parenthesesNests === 0

    function execEscaped() {
        return (++i <= len)
    }

    function execWave() {
        const entryi = i
        while (i < len) {
            const c = pattern.charCodeAt(i++)
            if (c === 0x7d) {
                // c === }
                return entryi < i-1
            }
            if (c === 0x2c) {
                // c === ,
                break
            }
            if (c < 0x30 || c > 0x39) {
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
            if (c < 0x30 || c > 0x39) {
                // not a number
                return false
            }
        }
        return false
    }

    function execSquare() {
        let prevLetter = false
        let prevHyphen = false
        let hasLetter = false
        while (i < len) {
            const c = pattern[i++]
            if (i <= 1 && c === '^') {
                continue
            }
            if (c === ']') {
                return !prevHyphen && hasLetter
            } else if (c === '-') {
                prevHyphen = prevLetter
                prevLetter = !prevLetter
            } else {
                hasLetter = true
                prevLetter = true
                prevHyphen = false
            }
        }
        return false
    }

}

export default m