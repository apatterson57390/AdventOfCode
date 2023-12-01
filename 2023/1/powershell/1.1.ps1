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

    $values = @()
    $CalibrationValue.ToCharArray() | ForEach-Object {
        $value = 0
        if ([int]::TryParse($_, [ref]$value))
        {
            $values += $value
        }
    }
    
    if ($values.Count -lt 2) {
        return [int]::Parse("$($values[0])$($values[0])")
    }

    return [int]::Parse("$($values[0])$($values[-1])")
}

$finalSum = 0
Get-Content ./puzzleInput.txt | ForEach-Object {
    $finalSum = $finalSum + $(Get-CalibrationValue -CalibrationValue $_)
}

Write-Host "Sum of all calibration values: $finalSum"