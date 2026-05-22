package parser
import
("lang/internal/lexer"
"fmt")
type Expression interface{
	expression()
}
type Literal struct{

	nodeName string
	value lexer.Token
}
type Groups struct{

	nodeName string
	value Expression
}
type Binary struct{
	nodeName string
	left Expression
	right Expression
	operator lexer.TokenType
}
type Unary struct{
	value Expression
}
func (s Literal) expression(){
fmt.Println(s.nodeName)
}
func (s Groups) expression(){

fmt.Println(s.nodeName)
}
func (s Unary) expression(){
}
func (s Binary) expression(){
fmt.Println(s.nodeName)}


