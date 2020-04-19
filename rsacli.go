package rsalib

import "fmt"

func RsaKey() string {
	// TODO: read information from the command line
	len := 4096
	pwd := "T3sT_demo!234"

	kp, err := GenerateKeyPair(len, pwd); if err != nil {
		fmt.Println("An error occurred during the generation process")
	} else {
		fmt.Println(kp.PrivateKey)
		fmt.Println("")
		fmt.Println(kp.PublicKey)
	}
	
	return "test"
}
