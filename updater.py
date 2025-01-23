import os

def main():
    file_path = "files/counter.txt"
    
    try:
        with open(file_path, "r") as file:
            content = file.read().strip()
            counter_value = int(content)
    except (FileNotFoundError, ValueError):
        counter_value = 0
    
    counter_value += 1
    
    with open(file_path, "w") as file:
        file.write(str(counter_value))
    
    print("Updated File Content:")
    print(counter_value)

if __name__ == "__main__":
    main()
