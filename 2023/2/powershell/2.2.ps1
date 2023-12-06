function Get-CubeChallengeCount {
    Param (
        [string]
        [ValidateSet("Red", "Blue", "Green")]
        $Color
    )

    @{
        "Red" = 12
        "Green" = 13
        "Blue" = 14
    }[$Color]
}

function Get-CubeCount {
    Param (
        [string]
        $CubeString
    )

    $builder = [System.Text.StringBuilder]::new()
    $CubeString.ToCharArray() | ForEach-Object {
        $value = ""
        try {
            $value = [int]::Parse($_)
        } catch {

        }
        if ($value -or ($value -eq 0)) {
            $builder.Append($value) | Out-Null
        }
    }
    $cubeCount = $builder.ToString()

    [PSCustomObject]@{
        Color = $CubeString.split($cubeCount)[1]
        Count = [int]::Parse($cubeCount)
    }
}

function Get-IsGamePossible {
    [OutputType([bool])]
    Param(
        [string]
        $Game
    )

    $rounds = $game.split(":")[1].replace(" ","").split(";")
    foreach ($round in $rounds) {
        foreach ($cubeSet in $round.split(",")) {
            $numberOf = (Get-CubeCount -CubeString $cubeSet).count
            if ([int]::Parse($numberOf) -gt (Get-CubeChallengeCount -Color $cubeset.Split($numberOf.ToString())[1])) {
                return $false
            }
        }
    }
    $true
}


function Get-Part1 {
    $total = 0
    (Get-Content ./puzzleInput.txt).forEach({
        $gameId = [int]::Parse($_.split(":")[0].replace("Game ",""))
        $isPossible = Get-IsGamePossible -Game $_

        if ($isPossible) {
            $total = $total + $gameId
        }
    })

    $total
}

function Get-Part2 {
    $total = 0
    (Get-Content ./puzzleInput.txt).forEach({
        $cubeCounts = @{}
        $game = $_
        $game.split(":")[1].replace(" ","").split(";").forEach({
            $round = $_
            $cubes = $round.split(",")
            $cubes.forEach({
                $result = Get-CubeCount -CubeString $_
                $cubeCounts[$result.Color] = $cubeCounts[$result.Color] + @($result.Count)
            })
            
        })
        $last = $null
        foreach($key in $cubeCounts.Keys.GetEnumerator()) {
            $current = ($cubeCounts[$key] | Sort-Object)[-1]
            if ($null -eq $last) {
                $last = $current
                continue
            }
            $last = ($current * $last)
        }
        $total = $last + $total
    })
    $total
}

Get-Part2
