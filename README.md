Curso alura: 
    -> Go: Orientação a Objetos Go: Orientação a Objetos

Professor: 
	-> Guilherme Lima
	
Objetivo:
	-> No curso aprenderemos a trabalhar com os principais tipos da linguagem e ainda com structs, ponteiros, referências e funções.  Veremos funções com um retorno só e outras com múltiplos retornos e saberemos como utilizá-las e trabalhar com métodos.
	Também aprenderemos a trabalhar com composição, encapsulamento e interface.
	O foco principal do curso será a linguagem Go.
	
Observação:
    -> Aprender sobre concorrência.


-----XX-----XX-----XX-----XX-----XX-----XX-----XX-----XX-----XX-----XX-----XX-----



O que são structs em golang?
-> As structs é um tipo que contém campos nomeados.

Como declarar um struct?
-> Começamos com o a palavra reservada TYPE, em seguida colocamos o nome do 
tipo acompanhado da palavra struct, pois assim informamos ao go que estamos definindo um tipo struct. Por fim, abrimos e fechamos {}.

Exemplo:
  type Pessoa struct{
	nome            string
	idade           int
	numeroDocumento int
  } 

Observação importante: Por padrão o golang define o valor zero correspondente aos tipos de campos declarados. 0 para ints, 0.0 para floats, "" para strings, nil para ponteiros etc.

Exemplo:
	nome string -> será inicializada com ""
	idade int -> será inicializada com 0
	numeroDocumento int -> será inicializada com 0
