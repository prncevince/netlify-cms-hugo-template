# copy nightowl.py to chroma directory
$CHROMA = $GOPATH/src/github.com/alecthomas/chroma
cp ./tools/nightowl.py $CHROMA/_tools
# run style generator
python $CHROMA/_tools/style.py nightowl nightowl.NightOwl > nightowl.go
mv nightowl.go $CHROMA/styles/
# change values in the go script ... or make it work from cli
go run new_chromastyles.go > ./static/nightowl-chroma.css
# also, need a configure option in config file (toml,yaml,json)
