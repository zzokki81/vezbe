package tempfunctions



import (

	"bufio"
	"os"
	"fmt"
	"errors"
)
//func Input() will take user input
func Input() (string,error){
reader := bufio.NewReader(os.Stdin)
fmt.Println("> Please input temperature:")
fmt.Print("< ")

userInput,err := reader.ReadString('\n')
if err != nil {
return "" ,errors.New("Unknown type..")
}
return userInput,nil
}
