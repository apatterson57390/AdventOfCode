#/bin/bash

PUZZLE_INPUT=$(cat ./puzzleInput.txt)
COUNT=0
for DEPTH in $PUZZLE_INPUT
    do
        if [[ -z $PREV_DEPTH ]]
        then
            echo "hit"
            PREV_DEPTH="$DEPTH"
            continue
        else
            if (("$DEPTH" > "$PREV_DEPTH"))
            then
                echo "$DEPTH > $PREV_DEPTH: $COUNT"
                PREV_DEPTH="$DEPTH"
                COUNT=$((COUNT + 1))
            else
                PREV_DEPTH="$DEPTH"
            fi
        fi
done
echo $COUNT
