# Default recipe
default:
    @just --list

# Clean, build, and run report6 with specified groups per page
report6 groups_per_page:
    make
    ./ourport report6 --groups-per-page={{ groups_per_page }}

# Generate reports with 5, 10, 20, 50, and 60 groups per page
generate-all:
    @just report6 5
    @just report6 10
    @just report6 20
    @just report6 50
    @just report6 60

# Move generated reports to Obsidian Vault
move-to-obsidian:
    mv ourport-images-* ~/Documents/Obsidian\ Vault

# Generate all reports and move them to Obsidian Vault
generate-and-move: generate-all move-to-obsidian

# Clean up generated files
clean:
    make clean
    rm -f ourport-images-*.md

format:
    just --unstable --fmt
