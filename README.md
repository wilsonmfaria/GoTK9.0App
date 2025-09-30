# GoTK9.0App

Mocking de uma ferramenta de Backup em Go utilizando a biblioteca **[tk9.0](https://pkg.go.dev/modernc.org/tk9.0)** para construção de interfaces gráficas.

Este projeto complementa o guia **Tk9.0 em Go**, trazendo um exemplo completo com:
- **Concorrência com goroutines e channels**
- **Atualização de UI com `NewTicker`**
- **Barras de progresso (`Progressbar`) personalizadas**
- **Confirmação de ações com `MessageBox`**
- **Estilização de widgets (`StyleConfigure`)**
- **Layout organizado via `Grid`**
- **Validação de dados de entrada**

---

## 🚀 Funcionalidades

- Criar **novo backup** a partir de um diretório local.
- Restaurar **backup existente** do servidor.
- Barra de progresso determinística, atualizada em tempo real.
- Feedback visual com rótulos dinâmicos e animação de status (`...`).
- Botões que desabilitam/reabilitam automaticamente durante o processo.
- Diálogo de confirmação antes de iniciar tarefas longas.

---

## 📦 Instalação

Pré-requisitos:
- **Go 1.21+** instalado.
- Tcl/Tk instalado no sistema (em distros Linux pode ser necessário instalar via gerenciador de pacotes).

Clone o repositório e baixe as dependências:

```bash
git clone https://github.com/wilsonmfaria/GoTK9.0App
cd GoTK9.0App
go get modernc.org/tk9.0
```

---

## ▶️ Execução

Para rodar o aplicativo:

```bash
go run .
```

A janela principal será aberta com duas opções:

1. **Criar Backup**
2. **Restaurar Backup**

---

## 📚 Recursos para estudo

* **Documentação oficial da API:**
  👉 [pkg.go.dev/modernc.org/tk9.0](https://pkg.go.dev/modernc.org/tk9.0)

* **Exemplos oficiais do projeto:**
  👉 [GitLab tk9.0 / _examples](https://gitlab.com/cznic/tk9.0/-/tree/v1.72.0/_examples)

---

## ✨ Autor

Desenvolvido por **Wilson M. Faria**
📘 Parte do material da apostila: *Tk9.0 em Go: Guia Essencial*.

---

## 📜 Licença

Este projeto é disponibilizado para fins educacionais.
Sinta-se à vontade para estudar, modificar e contribuir.