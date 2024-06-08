print("menu")
print("1. add")
print("2. subtract")
print("3. multiply")
print("4. devide")
choice = input("enter 1-4: ")
if choice == '1':
    num1 = input("first number: ")
    num2 = input("second number: ")
    print('the sum is', int(num1) + int(num2))
elif choice == '2':
    num1 = input("first number: ")
    num2 = input("second number: ")
    print('the differance is', int(num1) - int(num2))
elif choice == '3':
    num1 = input("first number: ")
    num2 = input("second number: ")
    print('the product is', int(num1) * int(num2))
elif choice == '4':
    num1 = input("first number: ")
    num2 = input("second number: ")
    print('the quotient is', float(int(num1)/int(num2)))
else:
    print("goodbye")
    exit