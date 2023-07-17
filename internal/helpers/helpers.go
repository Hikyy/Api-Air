package helpers

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/joho/godotenv"
	"github.com/pjebs/optimus-go"
	"golang.org/x/crypto/bcrypt"
)

var destinationPkg string
var Type string

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
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load .env file:", err)
	}
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

func EncodeId(id int) string {
	o := InitOptimus()
	newId := o.Encode(uint64(id))

	return strconv.FormatUint(newId, 10)
}

// Refacto this method @Hikyy
func FillStruct(destination interface{}, source interface{}) {
	destinationValue := reflect.ValueOf(destination)

	if reflect.TypeOf(destination).Elem().PkgPath() == "App/internal/resources" {
		// Type = reflect.TypeOf(source).Elem().Name()
		t := reflect.TypeOf(source)
		Type = ""
		if t.Kind() == reflect.Ptr {
			Type = t.Elem().Name()
		} else {
			Type = t.Name()
		}
		destinationPkg = reflect.TypeOf(destination).Elem().PkgPath()
	}

	if destinationValue.Kind() != reflect.Ptr || destinationValue.IsNil() {
		panic("destination must be a non-nil pointer")
	}

	destinationValue = destinationValue.Elem()
	sourceValue := reflect.ValueOf(source)

	if sourceValue.Kind() == reflect.Ptr {
		sourceValue = sourceValue.Elem()
	}

	if destinationValue.Kind() == reflect.Slice {
		for z := 0; z < destinationValue.Len(); z++ {
			FillStruct(destinationValue.Field(z).Addr(), sourceValue.Field(z).Addr())
		}
	}

	if sourceValue.Kind() == reflect.Slice {
		for i := 0; i < sourceValue.Len(); i++ {
			singleSource := sourceValue.Index(i).Interface()
			singleDestination := reflect.New(destinationValue.Type().Elem())
			FillStruct(singleDestination.Interface(), singleSource)
			destinationValue.Set(reflect.Append(destinationValue, singleDestination.Elem()))
		}
	} else {
		for i := 0; i < destinationValue.NumField(); i++ {
			destinationField := destinationValue.Field(i)
			for j := 0; j < sourceValue.NumField(); j++ {
				sourceField := sourceValue.Type().Field(j)
				if destinationValue.Type().Field(i).Name == sourceField.Name {

					if destinationPkg == "App/internal/resources" {
						id := ""
						if sourceField.Name == "Id" && destinationValue.Type().Field(i).Name == "Id" {
							id = EncodeId(int(sourceValue.Field(j).Int()))
						}
						if destinationField.CanSet() {
							if id == "" {
								destinationField.Set(sourceValue.Field(j))
							} else {
								if sourceValue.Kind() == reflect.Int && destinationValue.Kind() == reflect.Int {
									id, err := strconv.Atoi(id)

									if err != nil {
										fmt.Println("Error during conversion")
										return
									}

									destinationField.SetInt(int64(id))
								} else {
									destinationField.SetString(id)
								}
							}
						}
					} else {
						if destinationField.CanSet() {
							destinationField.Set(sourceValue.Field(j))
						}
					}
				}
				if destinationPkg == "App/internal/resources" && destinationValue.Type().Field(i).Name == "Type" && Type != "" {
					destinationField.SetString(Type)
				}
			}

			if destinationField.Kind() == reflect.Struct {
				FillStruct(destinationField.Addr().Interface(), source)
			} else if destinationField.Kind() == reflect.Slice {
				processSliceField(destinationField, source)
			} else {
				fmt.Printf("Unhandled field type: %s\n", destinationField.Kind())
			}
		}
	}
}

func processSliceField(field reflect.Value, source interface{}) {
	for i := 0; i < field.Len(); i++ {
		FillStruct(field.Index(i).Addr().Interface(), source)
	}
}
