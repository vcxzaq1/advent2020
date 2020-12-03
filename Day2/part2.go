package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

type Policy struct {
    firstPosition string
    secondPosition string
    check string
    password string
    og string
}

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    validPasswords := 0
    for scanner.Scan(){
        policy := getPolicy(scanner.Text())
        if policy.isValid() {
            validPasswords = validPasswords + 1
        }
    }

    fmt.Println(validPasswords)
}


func getPolicy(a string) Policy {
    rule, password := splitString(a, ":")
    counts, _ := splitString(rule, " ")
    min, max := splitString(counts, "-")
    one := getInt(min)
    two := getInt(max)

    checkIndex := strings.IndexByte(a, ':') - 1
    check := string(a[checkIndex])

    p := Policy{
        firstPosition: string(password[one]),
        secondPosition: string(password[two]),
        check: check,
        password: strings.TrimSpace(password),
    }

    return p
}

func (p Policy) isValid() bool {
    if p.firstPosition == p.secondPosition {
        return false
    }

    if !strings.Contains(p.password, p.check) {
        return false
    }

    if !strings.Contains(p.firstPosition + p.secondPosition, p.check) {
        return false
    }

    return true
}
func getInt(val string) int {
    num, _ := strconv.Atoi(val)
    return num
}

func splitString(s, sep string) (string, string) {
    split := strings.Split(s, sep)
    return split[0], split[1]
}

