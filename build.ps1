# Script for building all build types in a finger snaps
# args
if (($args[0] -ne $null) -AND ($args[1] -ne $null)) {
	$Env:GOOS="$($args[0])"
    $Env:GOOARCH="$($args[1])"
    $RELEASE_DIR="release/$($args[0])_$($args[1])"
    if ("$($args[0])" -eq "windows") {
        $RELEASE_EXTENSION=".exe"
    } else {
        $RELEASE_EXTENSION=""
    }
} else {
	$RELEASE_DIR="release/current"
    $RELEASE_EXTENSION=".exe"
}
mkdir release -ErrorAction SilentlyContinue
mkdir "$RELEASE_DIR" -ErrorAction SilentlyContinue
mkdir "$RELEASE_DIR/migration" -ErrorAction SilentlyContinue
cp .env.example "$RELEASE_DIR/.env"
Copy-Item -Path "migration/*.sql" -Destination "$RELEASE_DIR/migration/" -ErrorVariable capturedErrors -ErrorAction SilentlyContinue
go build -o "$RELEASE_DIR/server$RELEASE_EXTENSION" main.go
go build -o "$RELEASE_DIR/migrate$RELEASE_EXTENSION" migration/migrate.go