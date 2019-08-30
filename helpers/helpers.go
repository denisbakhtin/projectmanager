package helpers

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

//FileBasename returns base file name (without extension)
func FileBasename(s string) string {
	n := strings.LastIndexByte(s, '.')
	if n >= 0 {
		return s[:n]
	}
	return s
}

//GetUniqueFilename checks if file exists and attempts to find a unique name by appendind _num to original name
func GetUniqueFilename(dir string, fname string) (string, error) {
	if _, err := os.Stat(path.Join(dir, fname)); os.IsNotExist(err) {
		return fname, nil
	}
	ext := filepath.Ext(fname)
	attempt := 1
	for attempt < 100 {
		newname := FileBasename(fname) + "_" + strconv.Itoa(attempt) + ext
		if _, err := os.Stat(path.Join(dir, newname)); os.IsNotExist(err) {
			return newname, nil
		}
		attempt++
	}
	return "", errors.New("Please, rename file and try again")
}

//Substr returns substring of specified length
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//CheckNewPassword checks if password is safe
func CheckNewPassword(password string) error {
	if len(strings.Trim(password, " ")) < 8 {
		return errors.New("Password must be atleast 8 characters")
	}
	return nil
}

//NormalizeEmail strips spaces and converts to lower case
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.Trim(email, " "))
}

//CreatePasswordHash creates crypto hash of the password
func CreatePasswordHash(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

//CreateSecureToken creates new token for user activation / password recovery
func CreateSecureToken() string {
	str := fmt.Sprintf("t0kEn%v", time.Now().UnixNano())
	bytes, _ := bcrypt.GenerateFromPassword([]byte(str), 12)
	return base64.StdEncoding.EncodeToString(bytes)
}

//CheckPasswordHash checks if password has a specified hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
