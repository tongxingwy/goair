package utils

import (
	"fmt"
	"log"
	"crypto/rand"
)

const (
	HTTPProtocolType = "HTTP/1.1"
	RTSPProtocolType = "RTSP/1.0"
	carReturn        = "\r\n"
)

func CreateHttpResponse(success bool, isPlist bool, headers map[string]string, data []byte) []byte {
	s := HTTPProtocolType
	if success {
		s += " 200 OK" + carReturn
		if data != nil {
			if isPlist {
				s += fmt.Sprintf("Content-Type: text/x-apple-plist+xml%s", carReturn)
			} else {
				s += fmt.Sprintf("Content-Type: application/octet-stream%s", carReturn)
			}
			s += fmt.Sprintf("Content-Length: %d%s", len(data), carReturn)
		}
		for key, val := range headers {
			s += fmt.Sprintf("%s: %s%s", key, val, carReturn)
		}
	} else {
		s += " 400 Bad Request" + carReturn
	}
	log.Println("response is:", s)
	body := []byte(s + carReturn)
	if data != nil {
		body = append(body, data...)
	}
	return body
}

func CreateRtspResponse(success bool, isPlist bool, headers map[string]string, data []byte) []byte {
	s := RTSPProtocolType
	if success {
		s += " 200 OK" + carReturn
		if data != nil {
			if isPlist {
				s += fmt.Sprintf("Content-Type: text/x-apple-plist+xml%s", carReturn)
			} else {
				s += fmt.Sprintf("Content-Type: application/octet-stream%s", carReturn)
			}
			s += fmt.Sprintf("Content-Length: %d%s", len(data), carReturn)
		}
		for key, val := range headers {
			s += fmt.Sprintf("%s: %s%s", key, val, carReturn)
		}
	} else {
		s += " 400 Bad Request" + carReturn
	}
	log.Println("------------------- response -----------------")
	log.Println("response is:", s)
	log.Println("body: ", data)
	body := []byte(s + carReturn)
	if data != nil {
		body = append(body, data...)
	}
	return body
}

func RandomBytes(count int) []byte {
	b := make([]byte, count)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return b
}
