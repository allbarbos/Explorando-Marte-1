# Explorando Marte
**Problema:** Movimentação de sondas por um planalto retangular controladas por uma sequência de letras.
### ENTRADA

Arquivo com informações do ambiente. Deve-se entrar com o nome do arquivo completo, com extensão. Caso o arquivo esteja em pasta diferente do executável, deve-se adicionar o caminho completo até ele.

**Exemplo:** /text/input.txt

A primeira linha do arquivo deve conter as coordenadas do ponto superior-direito da malha do planalto. As linhas seguintes devem representar as informações das sondas.
Cada sonda é representada por duas linhas. A primeira indica sua posição inicial e a segunda uma série de instruções indicando para a sonda como ela deverá explorar o planalto.

A posição é representada por dois inteiros e uma letra separados por espaços, correpondendo à coordenada X-Y e à direção da sonda. 

**Exemplo de arquivo:**
```
5 5
1 2 N
LMLMLMLMM
3 3 E
MMRMMRMRRM
```

### SAÍDA
Para cada sonda, na mesma ordem de entrada, o programa exibirá sua coordenada final e direção.
