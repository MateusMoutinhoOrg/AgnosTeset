## Novas rotas:

- /admin/liist-users
  lista os usuarios , permitindo parametros de filtragem, quantidade, paginacao 
   permissao: qualquer usuario do backoffice 

- /admin/root/* 
  middlware que garante que o usuario seja root 


- /admin/root/remove-user/{id_do_usuario}
  remove o usuario
   permissao: apenas usuarios roots 

- /admin/root/edit-user/{id_do_usuario}
  edita o usuario
   permissao: apenas usuarios roots 

- /admin/root/add-user
  adiciona um usuario
   permissao: apenas usuarios roots 



## Importante 
precisa implementar as interfaces para controlar essas rotas. 
