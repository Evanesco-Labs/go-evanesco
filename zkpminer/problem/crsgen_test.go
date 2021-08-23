package problem

import (
	"os"
	"testing"
)

func TestGenCrs(t *testing.T) {
	r1csFile,err := os.OpenFile("./r1cstest.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil{
		panic(err)
	}
	defer r1csFile.Close()
	pkFile, err := os.OpenFile("./provekeytest.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil{
		panic(err)
	}
	defer pkFile.Close()
	skFile,err := 	os.OpenFile("./verifykeytest.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil{
		panic(err)
	}
	defer skFile.Close()

	r1cs := CompileCircuit()
	r1cs.WriteTo(r1csFile)

	pk,sk := SetupZKP(r1cs)
	pk.WriteTo(pkFile)
	sk.WriteTo(skFile)
}