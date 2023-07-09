package helpers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/pjebs/optimus-go"
	"golang.org/x/crypto/bcrypt"
)

// ParseForm pour décoder l'http request de Gorillan schema
func ParseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	dec := schema.NewDecoder()
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	fmt.Println(bytes)
	if err != nil {
		fmt.Println(err)
	}
	return string(bytes)
}

func InitOptimus() optimus.Optimus {
	optimusPrime, _ := strconv.Atoi(os.Getenv("OPTIMUS_PRIME"))
	optimusInverse, _ := strconv.Atoi(os.Getenv("OPTIMUS_INVERSE"))
	optimusRandom, _ := strconv.Atoi(os.Getenv("OPTIMUS_RANDOM"))

	fmt.Println("optimusPrime : ", optimusPrime, "optimusInverse : ", optimusInverse, "optimusRandom : ", optimusRandom)
	
	return optimus.New(uint64(optimusPrime), uint64(optimusInverse), uint64(optimusRandom))
}


func CheckPassword(hashedPassword string, password string) bool {
	bsp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword(bsp, []byte(hashedPassword))
	if err != nil {
		panic(err)
	} else {
		fmt.Println("password are equal")
		return true
	}
	return false
}
