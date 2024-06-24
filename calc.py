import os
def clrscr():
    os.system('cls' if os.name == 'nt' else 'clear')
while True:
    print("menu")
    print("1. add")
    print("2. subtract")
    print("3. multiply")
    print("4. devide")
    choice = input("enter 1-4: ")
    clrscr()
    if choice == '1':
        num1 = input("first number: ")
        num2 = input("second number: ")
        clrscr()
        print('the sum is', int(num1) + int(num2))
    elif choice == '2':
        num1 = input("first number: ")
        num2 = input("second number: ")
        clrscr()
        print('the differance is', int(num1) - int(num2))
    elif choice == '3':
        num1 = input("first number: ")
        num2 = input("second number: ")
        clrscr()
        print('the product is', int(num1) * int(num2))
    elif choice == '4':
        num1 = input("first number: ")
        num2 = input("second number: ")
        clrscr()
        print('the quotient is', float(int(num1)/int(num2)))
    else:
        print("goodbye")
        exit