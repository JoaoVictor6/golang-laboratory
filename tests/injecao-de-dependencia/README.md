# Resumo

Nossa primeira rodada de código não foi fácil de testar porque escrevemos dados em algum lugar que não podíamos controlar.

Graças aos nossos testes, refatoramos o código para que pudéssemos controlar para onde os dados eram escritos **injetando uma dependência** que nos permitiu:

- **Testar nosso código**: se você não consegue testar uma função de forma simples, geralmente é porque dependências estão acopladas em uma função ou estado global. Se você tem um pool de conexão global da base de dados, por exemplo, é provável que seja difícil testar e vai ser lento para ser execudado. A injeção de dependência te motiva a injetar em uma dependência da base de dados (através de uma interface), para que você possa criar um mock com algo que você possa controlar nos seus testes.
- Separar nossas preocupações, desacoplando onde os dados vão de como gerá-los. Se você já achou que um método/função tem responsabilidades demais (gerando dados e escrevendo na base de dados? Lidando com requisições HTTP e aplicando lógica a nível de domínio?), a injeção de dependência provavelmente será a ferramenta que você precisa.
- **Permitir** que nosso código seja reutilizado em contextos diferentes: o primeiro contexto "novo" do nosso código pode ser usado dentro dos testes. No entanto, se alguém quiser testar algo novo com nossa função, a pessoa pode injetar suas próprias dependências.

## A biblioteca padrão do Go é muito boa, leve um tempo para estudá-la
Ao termos familiaridade com a interface io.Writer, somos capazes de usar bytes.Buffer no nosso teste como nosso Writer para que depois possamos usar outros Writer da biblioteca padrão para usar na nossa função em uma aplicação de linha de comando ou em um servidor web.

Quanto mais familiar você for com a biblioteca padrão, mais vai ver essas interfaces de propósito geral que você pode reutilizar no seu próprio código para tornar o software reutilizável em vários contextos diferentes.