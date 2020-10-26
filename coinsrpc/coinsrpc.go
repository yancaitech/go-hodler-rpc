package coinsrpc

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// ENTROPY entry
type ENTROPY struct{}

// BTC entry
type BTC struct{}

// BCH entry
type BCH struct{}

// BSV entry
type BSV struct{}

// LTC entry
type LTC struct{}

// ETH entry
type ETH struct{}

// XRP entry
type XRP struct{}

// EOS entry
type EOS struct{}

// FIL entry
type FIL struct{}

// DOT entry
type DOT struct{}

const (
	// ServeTLS mode
	ServeTLS = true
)

var server *http.Server
var serverPort = 34911

// StartRPCMain func
func StartRPCMain(port int) {
	if port > 0 && port < 65535 {
		serverPort = port
	}
	go RPCMain()
}

// StopRPCMain func
func StopRPCMain() {
	if server != nil {
		server.Shutdown(context.Background())
	}
}

// RPCMain entry
func RPCMain() {
	s := rpc.NewServer()

	jsonCodec := json.NewCodec()
	s.RegisterCodec(jsonCodec, "application/json")
	s.RegisterCodec(jsonCodec, "application/x-www-form-urlencoded")
	s.RegisterCodec(jsonCodec, "charset=UTF-8")

	err := s.RegisterService(new(ENTROPY), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(BTC), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(BCH), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(BSV), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(LTC), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(ETH), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(XRP), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(EOS), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(FIL), "")
	if err != nil {
		fmt.Println(err)
	}
	err = s.RegisterService(new(DOT), "")
	if err != nil {
		fmt.Println(err)
	}

	localAddr := ":"
	localAddr += strconv.FormatInt(int64(serverPort), 10)
	envport := os.Getenv("COINSRPCPORT")
	if len(envport) > 2 {
		localAddr = ":" + envport
	}
	origins := handlers.AllowedOrigins([]string{"*"})
	server = &http.Server{Addr: localAddr, Handler: handlers.CORS(origins)(s)}

	if ServeTLS {
		crtFile, keyFile, err := generateCrtFiles()
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("load server.crt at:", crtFile)
		fmt.Println("load server.key at:", keyFile)

		err = server.ListenAndServeTLS(crtFile, keyFile)
	} else {
		err = server.ListenAndServe()
	}
	if err != nil {
		fmt.Println(err)
	}
}

func generateCrtFiles() (crtFile string, keyFile string, err error) {
	if runtime.GOOS == "android" {
		crtFile = "/sdcard/crpcser.crt"
	} else {
		crtFile, err = os.UserCacheDir()
		if err != nil {
			return "", "", err
		}
		crtFile = crtFile + "/crpcser.crt"
	}
	f, err := os.OpenFile(crtFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", "", err
	}
	f.WriteString(`-----BEGIN CERTIFICATE-----
MIICpjCCAY4CCQCJ+9nPUg96CzANBgkqhkiG9w0BAQsFADAUMRIwEAYDVQQDDAls
b2NhbGhvc3QwIBcNMTkwMzA3MDg1NTE4WhgPNDQ1MjExMTQwODU1MThaMBQxEjAQ
BgNVBAMMCWxvY2FsaG9zdDCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
AMeJtpGl3l9xn88xnf7NptCdZJETJAX4i2tqGzMutvpuDNN35N7AtbZjaavyktNK
1jFllA9lrGE1Fc5dlgKMjEb0KQdg4qI2cPv5+55THq2HJkNBi7iEACTRJ5nyJ7A9
blbtMkdM9v98RFkoRR/8EGaS04LHKQv0rwoncMpiAaw9J90XqOWEfLBWThVD/7bR
FE1l+keFnw40cVqvcnlkwL7KBh056O60zzL1asNqmvEY4b1qNedKj8xOQslQ9WrU
SLUym79Hv8Ga1sQcstPf0rGkkdGrS/Kfn552xD7jtYswhD1XZi7cHfw6ALR5fcti
Fxz3k39LOl5RPh5bi53wFE8CAwEAATANBgkqhkiG9w0BAQsFAAOCAQEAGydxxPUU
Nqio8lkBWgZBjLHbXZvanQpEKIMV33to0ZTsQAcjnd3Ic479t3ynY3jtr/1e2Q2v
ggjvVUlhlW3V3zZxqiQ6qwiM9DMOIUTJWRU8LC2bkRCPgS6KMWwOXpihodK0hqnU
jUmFzKhOZ7V6p9wyqHRaQRp04nCvb5jAoVcc4D+Tw0ySXbtyY4mioyjKSRmY4Zad
zhIDw6L9Vw25KRqgcQyOPFTff//BeasWeoup+85OkvhBBI98XVi9jwIjm6Fw4XvX
IUeWe4c20heY6PuVQk4gykOtgeqlA+2aAvwoRpPKZMeCdjDjErTFuCHzBBuBz9gz
L64pB/Z14W8e8Q==
-----END CERTIFICATE-----`)
	f.Close()

	if runtime.GOOS == "android" {
		keyFile = "/sdcard/crpcser.key"
	} else {
		keyFile, err = os.UserCacheDir()
		if err != nil {
			return "", "", err
		}
		keyFile = keyFile + "/crpcser.key"
	}
	f, err = os.OpenFile(keyFile, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return "", "", err
	}
	f.WriteString(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAx4m2kaXeX3GfzzGd/s2m0J1kkRMkBfiLa2obMy62+m4M03fk
3sC1tmNpq/KS00rWMWWUD2WsYTUVzl2WAoyMRvQpB2DiojZw+/n7nlMerYcmQ0GL
uIQAJNEnmfInsD1uVu0yR0z2/3xEWShFH/wQZpLTgscpC/SvCidwymIBrD0n3Reo
5YR8sFZOFUP/ttEUTWX6R4WfDjRxWq9yeWTAvsoGHTno7rTPMvVqw2qa8RjhvWo1
50qPzE5CyVD1atRItTKbv0e/wZrWxByy09/SsaSR0atL8p+fnnbEPuO1izCEPVdm
Ltwd/DoAtHl9y2IXHPeTf0s6XlE+HluLnfAUTwIDAQABAoIBAChTl7c124xUjWS1
dWXysB0HQjTjtmsSgTLTPe4JKefQ5/yxBWtTbSYQT2H142CahlFzYwn9lWoL7X1I
grG/L/aDD+uy0/rOn+T877JewBt37e63x4IiA1ltf+BoBUdkEvG0fG3WizC1L6cJ
MXc2XSYt1Ftg3tXQvLODpxLv3cCuO4+32PclcRcNnjWLvHC54lRtL/jQANnxBn+u
CEHfu/s7ONhl/iY2zGXLvmpQxUtezKhJQZWZR5acqY/BfvSGpTJ2FCFLLxM0hwJM
ZYt//qT5H3zC9l3W1HQodqQE8yVJTOYO4m+8FekEEyDXt2NhfjvRBheZxx1fy43M
/tYXGZkCgYEA+saebbwaPcrSkcC8CzOl4Q1C1zJnGhvwMhwGYLxXXp26qlqNs0+N
j5yU+RfOiFprZcTTPqXNbUQDVIh2/MEX8V0Ffh6PqKJ7+pCP+c67qFqwT1CioH+R
yoaaig+KG5l85wv3aCDHtkHRA2ytzW17LsE4lP4w1lQH7dVE6VAwhMsCgYEAy7HY
BZ+KSXqPDoxtyQPtCkThx/1JyDCQMqEWgYOv0awGVKykUCLCTjUOqbh1jzGcB87N
3CkjWjzokR9prlOx6M4mOTiXxEYmzML5rY+AFjhEBIphRFuCOGGx7A05fohMEfJZ
s4OafZ/4Ogynoz2+DjxBddLYCwgVsJXtFfEXQg0CgYBlWxF9WKFiiC9DKZrXDDDn
HOz+/SgerVwPZLRPNNA7NZTUdXUAHA8jFC5B3xVilukBYOPgVjMJDowqBl3RGloK
+4XUy5VUmxdw1iza0muWR9EqvXR9WhIawPyFAHLZZNfOqk98joMpbsCDmdFFThKT
exTbY0Fp3ty1i5Uml3qEsQKBgBBmwIsrXnouKSi1u/1MmKCUDU4KIg/BgUriV6qU
DOsoG9ZjlFNziQu9D/IwjR67kuG6EC9jDJquftd4nQzRZmjleIRw/x9puqQQKwSD
on+nhiqLbeuQJrsderMUGYYLuXxUdE6VETJ+WAFslW76gLwqs+al1ImG3CA84js7
D3FdAoGBAObx0oF9i5zXKA4hTyRts1L6+7l8VY928lxvKVpD4FMTBfrCcVlgcYXC
u/8I5NWNddfap/pI6XbKvBWpTIIb4bSTcXiVAQVXWO6SpOt6Sz+ve9bajZIyqFKb
+LBVZWMJC4LOAUltNVC8N4G4rzeaZGDadt/+8fcM2Vbl7vRhwfYO
-----END RSA PRIVATE KEY-----`)
	f.Close()
	return crtFile, keyFile, nil
}
