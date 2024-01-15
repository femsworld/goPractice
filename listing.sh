ls -la 
ls --ignore={.,..}
ls --quoting-style=shell | awk '{for (i=1; i<=NF; i++) printf "%s%s", $i, (i<NF ? ", " : RS)}'
ls -ltu
ls -lF --indicator-style=slash
