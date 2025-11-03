# 🚀 Guia de Instalação e Teste - SpeedTest GO

## ⚠️ ATENÇÃO: Go Não Encontrado

O Go não está instalado no seu sistema. Siga as instruções abaixo para instalar e testar o SpeedTest.

---

## 📥 PASSO 1: Instalar o Go

### Windows:

1. **Baixe o Go:**
   - Acesse: https://golang.org/dl/
   - Baixe o instalador: **go1.21.x.windows-amd64.msi** (ou versão mais recente)

2. **Instale:**
   - Execute o instalador `.msi`
   - Siga o assistente (Next, Next, Install)
   - A instalação padrão será em: `C:\Program Files\Go`

3. **Verifique a instalação:**
   ```cmd
   go version
   ```
   
   Deve mostrar: `go version go1.21.x windows/amd64`

---

## 🎯 PASSO 2: Testar o SpeedTest

### Abra um novo terminal (CMD ou PowerShell):

**Opção 1 - Via Terminal (Simples):**
```cmd
cd E:\MEU-REPOSITORIO-LOCAL-GIT-GITHUB\go+json
go run velocidade.go
```

**Opção 2 - Compilar e Executar:**
```cmd
cd E:\MEU-REPOSITORIO-LOCAL-GIT-GITHUB\go+json
go build -o speedtest.exe velocidade.go
speedtest.exe
```

**Opção 3 - PowerShell:**
```powershell
cd E:\MEU-REPOSITORIO-LOCAL-GIT-GITHUB\go+json
go run velocidade.go
```

---

## 📊 O Que Você Verá:

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
⏱️  Ping:        35.20 ms
⬇️  Download:    85.45 Mbps
⬆️  Upload:      23.67 Mbps
═══════════════════════════════════════

✅ Teste salvo em: results.json

📈 Últimas 5 medições:
1. [2024-01-15 10:00:00] ⬇️ 85.45 Mbps | ⬆️ 23.67 Mbps | ⏱️ 35.20 ms
```

---

## ✅ PASSO 3: Verificar Resultados

Após executar, você terá criado o arquivo `results.json`:

```json
{
  "results": [
    {
      "timestamp": "2024-01-15 14:30:00",
      "download_mbps": 85.45,
      "upload_mbps": 23.67,
      "ping_ms": 35.20
    }
  ]
}
```

Você pode abrir este arquivo em qualquer editor de texto!

---

## 🐛 Resolução de Problemas

### Problema: "go: command not found"

**Solução:**
- Instale o Go seguindo o PASSO 1
- Feche e reabra o terminal após instalar
- Verifique com: `go version`

### Problema: "Go não está no PATH"

**Solução:**
- O instalador adiciona automaticamente ao PATH
- Se não funcionar, adicione manualmente:
  - PATH do sistema: `C:\Program Files\Go\bin`

### Problema: "Erro ao conectar durante teste"

**Solução:**
- Verifique sua conexão com internet
- Verifique firewall/proxy
- Tente desabilitar temporariamente VPN

### Problema: "cannot find package"

**Solução:**
```cmd
cd E:\MEU-REPOSITORIO-LOCAL-GIT-GITHUB\go+json
go mod tidy
go run velocidade.go
```

---

## 📋 Checklist Rápido

Antes de executar, certifique-se:

- [ ] Go está instalado (`go version` funciona)
- [ ] Terminal está na pasta correta do projeto
- [ ] Você tem conexão com internet
- [ ] Arquivo `velocidade.go` existe na pasta

---

## 🎯 Próximos Passos Após Testar

1. ✅ Execute múltiplas vezes para ver histórico
2. ✅ Compare velocidades ao longo do dia
3. ✅ Compartilhe resultados com sua equipe
4. ✅ Explore o código em `velocidade.go`

---

## 📚 Recursos Úteis

- **Documentação Go:** https://golang.org/doc/
- **Tutorial Go:** https://golang.org/doc/tutorial/
- **Cloudflare Speed Test:** https://speed.cloudflare.com/

---

## 💡 Dica Pro

Execute o teste em diferentes horários do dia para ver como sua velocidade varia!

```cmd
# Manhã
go run velocidade.go

# Tarde
go run velocidade.go

# Noite
go run velocidade.go
```

