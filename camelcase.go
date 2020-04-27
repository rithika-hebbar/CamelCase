//Author: Rithika Hebbar

package main
import "fmt"
import "strings"

func main(){
    var input string
    fmt.Println("Enter a sequence of words in CamelCase:")
    fmt.Scanf("%s\n",&input)
	noOfWords:=1
	
    for _, char:=range input{

        if strings.ToUpper(string(char))==string(char){
            noOfWords++
        }
    }
    fmt.Printf("There are %d words.\n",noOfWords)
    
}