package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

const klangBinPathLine = `export PATH="$HOME/.klang/bin:$PATH"`
const klangBinSubstr = ".klang/bin"

func installCommand() {
	contentToInstall := []string{}

	if len(os.Args) <= 2 {
		contentToInstall = append(contentToInstall, "klang")
	} else {
		contentToInstall = os.Args[2:]
	}

	for _, item := range contentToInstall {
		if item != "klang" {
			fmt.Printf("Warning: loom does not support %s yet. Installing klang only.\n", item)
		}
	}

	shellConfigPath, err := determineShellConfigPath()
	if err != nil {
		fmt.Printf("Error: Could not determine shell config file: %v\n", err)
		return
	}
	fmt.Printf("Shell determined. Editing file: %s\n", shellConfigPath)

	currentUser, err := user.Current()
	if err != nil {
		fmt.Printf("Error getting user home path: %v\n", err)
		return
	}

	homeUserPath := currentUser.HomeDir
	klangBasePath := filepath.Join(homeUserPath, ".klang")

	paths := []string{
		klangBasePath,
		filepath.Join(klangBasePath, "bin"),
		filepath.Join(klangBasePath, "version"),
		filepath.Join(klangBasePath, "active"),
	}

	for _, path := range paths {
		if err = os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Printf("Permission error while creating '%s': %v\n", path, err)
			return
		}
	}

	kcPath := filepath.Join(klangBasePath, "bin", "kc")
	kcContent := []byte("#!/bin/sh\njava -jar \"$HOME/.klang/active/klang.jar\" \"$@\"")

	if err = makeFile(kcContent, kcPath); err != nil {
		return
	}

	os.Chmod(kcPath, 0755)

	expandedConfigPath := strings.Replace(shellConfigPath, "~", homeUserPath, 1)

	if found, err := fileContains(expandedConfigPath, klangBinSubstr); err == nil && !found {
		if err := appendLine(expandedConfigPath, klangBinPathLine); err != nil {
			fmt.Printf("Warning: Failed to add PATH to shell config file: %v\n", err)
		} else {
			fmt.Println("\nAdded ~/.klang/bin to your PATH. Restart your terminal or run:")
			fmt.Printf("  source %s\n", shellConfigPath)
		}
	}

	klangJarUrl, err := getLatestKlangJarURL()
	if err != nil {
		log.Fatalf("Error determining the latest download URL: %v", err)
		return
	}

	klangJarPath := filepath.Join(paths[3], "klang.jar")

	fmt.Printf("Downloading %s to %s...\n", klangJarUrl, klangJarPath)

	if err := downloadFile(klangJarPath, klangJarUrl); err != nil {
		log.Fatalf("Error downloading the file: %v", err)
		return
	}

	fmt.Println("Download complete!")
	fmt.Println("\n=============================================")
	fmt.Println("Klang installed successfully!")
	fmt.Println("Restart your terminal or run:")
	fmt.Printf("  source %s\n", shellConfigPath)
	fmt.Println("Then verify installation with:")
	fmt.Println("  kc --version")
	fmt.Println("=============================================")
}

type GitHubRelease struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func getLatestKlangJarURL() (string, error) {
	// Muda de /releases/latest para /releases para pegar TODOS os releases
	const apiURL = "https://api.github.com/repos/KlangLang/Klang/releases"

	client := &http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("failed to reach GitHub API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GitHub API returned status: %s", resp.Status)
	}

	var releases []GitHubRelease

	if err := json.NewDecoder(resp.Body).Decode(&releases); err != nil {
		return "", fmt.Errorf("failed to parse GitHub API response: %w", err)
	}

	if len(releases) == 0 {
		return "", fmt.Errorf("no releases found")
	}

	// O primeiro release da lista é o mais recente (incluindo pre-releases)
	release := releases[0]

	// Procura pelo asset klang.jar
	for _, asset := range release.Assets {
		if asset.Name == "klang.jar" {
			fmt.Printf("Found latest version: %s\n", release.TagName)
			return asset.BrowserDownloadURL, nil
		}
	}

	return "", fmt.Errorf("klang.jar not found in release %s", release.TagName)
}

func determineShellConfigPath() (string, error) {
	shellPath := os.Getenv("SHELL")
	if shellPath == "" {
		return "", fmt.Errorf("SHELL variable not defined")
	}

	shellName := filepath.Base(shellPath)
	switch shellName {
	case "bash":
		return "~/.bashrc", nil
	case "zsh":
		return "~/.zshrc", nil
	case "fish":
		return "~/.config/fish/config.fish", nil
	default:
		return "~/.profile", nil
	}
}

func appendLine(path, content string) error {
	line := content + "\n"
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("error opening file '%s': %w", path, err)
	}
	defer f.Close()

	if _, err := f.WriteString(line); err != nil {
		return fmt.Errorf("error writing to '%s': %w", path, err)
	}
	return nil
}

func makeFile(content []byte, path string) error {
	if err := os.WriteFile(path, content, 0644); err != nil {
		fmt.Printf("Permission error while creating '%s': %v\n", path, err)
		return err
	}
	return nil
}

func fileContains(path, substring string) (bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return false, fmt.Errorf("error reading '%s': %w", path, err)
	}
	return strings.Contains(string(data), substring), nil
}

func downloadFile(filepath string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: HTTP status %s", resp.Status)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
