# Relatório de Análise de Segurança — Backoffice AgnosTeset

**Data:** 2026-10-02
**Alvo:** branch `backoffice` — rotas HTML (`/admin`) e JSON (`/api/admin`), autenticação de backoffice, tokens de API e CLI.
**Escopo:** revisão de código + testes dinâmicos contra uma instância local (`start-server`), criada com um banco limpo e um usuário root de teste.
**Natureza:** teste de segurança autorizado sobre o próprio código do autor.

---

## Resumo executivo

A base aplica corretamente várias defesas importantes: **escape de saída** (sem XSS nos campos testados), **bloqueio de path traversal** no servidor de assets, **JWT restrito a HS256 com expiração obrigatória** (sem `alg:none`/confusão de algoritmo), **sessões aninhadas por usuário** (um JWT com `sub` trocado não alcança outra conta), e **autorização consistente** entre as rotas HTML e seus gêmeos JSON (sem credencial → 401; viewer em rota root → 403). Esses pontos estão descritos na seção "Pontos positivos".

Contudo, foram encontradas falhas que merecem correção. As de maior impacto são de **gestão de segredo e de credenciais**: o segredo do servidor trafega como argumento de linha de comando (visível a qualquer processo do host), as senhas usam **SHA-256 de passada única com um "salt" único e compartilhado**, e a **troca de senha não invalida sessões nem tokens** existentes. Somam-se a isso a **ausência de rate limiting** no login, o **bypass total de validação no comando CLI**, o controle de **IP burlável via `X-Forwarded-For`**, o cookie de sessão **sem o atributo `Secure`** e a **falta de cabeçalhos de segurança**.

### Quadro de severidade

| # | Severidade | Falha |
|---|------------|-------|
| 1 | **Alta** | Segredo do servidor passado como argumento de CLI (exposto em `ps`/histórico/logs) |
| 2 | **Alta** | Hash de senha fraco: SHA-256 de passada única com salt único e compartilhado |
| 3 | **Média** | Troca/reset de senha não invalida sessões nem tokens de API existentes |
| 4 | **Média** | `X-Forwarded-For` forjável anula a allowlist de IP de tokens e o vínculo de IP da sessão |
| 5 | **Média** | Ausência de rate limiting / bloqueio de conta no login e na API |
| 6 | **Média** | CLI `add-backoffice-user` ignora toda validação (unicidade, e-mail, senha) e sempre cria root |
| 7 | **Média** | Cookie de sessão sem `Secure`; sem HSTS; servidor em HTTP puro |
| 8 | **Baixa** | Ausência de cabeçalhos de segurança (CSP, X-Frame-Options, X-Content-Type-Options, etc.) |
| 9 | **Baixa** | Sem defesa-em-profundidade contra CSRF (depende só de `SameSite=Strict`) |
| 10 | **Baixa/Informativa** | Enumeração de usuários por diferença de tempo no login |

---

## Metodologia

1. Leitura das 23 declarações `route.yaml` e dos respectivos `InternalPureHandler.go`.
2. Leitura da lógica central: `backofficeauth`, `backofficetokens`, `backofficeusers`, `backofficeapi`, `render`, despacho gerado e adaptadores (`serverdeps`, `jwtdeps`, `hashdeps`, `randdeps`, `timedeps`) e o backend de armazenamento `Keep`.
3. Compilação (`go build ./cmd/main`) e execução local em `127.0.0.1:38080` com banco limpo.
4. Testes dinâmicos: acesso sem credencial a todas as rotas; variações de caminho/método; path traversal; login válido/inválido; autorização de viewer × root; tokens de API; spoofing de `X-Forwarded-For`; persistência de sessão/token após troca de senha; força bruta de login; XSS armazenado/refletido; CSRF; inspeção de cabeçalhos e cookies.

---

## Achados

### 1. [ALTA] Segredo do servidor passado como argumento de linha de comando

**Onde:** `sandbox/internal/commands/start_server/command.yaml` (flag `Secret`), `sandbox/internal/commands/start_server/InternalPureHandler.go:14`; `sandbox/internal/commands/add_backoffice_user/command.yaml` (flag `Secret`).

**Descrição:** o servidor é iniciado com `start-server --secret <s>` e usuários são criados com `add-backoffice-user --secret <s>`. Esse segredo é, ao mesmo tempo, **a chave de assinatura do JWT de sessão** e o **"salt" usado no hash das senhas** (`SHA256(secret + password)`).

**Evidência:** com o servidor rodando, o segredo aparece em texto claro na lista de processos:

```
$ ps aux | grep teste
... teste start-server --secret s3cr3t-test-key --addr 127.0.0.1:38080
```

Argumentos de CLI são legíveis por outros usuários do host via `/proc`/`ps`, e costumam ser capturados em histórico de shell, logs de processo, dumps de monitoramento e orquestradores de container.

**Impacto:** quem obtiver o segredo detém a chave que assina as sessões e o salt comum do hash de senhas; combinado com leitura dos hashes armazenados (arquivos do `maindatabase`), viabiliza ataque offline de senhas muito mais rápido. É a raiz de confiança de toda a autenticação exposta a um canal lateral trivial.

**Recomendação:** ler o segredo de variável de ambiente ou de arquivo com permissão restrita (ex.: `--secret-file`), nunca de argumento direto. Rotacionar o segredo (invalida sessões e exige re-hash de senhas — ver #2).

---

### 2. [ALTA] Hash de senha fraco: SHA-256 de passada única com salt único e compartilhado

**Onde:** `sandbox/internal/backofficeauth/backofficeauth.go:69-71`
```go
func PasswordSha(sandbox *api.Sandbox, password string) string {
    return sandbox.Deps.Hashdeps.Sha256Hex([]byte(sandbox.Config.Secret + password))
}
```

**Descrição:** as senhas são guardadas como um único SHA-256 de `secret + password`. Três problemas somados:
- **Função rápida:** SHA-256 de passada única permite bilhões de tentativas por segundo em GPU num cenário de vazamento dos hashes.
- **Sem salt por usuário:** o "salt" é o mesmo segredo global para todos. Senhas iguais geram hashes iguais (confirmado: dois usuários com a mesma senha produzem o mesmo `passwordsha`), o que vaza reuso de senha entre contas e permite uma única tabela de ataque para todos os usuários.
- **Não é um KDF:** ausência de fator de custo (bcrypt/scrypt/Argon2id).

**Evidência:** `maindatabase/backofficeuser/<id>/values/passwordsha` contém exatamente `SHA256(secret+senha)` (reproduzido com `shasum -a 256`). Os arquivos são gravados com permissão `0644`.

**Impacto:** vazamento do diretório do banco (backup, má configuração, acesso ao host) expõe senhas a quebra offline rápida; o salt compartilhado amplia o dano.

**Recomendação:** adotar Argon2id (ou bcrypt/scrypt) com **salt aleatório por usuário** armazenado junto ao hash. Migrar de forma incremental (re-hash no próximo login bem-sucedido).

---

### 3. [MÉDIA] Troca/reset de senha não invalida sessões nem tokens de API

**Onde:** `sandbox/internal/backofficeusers/backofficeusers.go:158-199` (`Update` — altera `passwordsha` sem encerrar sessões); comentário em `:153-157` confirma que as sessões permanecem abertas.

**Descrição:** ao editar a senha de um usuário (incluindo um reset feito por um root quando a conta foi comprometida), nenhuma sessão ativa nem token de API é revogado.

**Evidência (teste dinâmico):** após o root resetar a senha da usuária `carol`:
- a senha antiga deixa de autenticar (esperado);
- **o cookie de sessão antigo continua válido** → `GET /admin/home` = 200;
- **o token de API antigo continua válido** → `GET /api/admin/me` = 200;
- a sessão antiga ainda consegue **criar novos tokens**.

**Impacto:** o reset de senha — ação típica de resposta a incidente — não expulsa o atacante. Quem tiver capturado um cookie ou um token mantém acesso indefinidamente (tokens `never` não expiram).

**Recomendação:** ao alterar a senha, encerrar todas as sessões do usuário (remover os registros `sessions`) e, idealmente, revogar/renovar os tokens de API, ou oferecer um "sair de todos os dispositivos" explícito acionado pela troca de senha.

---

### 4. [MÉDIA] `X-Forwarded-For` forjável anula a allowlist de IP e o vínculo de IP da sessão

**Onde:** `adapters/libs/serverdeps/serverdeps.go:174-194` (`clientIp`). O próprio comentário reconhece: *"a client on its private network can forge X-Forwarded-For"*.

**Descrição:** quando a conexão chega de endereço **loopback ou privado** (cenário comum atrás de proxy reverso nginx/Docker, ou em acesso local), o IP do cliente é lido do **último** item de `X-Forwarded-For`. Em implantações sem proxy, ou com proxy que não sanitiza o cabeçalho, o cliente controla esse valor. Esse IP alimenta dois controles de segurança: a **allowlist de IP dos tokens de API** (`backofficetokens.accepts`) e o **vínculo de IP da sessão** (claim `ip` do JWT).

**Evidência (teste dinâmico):** token restrito ao IP `203.0.113.50` (que o cliente não possui):
- sem cabeçalho → `401` (correto);
- com `X-Forwarded-For: 203.0.113.50` → **`200`** (controle burlado);
- com `X-Forwarded-For: 1.1.1.1, 203.0.113.50` → **`200`** (vence o último item).

O mesmo vale para a sessão: um cookie emitido com `ip` forjado é aceito ao reenviar o mesmo `X-Forwarded-For`. (O cabeçalho `X-Client-Ip` enviado pelo cliente é corretamente ignorado — essa parte está certa.)

**Impacto:** a restrição de IP dos tokens e o vínculo de IP da sessão são, na prática, **não confiáveis** nesses cenários — são apresentados como controle de segurança, mas podem ser contornados por quem controla o cabeçalho.

**Recomendação:** não derivar IP confiável de `X-Forwarded-For` por padrão. Tornar a origem do IP configurável (contagem de proxies confiáveis, ou IP de conexão direto) e documentar que a allowlist só é efetiva atrás de um proxy que **reescreva** o cabeçalho. Tratar o vínculo de IP como reforço, não como fronteira.

---

### 5. [MÉDIA] Ausência de rate limiting / bloqueio de conta

**Onde:** `sandbox/internal/routeslist/admin/login/InternalPureHandler.go` e `.../api/admin/api_autentication/InternalPureHandler.go` — nenhum limite de tentativas.

**Descrição:** não há limitação de taxa nem bloqueio temporário após falhas de login, nem na validação de tokens.

**Evidência (teste dinâmico):** 300 tentativas de senha incorreta foram respondidas em ~1,8 s (todas `401`), sem `429` e sem bloqueio; a senha correta funcionou imediatamente depois.

**Impacto:** viabiliza força bruta online e *password spraying* contra o backoffice e adivinhação de tokens.

**Recomendação:** rate limiting por IP e por usuário, backoff progressivo e/ou bloqueio temporário após N falhas; considerar atraso artificial e monitoração de tentativas.

---

### 6. [MÉDIA] CLI `add-backoffice-user` ignora toda a validação e sempre cria root

**Onde:** `sandbox/internal/commands/add_backoffice_user/InternalPureHandler.go:13-29` — chama `db.AddBackofficeuser(...)` diretamente, sem passar por `backofficeusers.Add`; o `command.yaml` não tem flag de `role`, então `role` fica no default `0` (= **RoleRoot**).

**Descrição:** o comando de CLI não aplica nenhuma das validações que a camada web aplica: não verifica **unicidade** de username/e-mail, não valida **formato de e-mail** nem **tamanho mínimo de senha**, e cria sempre um usuário **root**.

**Evidência (teste dinâmico):**
- criado usuário `x` com e-mail `notanemail` e senha de 1 caractere — aceito;
- criados **dois** usuários com username `alice`. O login resolve por `FindByLogin` → `found[0]` (`backofficeauth.go:74-89`): a senha do `alice` **original** autentica; a do `alice` duplicado **nunca** autentica (fica "sombreado").

**Impacto:** quebra a invariante de unicidade que o restante do sistema assume, tornando a autenticação **dependente da ordem** de armazenamento; permite criar contas com credenciais fracas e múltiplos roots silenciosamente. (Mitigação: exige acesso ao host para rodar o CLI; ainda assim corrói uma garantia de autenticação.)

**Recomendação:** o CLI deve reutilizar `backofficeusers.Add` (mesmas validações) e expor uma flag `--role` com default seguro; recusar duplicatas.

---

### 7. [MÉDIA] Cookie de sessão sem `Secure`; sem HSTS; HTTP puro

**Onde:** `sandbox/internal/backofficeauth/backofficeauth.go:134-136` (`SessionCookie`).
```go
"%s=%s; Path=/; Max-Age=%d; HttpOnly; SameSite=Strict"
```

**Descrição:** o cookie `admin_token` define `HttpOnly` e `SameSite=Strict` (bom), mas **não** define `Secure`. O servidor fala HTTP puro e não emite HSTS.

**Evidência:** cabeçalho `Set-Cookie` observado no login não contém `Secure`.

**Impacto:** em qualquer acesso via HTTP (ou downgrade), o JWT de sessão trafega em claro e pode ser capturado na rede; sem `Secure`, o navegador também o envia em conexões não cifradas.

**Recomendação:** adicionar `Secure` ao cookie, servir atrás de TLS e enviar `Strict-Transport-Security`. Se houver modo de desenvolvimento em HTTP, tornar `Secure` condicional à implantação, com default seguro.

---

### 8. [BAIXA] Ausência de cabeçalhos de segurança

**Onde:** `sandbox/internal/render/render.go:11-19` (`Html`) define apenas `Content-Type` e `Cache-Control`.

**Descrição:** as páginas do backoffice não enviam `Content-Security-Policy`, `X-Frame-Options`/`frame-ancestors` (clickjacking), `X-Content-Type-Options: nosniff`, `Referrer-Policy` nem `Permissions-Policy`.

**Evidência:** inspeção dos cabeçalhos de `GET /admin/home` — presente apenas `Cache-Control: no-store`.

**Impacto:** reduz a defesa-em-profundidade. Sem `X-Frame-Options` o backoffice é enquadrável (clickjacking); sem CSP, qualquer XSS futuro fica sem contenção; sem `nosniff`, risco de MIME sniffing.

**Recomendação:** enviar um conjunto padrão de cabeçalhos de segurança em todas as respostas HTML (idealmente centralizado num middleware), incluindo CSP restritiva.

---

### 9. [BAIXA] Sem defesa-em-profundidade contra CSRF

**Onde:** formulários em `assets/templates/*.html` — nenhum token anti-CSRF; proteção depende exclusivamente de `SameSite=Strict`.

**Descrição:** as ações que alteram estado (`/admin/root/...`, revogação de token, logout) não têm token anti-CSRF nem verificação de `Origin`/`Referer`. O `SameSite=Strict` do cookie mitiga CSRF de navegador (e a API JSON não lê cookie, logo não é alvo de CSRF), mas não há segunda camada. *(A requisição com `Origin` hostil "passou" apenas porque `curl` não é navegador e envia o cookie de qualquer forma; um navegador real não enviaria o cookie cross-site com `SameSite=Strict`.)*

**Impacto:** baixo hoje; porém, se `SameSite` for enfraquecido (navegadores antigos, ou mudança futura para `Lax`), as ações de alteração ficam sem proteção.

**Recomendação:** adicionar verificação de `Origin`/`Referer` nas rotas que alteram estado e/ou token anti-CSRF por sessão, como reforço.

---

### 10. [BAIXA/INFORMATIVA] Enumeração de usuários por tempo de resposta

**Onde:** `sandbox/internal/backofficeauth/backofficeauth.go:93-102` (`Authenticate`) — o hash só é calculado quando o usuário existe.

**Descrição:** como o cálculo do hash (e os dois scans de `FindByLogin`) só ocorrem para logins existentes, há diferença mensurável de tempo entre usuário existente e inexistente. A mensagem de erro é genérica (bom), mas o tempo não é constante.

**Evidência:** sobre 500 amostras, diferença de mediana da ordem de centenas de microssegundos entre `root` (existente) e um nome inexistente. É pequena e ruidosa em rede, mas explorável localmente/em larga escala.

**Impacto:** permite inferir quais usernames/e-mails existem. Baixo isoladamente; combina-se com a falta de rate limiting (#5).

**Recomendação:** executar sempre uma comparação de hash "dummy" com custo equivalente quando o usuário não existe, para uniformizar o tempo.

---

## Pontos positivos (defesas corretas observadas)

- **XSS:** saída escapada via `text/template` com `{{html ...}}` e `{{urlquery ...}}`. Payloads armazenados (username, e-mail, nome de token) e refletidos (busca, username no login) apareceram **somente em forma escapada** — nenhum executou.
- **Path traversal:** `frontio.SafePath` (`sandbox/internal/generated/frontio/frontio.go:70-86`) rejeita `.`, `..`, barra invertida e NUL, inclusive em variações percent-encoded — todas resultaram em `404`.
- **JWT robusto:** `jwtdeps` força `HS256` via `WithValidMethods` e exige `exp` (`WithExpirationRequired`) — sem `alg:none`, sem confusão de algoritmo, sem token sem expiração (`adapters/libs/jwtdeps/jwtdeps.go:41-48`).
- **Isolamento de sessão por usuário:** o `jti` referencia um registro `sessions` **aninhado sob o usuário**; trocar o `sub` de um JWT não alcança a sessão de outra conta (`backofficeauth.go:174-219`).
- **Autorização consistente:** sem credencial → `401`; viewer em rota `/admin/root/*` e `/api/admin/root/*` → `403`; verificado nas rotas HTML e nos gêmeos JSON. A API **não lê o cookie**, eliminando CSRF na superfície JSON.
- **Invariante de root:** auto-remoção bloqueada e último root não pode ser rebaixado → sempre resta ≥ 1 root (`backofficeusers.go` `Remove`/`Update`).
- **Tokens de API:** mostrados uma única vez; armazenado apenas o `SHA-256` (campo indexado `tokensha`); gerados com 32 bytes de `crypto/rand`; hash de senha nunca retornado por API nem por template.
- **Limites de corpo:** `max-bytes` (413) e checagem de `Content-Type` (415) aplicados antes de ler o corpo.

---

## Recomendações priorizadas

1. **Parar de passar o segredo por argumento de CLI** (env/arquivo com permissão restrita) — #1.
2. **Trocar o hash de senha para Argon2id/bcrypt com salt por usuário** e migrar incrementalmente — #2.
3. **Invalidar sessões e tokens ao trocar/resetar senha** (e oferecer "sair de todos os dispositivos") — #3.
4. **Adicionar rate limiting/bloqueio** no login e na validação de tokens — #5.
5. **Fazer o CLI reutilizar `backofficeusers.Add`** e expor `--role` com default seguro — #6.
6. **Marcar o cookie como `Secure`, servir sob TLS, enviar HSTS** — #7.
7. **Tratar o IP de `X-Forwarded-For` como não confiável por padrão**; documentar a allowlist como efetiva só atrás de proxy que reescreva o cabeçalho — #4.
8. **Enviar cabeçalhos de segurança** (CSP, X-Frame-Options, nosniff, Referrer-Policy) e **reforço anti-CSRF** (Origin/Referer) — #8, #9.
9. **Uniformizar o tempo de resposta do login** — #10.

---

## Notas de reprodução

- Ambiente de teste: binário compilado em sandbox, banco limpo, root de teste criado via `add-backoffice-user`, servidor em `127.0.0.1:38080`.
- Nenhum dado ou sistema de produção foi tocado. O diretório `maindatabase` de teste e os scripts auxiliares ficaram fora da árvore do projeto (no scratchpad da sessão).
- Os achados #1, #4, #5, #6 foram confirmados tanto por leitura de código quanto por teste dinâmico; #2, #3, #7, #8, #9 por leitura de código e inspeção de respostas; #10 por medição (sinal pequeno — classificado como baixo/informativo).
