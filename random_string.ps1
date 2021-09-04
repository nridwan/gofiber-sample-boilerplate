$FIRST=(-join ((48..57) + (65..90) + (97..122) | Get-Random -Count 32 | % {[char]$_}))
$SECOND=(-join ((48..57) + (65..90) + (97..122) | Get-Random -Count 32 | % {[char]$_}))
echo "${FIRST}${SECOND}"
