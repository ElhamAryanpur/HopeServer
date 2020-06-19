package hopeserver

import b64 "encoding/base64"

// Encrypt is where base64 encryption happen 
func Encrypt(data string) string {
	return b64.RawURLEncoding.EncodeToString([]byte(data));
}

// Decrypt is where base64 decryption happen 
func Decrypt(data string) string {
	dec, err := b64.RawURLEncoding.DecodeString(data);
	if (err != nil){
		print((err))
	}
	return string(dec)
}
