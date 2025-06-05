package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "regexp"
    "strings"
)

func main() {
    fmt.Println("Email Checker Tool")
    fmt.Print("Enter email address to check: ")
    
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    email := scanner.Text()
    
    if isValidEmail(email) {
        fmt.Printf("✅ %s is a valid email format\n", email)
        checkDomain(email)
    } else {
        fmt.Printf("❌ %s is not a valid email format\n", email)
    }
}

func isValidEmail(email string) bool {
    emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
    re := regexp.MustCompile(emailRegex)
    return re.MatchString(email)
}

func checkDomain(email string) {
    parts := strings.Split(email, "@")
    if len(parts) != 2 {
        return
    }
    
    domain := parts[1]
    _, err := net.LookupMX(domain)
    if err != nil {
        fmt.Printf("⚠️  Domain %s may not accept emails\n", domain)
    } else {
        fmt.Printf("✅ Domain %s can receive emails\n", domain)
    }
}
