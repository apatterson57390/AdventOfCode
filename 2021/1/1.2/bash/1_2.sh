# /bin/bash

PUZZLE_INPUT=()

mapfile -t PUZZLE_INPUT < puzzleInput.txt

INDEX=0
PREV_SUM=0
COUNT=0

for ((i=0; i<=${#PUZZLE_INPUT[@]};i++)); do
    CURRENT_SUM=0
    DEPTH_VALUES=${PUZZLE_INPUT[@]:$i:3}
    
    for DEPTH in $DEPTH_VALUES; do
        let CURRENT_SUM+=$DEPTH
    done

    if (($CURRENT_SUM > $PREV_SUM)); then
        COUNT=$((COUNT + 1))
        PREV_SUM=$CURRENT_SUM
    else
        PREV_SUM=$CURRENT_SUM
    fi
done

# subtract 1 from the result to account for first comparison which doesn't count
let COUNT=$(($COUNT-1))
echo $COUNT