# 🚀 SpeedTest GO

Um testador de velocidade de internet estilo Speedtest desenvolvido em Go com armazenamento de resultados em JSON.

## 📋 Funcionalidades

- ✅ Teste de **Ping/Latência** - mede a resposta da conexão
- ✅ Teste de **Download** - testa velocidade de download
- ✅ Teste de **Upload** - testa velocidade de upload
- ✅ Armazenamento em **JSON** - histórico de resultados
- ✅ **Duas Versões**: Interface CLI (Go) e Web (HTML/JS)
- ✅ Interface amigável e visual - design moderno

## 🔧 Pré-requisitos

- Go 1.21 ou superior
- Conexão com a internet

## 📦 Instalação

1. Clone ou baixe este repositório
2. Instale as dependências:
```bash
go mod download
```

## 🎮 Como Usar

### 🖥️ Versão Go (CLI)

Execute o programa:
```bash
# Via .bat (Windows)
EXECUTAR-SPEEDTEST.bat

# Ou via terminal
& "C:\Program Files\Go\bin\go.exe" run velocidade.go
```

Ou compile e execute:
```bash
go build -o speedtest
./speedtest
```

### 🌐 Versão Web (Recomendado!)

**Não precisa instalar nada!**
- Abra `speedtest-web.html` no navegador
- Clique em "🚀 Iniciar Teste"
- Veja resultados em tempo real com interface bonita!

## 📊 Resultados

Os resultados são salvos automaticamente no arquivo `results.json` com o seguinte formato:

```json
{
  "results": [
    {
      "timestamp": "2024-01-15 14:30:00",
      "download_mbps": 45.23,
      "upload_mbps": 12.45,
      "ping_ms": 25.50
    }
  ]
}
```

## 📈 Histórico

O programa exibe automaticamente as últimas 5 medições realizadas.

## 🔍 Detalhes Técnicos

- Utiliza servidores Cloudflare para testes reais
- Download: 10MB de dados
- Upload: 5MB de dados
- Ping: Teste contra Google.com
- Cálculo em Mbps (Megabits por segundo)

## 📝 Estrutura do Projeto

```
go+json/
├── velocidade.go            # Código principal (Go)
├── speedtest-web.html       # Versão web (JavaScript)
├── EXECUTAR-SPEEDTEST.bat   # Script de execução (Windows)
├── go.mod                   # Gerenciamento de dependências
├── results.json             # Resultados dos testes (gerado)
├── README.md                # Documentação principal
├── COMO-EXECUTAR.md         # Guia de execução
├── COMO-USAR-SPEEDTEST-WEB.md # Guia versão web
└── CORRECAO-UPLOAD.md       # Notas sobre upload
```

## 🌟 Características

- 🎨 Interface amigável com emojis
- ⚡ Testes rápidos e eficientes
- 💾 Persistência de dados
- 📊 Visualização clara dos resultados
- 🔄 Histórico de medições

## 📄 Licença

Este projeto é de código aberto e está disponível para uso livre.

## 👨‍💻 Desenvolvido com

- Go 1.21+
- Cloudflare Speed Test API


"# speedtestprojetgo-json"  
# speedtestprojetgo-json
