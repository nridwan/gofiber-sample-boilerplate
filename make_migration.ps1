# args 0 -> migration name
if (($args[0] -ne $null)) {
    $version = [int][double]::Parse((Get-Date -UFormat %s))
    echo $null > "migration/${version}_$($args[0]).up.sql"
    echo $null > "migration/${version}_$($args[0]).down.sql"
} else {
	exit 1
}