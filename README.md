## Balance

This package primarily provides a function for testing whether an input string has a balanced set of braces ('{' and '}'). The data structure used to perform the balance test is also exported.

The Balance function takes an input string and returns then index of the first unbalanced brace. If the string is balanced, a -1 is returned. Below are a few example inputs and outputs:

+ Valid cases
	- "hello world", -1,  
	- "{}", -1,           
	- "{{{foo();}}}{}", -1
	- "{{}{}}", -1        
	- "valid {} case", -1 

+ Invalid cases
	- "{I", 0            
	- "{{used{to}", 0    
	- "{be}{an", 2       
	- "{{adventurer}", 0 
	- "{like}{you}{{}", 4
	- "}But", 0            
	- "}then}}", 0         
	- "{I}{took}{}an}", 6  
	- "}{arrow}{}to", 0    
	- "{{the}} knee} {}", 4
