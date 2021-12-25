package controllers

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"

	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	encryptor := sha1.New()
	token := "XXX"
	signature := c.Params.Query.Get("signature")
	timestamp := c.Params.Query.Get("timestamp")
	nonce := c.Params.Query.Get("nonce")
	echostr := c.Params.Query.Get("echostr")

	tmpArray := []string{token, timestamp, nonce}
	sort.Strings(tmpArray)
	tmpStr := strings.Join(tmpArray[:], ",")
	encryptor.Write([]byte(tmpStr))
	encryptedString := base64.URLEncoding.EncodeToString(encryptor.Sum(nil))

	fmt.Printf("tmpArray: %s\n", tmpArray)
	fmt.Printf("signature: %s\n", signature)
	fmt.Printf("encryptedString: %s\n", encryptedString)

	if signature == encryptedString {
		return c.RenderText(echostr)
	}

	return c.RenderText(encryptedString)
}
