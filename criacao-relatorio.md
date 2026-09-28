### objetivo:
Use o agnos cli, e crie diversos projetos com ele, desde servidores, clis, etc .. entao crie um relatorio apontando tudo que deve mudar no projeto para se aproximar dos vetores guia.

### Vetores Guia 
- o agnos deve ser projetado para ser usado primariamente por LLMs.
- os projetos criados pelo agnos devem ser compreensiveis por LLMs
- os projetos criados devem ser modulares (por isso o sistema de sandbox + dependencias), de modo que qualquer llm conssiga entender toda a api.
- a documentacao publica deve ser facil de ser entendia por humanos. 
- a documentacao privada deve ser focada em llms.
- os projetos criados pelo agnos devem ter conssistencia, de modo que apenas usando o agnos, uma empresa conssiga criar um ecossistema inteiro (importando um projeto agnos dentro de outro projeto agnos).
- Apesar de nao ser a prioridade 1, o consumo de tokens deve ser levado em conta, entao verbosidade excessiva e pessimo.

## Pontos que eu achei ruins, mas posso estar errado
- Acho que os projetos gerados estao excessivamente verbosos, tipo, qualquer coisa que eu crio , gera +100 arquivos
- a documentacao gerada esta excessivamente redundante, tem muita coisa na doc, que tambem tem na cli do proprio agnos.
- a documentacao de rotas e commands, esta confusa.(essa deveria ser simples de entender).


## O Relatorio
- salve o relatorio em relatorio.md 
- escreva o relatorio de maneira consisa e use index.
- escreva tudo que tem que ser adicionado, modificado e removido.
- caso encontre bugs, reporte esses bugs no relatorio.
- aponte os nomes de funcoes, struct,s docs ,que acha que deveriam ter nomes diferentes.

