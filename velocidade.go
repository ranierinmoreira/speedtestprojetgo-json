package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// TestResult representa o resultado de um teste de velocidade
type TestResult struct {
	Timestamp string  `json:"timestamp"`
	Download  float64 `json:"download_mbps"`
	Upload    float64 `json:"upload_mbps"`
	Ping      float64 `json:"ping_ms"`
}

// SpeedTest armazena todos os resultados
type SpeedTest struct {
	Results []TestResult `json:"results"`
}

// Servidores de teste (URLs públicas para download/upload)
var (
	downloadURL = "https://speed.cloudflare.com/__down?bytes=%d"
	uploadURL   = "https://speed.cloudflare.com/__up"
)

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║     🚀 SpeedTest GO v1.0 🚀           ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	// Carregar resultados anteriores
	speedTest := loadResults()

	// Executar testes
	fmt.Println("🔍 Testando Ping (Latência)...")
	ping := testPing()

	fmt.Println("⬇️  Testando Velocidade de Download...")
	download := testDownload()

	fmt.Println("⬆️  Testando Velocidade de Upload...")
	upload := testUpload()

	// Criar resultado
	result := TestResult{
		Timestamp: time.Now().Format("2006-01-02 15:04:05"),
		Download:  download,
		Upload:    upload,
		Ping:      ping,
	}

	// Adicionar ao histórico
	speedTest.Results = append(speedTest.Results, result)

	// Salvar resultados
	saveResults(speedTest)

	// Exibir resultados
	fmt.Println()
	fmt.Println("═══════════════════════════════════════")
	fmt.Println("📊 RESULTADOS DO TESTE")
	fmt.Println("═══════════════════════════════════════")
	fmt.Printf("⏱️  Ping:        %.2f ms\n", result.Ping)
	fmt.Printf("⬇️  Download:    %.2f Mbps\n", result.Download)
	fmt.Printf("⬆️  Upload:      %.2f Mbps\n", result.Upload)
	fmt.Println("═══════════════════════════════════════")
	fmt.Println()
	fmt.Printf("✅ Teste salvo em: %s\n", "results.json")
	
	// Exibir histórico
	fmt.Println()
	fmt.Println("📈 Últimas 5 medições:")
	displayHistory(speedTest.Results)
}

// testPing testa a latência da conexão
func testPing() float64 {
	start := time.Now()
	
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	
	resp, err := client.Get("https://www.google.com")
	if err != nil {
		fmt.Printf("❌ Erro no teste de ping: %v\n", err)
		return 0
	}
	defer resp.Body.Close()
	
	elapsed := time.Since(start)
	
	return float64(elapsed.Milliseconds())
}

// testDownload testa a velocidade de download
func testDownload() float64 {
	// Teste com 10MB
	testSize := 10 * 1024 * 1024
	url := fmt.Sprintf(downloadURL, testSize)
	
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	
	start := time.Now()
	
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("❌ Erro no teste de download: %v\n", err)
		return 0
	}
	defer resp.Body.Close()
	
	// Ler dados para medir velocidade real
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		fmt.Printf("❌ Erro ao ler dados: %v\n", err)
		return 0
	}
	
	elapsed := time.Since(start)
	
	// Calcular Mbps (Megabits por segundo)
	bitsDownloaded := float64(testSize) * 8
	megabitsDownloaded := bitsDownloaded / 1000000
	seconds := elapsed.Seconds()
	
	mbps := megabitsDownloaded / seconds
	
	return mbps
}

// testUpload testa a velocidade de upload
func testUpload() float64 {
	// Gerar dados de teste (5MB)
	testSize := 5 * 1024 * 1024
	testData := make([]byte, testSize)
	
	// Preencher com dados aleatórios
	for i := range testData {
		testData[i] = byte(i % 256)
	}
	
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	
	start := time.Now()
	
	// Criar POST request
	resp, err := client.Post(uploadURL, "application/octet-stream", 
		io.NopCloser(bytes.NewReader(testData)))
	if err != nil {
		fmt.Printf("❌ Erro no teste de upload: %v\n", err)
		return 0
	}
	defer resp.Body.Close()
	
	_, err = io.Copy(io.Discard, resp.Body)
	if err != nil {
		fmt.Printf("❌ Erro ao ler resposta: %v\n", err)
		return 0
	}
	
	elapsed := time.Since(start)
	
	// Calcular Mbps
	bitsUploaded := float64(testSize) * 8
	megabitsUploaded := bitsUploaded / 1000000
	seconds := elapsed.Seconds()
	
	mbps := megabitsUploaded / seconds
	
	return mbps
}

// loadResults carrega resultados anteriores do arquivo JSON
func loadResults() SpeedTest {
	var speedTest SpeedTest
	
	file, err := os.Open("results.json")
	if err != nil {
		// Se o arquivo não existe, retorna estrutura vazia
		return speedTest
	}
	defer file.Close()
	
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&speedTest); err != nil {
		fmt.Printf("⚠️  Aviso: Não foi possível ler results.json: %v\n", err)
		return SpeedTest{Results: []TestResult{}}
	}
	
	return speedTest
}

// saveResults salva os resultados no arquivo JSON
func saveResults(speedTest SpeedTest) {
	file, err := os.Create("results.json")
	if err != nil {
		fmt.Printf("❌ Erro ao criar arquivo: %v\n", err)
		return
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(speedTest); err != nil {
		fmt.Printf("❌ Erro ao salvar resultados: %v\n", err)
	}
}

// displayHistory exibe as últimas medições
func displayHistory(results []TestResult) {
	count := len(results)
	if count == 0 {
		fmt.Println("Nenhuma medição anterior encontrada.")
		return
	}
	
	// Mostrar últimas 5
	start := 0
	if count > 5 {
		start = count - 5
	}
	
	for i := start; i < count; i++ {
		r := results[i]
		fmt.Printf("%d. [%s] ⬇️ %.2f Mbps | ⬆️ %.2f Mbps | ⏱️ %.2f ms\n",
			i+1, r.Timestamp, r.Download, r.Upload, r.Ping)
	}
}
