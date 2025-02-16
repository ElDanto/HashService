package redis

import (
	"encoding/base64"
	"github.com/google/uuid"
	"regexp"
)

type Hash struct {
	Hash string `redis:"hash"`
}

func (h *Hash) Generate() string {
	v4, _ := uuid.NewRandom()
	re := regexp.MustCompile(`^(.{8})-(.{4})-4(.{3})-(.{4})-(.{12})$`)
	hash := re.ReplaceAllString(v4.String(), "$1$2$3$4$5")
	hash = base64.StdEncoding.EncodeToString([]byte(hash))
	hash = hash[:3] + hash[18:21] + hash[32:35]
	return hash
}
