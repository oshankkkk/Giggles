package lexer 
import (
	"os"
	"bufio"
	"fmt"
	"strings"
)
func ReadFile(file *os.File){
	//separators := [4]string{"(", ")", ",", "\""}
	separators :="()\","
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		lexer(scanner.Text(),separators)
		// scanner.Text() gives you the line as a string
		fmt.Println(scanner.Text())
	}


}

func lexer(source string,seperators string){

// use case statement to look at each string and put them into tokens
tokens:=[]string{}
buff:=""
for _,value:=range source{
	
	if strings.Contains(seperators,string(value)){
		tokens = append(tokens, buff)
		buff=""
	}
	buff+=string(value)
}
fmt.Println("this is buff:",tokens)
}






