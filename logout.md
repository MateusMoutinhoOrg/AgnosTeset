
- adicione uma subtabela chamada hosts na tabela de backofficeuser , essa tabela tera as seguintes colunas~~
  - host string 
  - mincreation time.time

- Adicione a propriedade  Host ao token jtw 

- Na autenticacao, alem de verificar o token, verifique se o host bate. (impedimento de cross site request)
 e se o tokenjtw foi criado apos o mintimecreation do host

- adicione a rota /admin/logout, essa rota ira buscar o host, e adicionar o mincreation com o tempo unix atual, de modo que invalide
qualquer token criado antes do mincreation atual.

- adicione na home um botao de logunt.