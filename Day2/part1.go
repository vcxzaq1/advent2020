package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func main() {
    file, err := os.Open("input.txt")

    if err != nil {
        log.Fatal(err)
    }

    scanner := bufio.NewScanner(file)

    validPasswords := 0
    for scanner.Scan(){
        if validatePassword(scanner.Text()) == true {
            validPasswords = validPasswords + 1
        }
    }

    fmt.Println(validPasswords)

}

func validatePassword(input string) bool {
    line := strings.Split(input, ":")
    rule, password := strings.TrimSpace(line[0]), strings.TrimSpace(line[1])

    rulesSplit := strings.Split(rule, " ")
    ruleCounts, letter := strings.TrimSpace(rulesSplit[0]), strings.TrimSpace(rulesSplit[1])

    maxMin := strings.Split(ruleCounts, "-")
    min, max := maxMin[0], maxMin[1]

    minInt, _ := strconv.Atoi(min)
    maxInt, _ := strconv.Atoi(max)

    occurences := strings.Count(password, letter)

    if occurences >= minInt && occurences <= maxInt {
        return true
    }
    return false
}
