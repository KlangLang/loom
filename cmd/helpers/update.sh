#!/usr/bin/env sh
set -eu

REPO="KlangLang/loom"
VERSION=$(curl -s https://api.github.com/repos/KlangLang/loom/releases/latest \
    | grep tag_name \
    | cut -d '"' -f 4)

if ! [ -e "$HOME/loom" ]; then
    echo "loom not installed!" 
    echo "running install.sh..."

    ./install.sh

    exit 0
fi

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
  Linux)  OS="Linux" ;;
  Darwin) OS="Darwin" ;;
  MINGW*|MSYS*|CYGWIN*) OS="Windows" ;;
  *) OS="$(uname -s)" ;;
esac

case "$ARCH" in
  x86_64|amd64) ARCH="x86_64" ;;
  aarch64|arm64) ARCH="arm64" ;;
  i386|i686) ARCH="i386" ;;
esac

FILE_NAME="loom_${OS}_${ARCH}.tar.gz"
URL="https://github.com/$REPO/releases/download/$VERSION/$FILE_NAME"

echo "Downloading: $URL"
echo "Updating to $VERSION..."

TMPDIR="$(mktemp -d 2>/dev/null || (printf '%s\n' "/tmp/loom_install.$$" && mkdir -p "/tmp/loom_install.$$" && printf "%s\n" "/tmp/loom_install.$$"))"
trap 'rm -rf "$TMPDIR"' EXIT

curl -L --fail "$URL" -o "$TMPDIR/$FILE_NAME"

tar -xzf "$TMPDIR/$FILE_NAME" -C "$TMPDIR"

LOOM_BIN="$(find "$TMPDIR" -type f -name loom -print | head -n 1 || true)"

if [ -z "$LOOM_BIN" ]; then
  echo "❌ loom binary not found inside archive."
  echo "Contents:"
  tar -tzf "$TMPDIR/$FILE_NAME" | sed -n '1,50p'
  exit 1
fi

chmod +x "$LOOM_BIN"

# install target: prefer /usr/local/bin when sudo available
if command -v sudo >/dev/null 2>&1; then
  echo "Installing to /usr/local/bin (requires sudo)..."
  sudo mv "$LOOM_BIN" /usr/local/bin/loom
else
  echo "Installing to ~/.local/bin (no sudo)..."
  mkdir -p "$HOME/.local/bin"
  mv "$LOOM_BIN" "$HOME/.local/bin/loom"
  printf '\nNote: add ~/.local/bin to your PATH if needed (e.g. export PATH="$HOME/.local/bin:$PATH")\n'
fi

OLD_VERSION=$(loom -V | grep -E '^loom ' | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+')
echo "$INSTALLED_VERSION"

echo
echo "✔ Loom updated from $OLD_VERSION to $VERSION"
echo "→ Run: loom --version"