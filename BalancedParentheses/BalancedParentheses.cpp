// BalancedParentheses.cpp : This file contains the 'main' function. Program execution begins and ends there.
//

#include <iostream>
#include <string>
#include <stack>

bool IsBalancedParentheses(std::string& str)
{
    std::stack<char> s;
    for (auto c : str)
    {
        switch (c)
        {
        case '(':
            s.push(')');
            break;

        case '{':
            s.push('}');
            break;

        case '[':
            s.push(']');
            break;

        case ')':
        case '}':
        case ']':
            if (s.empty() || s.top() != c)
                return false;
            s.pop();
            break;

        default:
            break;
        }
    }

    return true;
}

int main()
{
    while (true)
    {
        std::cout << "Please enter a string (\"q\" to exit): ";
        std::string str;
        std::cin >> str;
        if (str == "q" || str == "Q")
            break;

        std::cout << "String \"" << str << "\" " << (IsBalancedParentheses(str) ? "is" : "is not") << " Parentheses balanced" << std::endl;
    }
}
