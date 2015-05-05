# Senac-TCC 
### Bacharelado em ciência da computação

#### Alunos
 - [Lucas Teles](https://github.com/lucasteles)
 - [Rodrigo Mendonça](https://github.com/rodrigo-mendonca)

#### Orientador
 - [Eduardo Heredia](https://github.com/eheredia2511)

-----------

## Bibliografia Teorica
- [Fuzzy Network Profiling for Intrusion Detection](http://home.engineering.iastate.edu/~julied/publications/NAFIPSpaper2000.pdf)
- [An Improved Intrusion Detection based on Neural Network and Fuzzy Algorithm](http://ojs.academypublisher.com/index.php/jnw/article/download/jnw090512741280/9265)
- [Improving Network Attack Alarm System: A Proposed Hybrid Intrusion Detection System Model](http://www.computerscijournal.org/download/1613)
- [Artificial Intelligence in Network Intrusion Detection](http://www.fer.unizg.hr/_download/repository/KDI,_Miroslav_Stampar.pdf)
- [Current Studies On Intrusion Detection System, Genetic Algorithm And Fuzzy Logic](http://arxiv.org/ftp/arxiv/papers/1304/1304.3535.pdf)

## Ferramentas
- [GO](http://arxiv.org/ftp/arxiv/papers/1304/1304.3535.pdf)
- [WireShark](https://www.wireshark.org/)
- [PCAP](https://godoc.org/code.google.com/p/gopacket/pcap) 
- [GNS3](http://www.gns3.com/)
 

### Compilando arquivos LATEX
----------------------------
#### Dependencias
 - [TexLive](https://www.tug.org/texlive/)
 - [Node.JS](https://nodejs.org/)

#### Instale o gulpJS 
Rode no bash/zsh/cmd/PowerShell
```sh
npm install -g gulp
```

Para instalar as dependencias, na pasta `docs` rode
```sh
npm install
```

#### Iniciando monitor de compilação do latex
Rode na pasta docs/ onde se encontra o arquivo `gumpfile.js`

```sh
gulp
```

toda alteração no arquivo `.tex` na pasta tex ira gerar automaticamente o pdf na pasta pdf/

