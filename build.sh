# Script for building all build types in a finger snap
# args
# $1 - first argument, set JDK path
if [ "$1" "!=" "" ] && [ "$2" "!=" "" ]; then
	export GOOS="$1"
    export GOOARCH="$2"
    RELEASE_DIR="release/$1_$2"
    if [ "$1" "==" "windows" ]; then
        RELEASE_EXTENSION=".exe"
    else
        RELEASE_EXTENSION=""
    fi
else
	RELEASE_DIR="release/current"
    RELEASE_EXTENSION=""
fi
mkdir -p release && mkdir -p "$RELEASE_DIR" && mkdir -p "$RELEASE_DIR/migration"
cp .env.example "$RELEASE_DIR/.env"
for file in migration/*.sql; do cp "$file" "$RELEASE_DIR/$file";done
go build -o "$RELEASE_DIR/server$RELEASE_EXTENSION" main.go
go build -o "$RELEASE_DIR/migrate$RELEASE_EXTENSION" migration/migrate.go