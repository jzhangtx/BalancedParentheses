BalancedParentheses: BalancedParentheses.o
	g++ -g -o BalancedParentheses.exe BalancedParentheses.o -pthread    

BalancedParentheses.o: BalancedParentheses/BalancedParentheses.cpp
	g++ -g  -c -pthread BalancedParentheses/BalancedParentheses.cpp
