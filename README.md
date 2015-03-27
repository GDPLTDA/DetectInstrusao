# Senac-TCC
-----------
Trabalho de conclusão de curso de ciência da computação.


## Bibliografia Teorica
- [Modelo HAND](http://modelohand.blogspot.com.br/2012/03/modelo-hand-na-previsao-de-enchentes_08.html)
- [Vulnerabilidades das Megacidades Brasileiras às Mudanças Climáticas: Sumario Executivo](http://mudancasclimaticas.cptec.inpe.br/~rmclima/pdfs/publicacoes/2010/SumarioExecutivo_megacidades.pdf)
- [Estudo aponta que enchentes e deslizamentos serão mais frequentes na capital paulista](http://www.inpe.br/noticias/noticia.php?Cod_Noticia=2215)
- [Vulnerabilidades Das Megacidades Brasileiras Às Mudanças Climáticas: Final](http://www.nepo.unicamp.br/textos/publicacoes/livros/megacidades/megacidades_RMSP.pdf)
- [Aplicação do modelo HAND](http://4ccr.pgr.mpf.mp.br/documentos-e-publicacoes/adis-propostas/adi_4901_peticao_inicial_-_parte_2.pdf)
- [Uso Multiplo da agua e multiplos conflitos em contextos urbanos](http://www.nepo.unicamp.br/textos/publicacoes/livros/migracao_urbanas/02pronex_14_Uso_Multiplo.pdf)
- [Entrevista Renato Tagnin - Gazeta](https://www.youtube.com/watch?v=fZJmAj2sFaE)


### Compilando arquivos LATEX
----------------------------
#### Dependencias
 - [TexLive](https://www.tug.org/texlive/)
 - [Node.JS](https://nodejs.org/)

#### Instale o gulpJS 
No bash/zsh/cmd/PowerShell
```sh
npm install -g gulp
```

Para isntalar as dependencias, na pasta `docs` rode
```sh
npm install
```

#### Iniciando monitor de compilação do latex
na pasta docs/ onde se encontra o arquivo `gumpfile.js`

```sh
gulp
```

toda alteração no arquivo `.tex` na pasta tex ira gerar altomaticamente o pdf na pasta pdf/

