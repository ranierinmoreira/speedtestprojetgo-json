# ⚡ Guia Rápido - SpeedTest GO

## 🚀 Como Executar

### Windows:

```cmd
go run velocidade.go
```

### Compilar (Windows):

```cmd
go build -o speedtest.exe velocidade.go
speedtest.exe
```

### Linux/Mac:

```bash
go run velocidade.go
```

Ou compile:

```bash
go build -o speedtest velocidade.go
./speedtest
```

## 📋 O que o programa faz:

1. ✅ **Ping** - Testa latência contra Google.com
2. ✅ **Download** - Faz download de 10MB da Cloudflare
3. ✅ **Upload** - Faz upload de 5MB para Cloudflare
4. ✅ **Salva** - Grava resultados em `results.json`
5. ✅ **Exibe** - Mostra últimas 5 medições

## 📊 Saída Exemplo:

```
╔════════════════════════════════════════╗
║     🚀 SpeedTest GO v1.0 🚀           ║
╚════════════════════════════════════════╝

🔍 Testando Ping (Latência)...
⬇️  Testando Velocidade de Download...
⬆️  Testando Velocidade de Upload...

═══════════════════════════════════════
📊 RESULTADOS DO TESTE
═══════════════════════════════════════
⏱️  Ping:        25.50 ms
⬇️  Download:    45.23 Mbps
⬆️  Upload:      12.45 Mbps
═══════════════════════════════════════

✅ Teste salvo em: results.json

📈 Últimas 5 medições:
1. [2024-01-15 10:00:00] ⬇️ 45.23 Mbps | ⬆️ 12.45 Mbps | ⏱️ 25.50 ms
2. [2024-01-15 11:30:00] ⬇️ 48.12 Mbps | ⬆️ 13.20 Mbps | ⏱️ 22.30 ms
```

## 🔍 Requisitos:

- ✅ Go instalado (versão 1.21 ou superior)
- ✅ Conexão com internet
- ✅ Firewall permitindo conexões HTTP/HTTPS

## 🐛 Solução de Problemas:

**Erro: "cannot find package"**
```bash
go mod tidy
```

**Erro: "command not found: go"**
- Instale o Go: https://golang.org/dl/

**Erro de conexão durante testes**
- Verifique sua conexão com internet
- Firewall corporativo pode estar bloqueando

## 📝 Arquivos Criados:

- `results.json` - Histórico de testes (gerado automaticamente)
- `go.mod` - Configuração do módulo Go

## 🎯 Próximos Passos:

1. Execute o programa: `go run velocidade.go`
2. Analise os resultados
3. Compare medições ao longo do tempo
4. Verifique `results.json` para histórico completo


