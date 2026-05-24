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
	line int
	column int


}
type Groups struct{

	nodeName string
	value Expression
	line int
	column int


}
type Binary struct{

	nodeName string
	left Expression
	right Expression
	operator lexer.TokenType
	line int
	column int


}
type Unary struct{

	nodeName string
	value Expression
	line int
	column int


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


