package hw10programoptimization

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	u, err := getUsers(r)
	if err != nil {
		return nil, fmt.Errorf("get users error: %w", err)
	}
	return countDomains(u, domain)
}

type users [100000]User

func getUsers(r io.Reader) (result users, err error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	i := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		user := &User{}
		if err = easyjson.Unmarshal(line, user); err != nil {
			return
		}
		result[i] = *user
		i++
	}
	return
}

func countDomains(u users, domain string) (DomainStat, error) {
	result := make(DomainStat)
	searchStr := fmt.Sprintf(".%s", domain)
	for _, user := range u {
		matched := strings.HasSuffix(user.Email, searchStr)

		if matched {
			num := result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])]
			num++
			result[strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])] = num
		}
	}
	return result, nil
}
