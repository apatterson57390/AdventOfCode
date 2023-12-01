function Get-IndicesForWord
{
    Param
    (
        [string]
        $CalibrationValue, 

        [string]
        $Word
    )

    $afterIndex = 0
    $result = @()

    do {
        $index = $CalibrationValue.IndexOf($Word, $afterIndex)
        if ($index -eq -1) {
            break
        }

        $result += [PSCustomObject]@{
            Index = $index;
            Value = $word;
            IsWord = $true
        }
        $afterIndex = $index + 1
    }
    while ($true)

    return $result
}

function Get-CalibrationValue
{
    [cmdletbinding()]
    [OutputType([int])]
    Param
    (
        [Parameter(Mandatory=$true)]
        [String]
        $CalibrationValue
    )
    
    $numstrings = @{
        "one" = 1
        "two" = 2
        "three" = 3
        "four" = 4
        "five" = 5
        "six" = 6
        "seven" = 7
        "eight" = 8
        "nine" = 9
    }
    
    $numvalues = @()
    $idx = 0
    $CalibrationValue.ToCharArray() | ForEach-Object {
        $value = 0
        if ([int]::TryParse($_, [ref]$value))
        {
            $numvalues += [PSCustomObject]@{
                Index = $idx;
                Value = $value;
                IsWord = $false
            }
        }
        $idx = $idx + 1
    }

    $wordvalues = @()
    $numstrings.Keys | ForEach-Object {
        $wordvalues += (Get-IndicesForWord -CalibrationValue $CalibrationValue -Word $_)
    }

    $finalList = $numvalues + $wordvalues
    $finalList = $finallist | Sort-Object Index
    
    $firstValue = 0
    $lastValue = 0

    if ($finalList[0].IsWord) {
        $firstValue = $numstrings[$finalList[0].Value]
    } else {
        $firstValue = $finalList[0].Value
    }

    if ($finallist.Count -lt 2) {
        return [int]::Parse("$firstValue$firstValue") 
    }

    if ($finalList[-1].IsWord) {
        $lastValue = $numstrings[$finalList[-1].Value]
    } else {
        $lastValue = $finalList[-1].Value
    }
    return [int]::Parse("$($firstValue)$($lastValue)")
}   

$finalSum = 0
Get-Content ./puzzleInput.txt | ForEach-Object {
    $finalSum = $finalSum + $(Get-CalibrationValue -CalibrationValue $_)
}

$finalSum
