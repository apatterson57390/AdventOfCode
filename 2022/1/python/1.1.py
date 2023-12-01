def read_input() -> list:
    fileContents = ""
    with open("./puzzleInput.txt") as file:
        fileContents = file.read()
    return fileContents.split("\n")

def main():    
    puzzleInput = read_input()

    currentCalorieHigh = 0
    calorieSum = 0
    for calorie in puzzleInput:
        # this is a newline, swtich to the next food pack
        if not calorie:
            if calorieSum > currentCalorieHigh:
                currentCalorieHigh = calorieSum
            calorieSum = 0
            continue
        calorie = int(calorie)
        calorieSum += calorie
    
    print(currentCalorieHigh)

if __name__ == "__main__":
    main()
            

        
        