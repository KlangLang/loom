package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type BackendInfo struct {
	Name        string
	Command     string
	VersionFlag string
	Installed   bool
	Version     string
}

func detectBackends(log Log) map[string]BackendInfo {
	backends := map[string]BackendInfo{
		"java": {
			Name:        "Java",
			Command:     "java",
			VersionFlag: "-version",
		},
		"c": {
			Name:        "C (GCC)",
			Command:     "gcc",
			VersionFlag: "--version",
		},
		"python": {
			Name:        "Python",
			Command:     "python3",
			VersionFlag: "--version",
		},
		"rust": {
			Name:        "Rust",
			Command:     "rustc",
			VersionFlag: "--version",
		},
	}

	fmt.Printf("%s◉%s Detecting installed backends...\n", log.PRIMARY_COLOR, log.RESET_COLOR)

	for key, backend := range backends {
		cmd := exec.Command(backend.Command, backend.VersionFlag)
		output, err := cmd.CombinedOutput()

		if err == nil {
			backend.Installed = true
			// Pega a primeira linha da versão
			lines := strings.Split(string(output), "\n")
			if len(lines) > 0 {
				backend.Version = strings.TrimSpace(lines[0])
			}
			fmt.Printf("  %s✔%s %s detected: %s\n", 
				log.SUCESS_COLOR, log.RESET_COLOR, backend.Name, backend.Version)
		} else {
			backend.Installed = false
			fmt.Printf("  %s✖%s %s not found\n", 
				log.ERROR_COLOR, log.RESET_COLOR, backend.Name)
		}

		backends[key] = backend
	}

	return backends
}

func selectDefaultBackend(backends map[string]BackendInfo, log Log) string {
	installedBackends := []string{}
	
	for key, backend := range backends {
		if backend.Installed {
			installedBackends = append(installedBackends, key)
		}
	}

	if len(installedBackends) == 0 {
		fmt.Printf("\n%s⚠ Warning:%s No backends detected. Using Java as default.\n", 
			log.WARNING_COLOR, log.RESET_COLOR)
		return "java"
	}

	if len(installedBackends) == 1 {
		fmt.Printf("\n%s◉%s Only one backend found, using %s as default.\n", 
			log.PRIMARY_COLOR, log.RESET_COLOR, installedBackends[0])
		return installedBackends[0]
	}

	fmt.Printf("\n%s◉%s Multiple backends available. Choose default:\n", 
		log.PRIMARY_COLOR, log.RESET_COLOR)
	
	for i, backend := range installedBackends {
		fmt.Printf("  %s[%d]%s %s\n", log.ACCENT_COLOR, i+1, log.RESET_COLOR, backends[backend].Name)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nSelect [1]: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return installedBackends[0]
	}

	var selection int
	fmt.Sscanf(input, "%d", &selection)

	if selection > 0 && selection <= len(installedBackends) {
		return installedBackends[selection-1]
	}

	return installedBackends[0]
}

func generateBackendConfig(backends map[string]BackendInfo, defaultBackend string) string {
	config := "[targets]\n"
	
	order := []string{"java", "c", "cpp", "python", "rust", "go", "node"}
	
	for _, key := range order {
		backend, exists := backends[key]
		if !exists {
			continue
		}
		
		padding := strings.Repeat(" ", 7-len(key))
		config += fmt.Sprintf("%s%s = %v\n", key, padding, backend.Installed)
	}

	return config
}

func newCommand() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: loom new <project_name>")
		return
	}

	log := NewLog()
	name := os.Args[2]

	log.Header()
	log.finalizeBottomheader()
	fmt.Println()
	log.Line()

	backends := detectBackends(log)
	fmt.Println()

	defaultBackend := selectDefaultBackend(backends, log)
	fmt.Printf("%s◉%s Default backend set to: %s\n", 
		log.PRIMARY_COLOR, log.RESET_COLOR, defaultBackend)
	fmt.Println()

	readme := question("Do you want to create a base README for " + name)

	author := askAuthor(log)

	fmt.Printf("%s◉%s Ok, creating project structure...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	createFoldersTree(name, log)

	fmt.Printf("%s◉%s Almost there, writing manifest...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	createManifest(name, log, readme, backends, defaultBackend, author)

	fmt.Printf("%s◉%s Initializing Klang module...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	fmt.Printf("%s◈%s Finalizing...\n", log.ACCENT_COLOR, log.RESET_COLOR)

	fmt.Printf("\n%s✔%s Project '%s' created.\n",
		log.SUCESS_COLOR, log.RESET_COLOR, name)
	
	fmt.Printf("\n%s◉%s Next steps:%s\n", log.PRIMARY_COLOR, log.RESET_COLOR, log.RESET_COLOR)
	fmt.Printf("  cd %s\n", name)
	fmt.Printf("  kc run src/main.k\n")
	
	log.Line()
}

func createManifest(name string, log Log, readme bool, backends map[string]BackendInfo, defaultBackend string, author string) {
	SRC := filepath.Join(name, "src")

	backendConfig := generateBackendConfig(backends, defaultBackend)

	tomlContent := fmt.Sprintf(`[project]
name = "%s"
version = "0.1.0"
author = "%s"

[build]
default-backend = "%s"

%s`, name, author, defaultBackend, backendConfig)

	readmeContent := fmt.Sprintf(`# %s

Projeto criado com o Loom — o CLI oficial da linguagem Klang.

## Rodar

kc run src/main.k

## Compilar

kc build

## Backends disponíveis

`, name)

	for _, backend := range backends {
		if backend.Installed {
			readmeContent += fmt.Sprintf("- **%s**: %s\n", backend.Name, backend.Version)
		}
	}

	files := []string{
		filepath.Join(SRC, "main.k"),
		filepath.Join(SRC, "main_test.k"),
		filepath.Join(name, "loom.toml"),
		filepath.Join(name, "README.md"),
	}

	filesContents := [][]byte{
		[]byte("@Use(\"" + defaultBackend + "\")\npublic static void main(String[] args){\n\tprintln(\"Hello from Klang!\");\n}"),
		[]byte("test \"project boots\" {\n\tassert(true);\n}"),
		[]byte(tomlContent),
		[]byte(readmeContent),
	}

	for i := 0; i < len(files); i++ {
		file := files[i]
		content := filesContents[i]

		if strings.HasSuffix(file, "README.md") && !readme {
			continue
		}

		if err := makeFile(content, file); err != nil {
			fmt.Printf("%s✖ Error:%s %s\n", log.ERROR_COLOR, log.RESET_COLOR, err)
			return
		}
	}
}

func createFoldersTree(name string, log Log) {
	BASE_FOLDER := name

	SRC := filepath.Join(BASE_FOLDER, "src")
	TESTS := filepath.Join(BASE_FOLDER, "tests")
	BUILD := filepath.Join(BASE_FOLDER, "build")
	BUILD_CACHE := filepath.Join(BUILD, "cache")
	BUILD_BACK := filepath.Join(BUILD, "backends")
	KLANG_CACHES := filepath.Join(BASE_FOLDER, ".klang")

	folders := []string{SRC, TESTS, BUILD, BUILD_CACHE, BUILD_BACK, KLANG_CACHES}

	fmt.Printf("%s◉%s Creating folder tree...\n", log.PRIMARY_COLOR, log.RESET_COLOR)
	for _, folder := range folders {
		err := os.MkdirAll(folder, 0755)

		if err != nil {
			fmt.Printf("%s✖ Error:%s %s\n", log.ERROR_COLOR, log.RESET_COLOR, err)
			return
		}
	}
}

var respostasSim = map[string]bool{
	// Portuguese
	"sim": true, "s": true, "si": true, "claro": true, "certeza": true,
	"com certeza": true, "certo": true, "correto": true, "exato": true,
	"ok": true, "aceito": true, "bora": true, "é isso": true, "isso": true,
	"beleza": true, "positivo": true, "afirmativo": true, "confirmo": true, "pode": true,

	// English
	"yes": true, "y": true, "yeah": true, "yep": true, "yup": true,
	"sure": true, "okay": true, "right": true, "correct": true,
	"absolutely": true, "definitely": true, "affirmative": true,
	"aye": true, "indeed": true, "certainly": true, "roger": true,

	// Spanish
	"sí": true, "vale": true, "bueno": true, "por supuesto": true,
	"desde luego": true, "correcto": true, "exacto": true,
	"de acuerdo": true, "está bien": true, "okey": true, "dale": true, "va": true,

	// Other languages...
	"oui": true, "ja": true, "da": true, "はい": true, "hai": true,
	"是": true, "네": true, "tak": true, "ano": true, "evet": true,
	
	// Universals
	"1": true, "true": true, "+": true, "✓": true, "✔": true,
}

func question(quest string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(quest + " (s/n) [default = n]: ")

	input, _ := reader.ReadString('\n')
	input = strings.ToLower(strings.TrimSpace(input))

	return respostasSim[input]
}

func askAuthor(log Log) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("%s◉%s Author name [default = ~K']: ", log.PRIMARY_COLOR, log.RESET_COLOR)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	if input == "" {
		return "foreigner"
	}

	return input
}