# Build the application
go build -o commanderapp ./cmd/main.go

# Create the directory structure for the binary
mkdir -p pkgroot/usr/local/bin
cp commanderapp pkgroot/usr/local/bin/

# Add the LaunchAgent to the directory structure
mkdir -p pkgroot/Library/LaunchAgents
cp com.monomial.commanderapp.plist pkgroot/Library/LaunchAgents/

# Create the .pkg installer
pkgbuild --root ./pkgroot \
         --identifier com.monomial.commanderapp \
         --version 1.0 \
         --install-location / \
         commanderapp-installer.pkg
