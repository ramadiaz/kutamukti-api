package helpers

import (
	"encoding/json"
	"io"
	"kutamukti-api/api/complaint/dto"
	"log"
	"net/http"
	"net/url"
	"os"
)

func VerifyRecaptcha(captchaResponse string) bool {
	secretKey := os.Getenv("RECAPTCHA_SECRET_KEY")
	if secretKey == "" {
		log.Println("RECAPTCHA_SECRET_KEY environment variable not set")
		return false
	}

	postData := url.Values{}
	postData.Set("secret", secretKey)
	postData.Set("response", captchaResponse)

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", postData)
	if err != nil {
		log.Printf("Error verifying reCAPTCHA: %v", err)
		return false
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading reCAPTCHA response: %v", err)
		return false
	}

	var recaptchaResp dto.RecaptchaResponse
	if err := json.Unmarshal(body, &recaptchaResp); err != nil {
		log.Printf("Error parsing reCAPTCHA response: %v", err)
		return false
	}

	if len(recaptchaResp.ErrorCodes) > 0 {
		log.Printf("reCAPTCHA error codes: %v", recaptchaResp.ErrorCodes)
	}

	return recaptchaResp.Success
}