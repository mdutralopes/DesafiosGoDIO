package main
import "testing"
func TestSoma(t *testing.T) { //convensão de nomes ShouldSumCorrect 
(assinatura do método)
//arrange
teste := soma(3, 2, 1)
//act
resultado := 6
//assert
if teste != resultado { //assert
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestSoma2(t *testing.T) { //ShouldSumIncorrect
teste := soma(3, 2, 1) //arrange
resultado := 7 //act
if teste != resultado { //assert
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestMult(t *testing.T) {
teste := multiplica(10, 10)
resultado := 100
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestMult2(t *testing.T) {
teste := multiplica(10, 10)
resultado := 2560
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestSub(t *testing.T) {
teste := subtrai(10, 5)
resultado := -5
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestSub2(t *testing.T) {
teste := subtrai(10, 10)
resultado := 5
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestDiv(t *testing.T) {
teste := divide(10)
resultado := 10
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
}
func TestDiv2(t *testing.T) { 
teste := divide(10)
resultado := 5
if teste != resultado {
t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
}
