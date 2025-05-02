## Learning Plan for Mastering a New Backend Programming Language

### Week 1: Language Fundamentals & Environment Setup

**Objectives:**  
- Set up development environment  
- Understand syntax, data types, control structures  
- Write simple programs and functions  

**Core Concepts:**  
- Installation & IDE setup  
- Variables, data types (strings, integers, floats, booleans)  
- Conditional statements (if/else, switch/case)  
- Loops (for, while)  
- Functions and basic I/O operations  

**Hands-on Exercises:**  
- Print “Hello, World!” and basic arithmetic operations  
- Write functions to compute factorial and Fibonacci numbers  
- Create a simple command-line calculator  
- Implement basic string manipulations (e.g., palindrome checker)  

---

### Week 2: Data Structures & Error Handling

**Objectives:**  
- Learn built-in data structures  
- Understand exception handling  
- Introduce modular programming and code organization  

**Core Concepts:**  
- Arrays, lists, dictionaries/maps, sets  
- Exception handling (try/catch/except)  
- Functions with parameters and return values  
- Modules and packages (import/export)  

**Hands-on Exercises:**  
- Implement a contact list using dictionaries/maps  
- Write a program to parse and validate user input with error handling  
- Create a module for basic math utilities (e.g., gcd, lcm)  
- Practice writing unit tests for functions  

---

### Week 3: Object-Oriented Programming & File I/O

**Objectives:**  
- Learn classes, objects, and OOP principles  
- Work with file input/output  
- Understand simple database or data persistence concepts (optional)  

**Core Concepts:**  
- Classes, objects, constructors, methods  
- Encapsulation, inheritance, polymorphism basics  
- Reading from and writing to files (text and JSON)  
- (Optional) Basics of database connectivity or ORM  

**Hands-on Exercises:**  
- Design a class to represent a Product with attributes and methods  
- Implement reading/writing product data to/from a JSON file  
- Extend product class with inheritance (e.g., PerishableProduct)  
- Write a small CLI to load products and display them  

---

### Week 4: Web Backend Basics & Final Mini-Project Preparation

**Objectives:**  
- Understand basic web backend concepts (APIs, routing)  
- Learn to structure a simple backend app  
- Design and start the final mini-project  

**Core Concepts:**  
- HTTP basics, REST principles  
- Simple routing and request handling  
- JSON serialization/deserialization  
- Structuring code for scalability and maintainability  

**Hands-on Exercises:**  
- Create a simple REST API endpoint that returns product data  
- Implement POST/GET requests with basic validation  
- Use a lightweight web framework or standard library HTTP server  
- Plan the final mini-project architecture and data structures  

---

## Final Mini-Project: Shopping Cart with Multiple Discount Types

### Project Overview:  
Build a backend module/application that simulates a shopping cart system supporting multiple types of discounts. The system should expose REST API endpoints to add/remove products, apply discounts, and calculate the final price.

---

### Functional Specifications:

1. **Products:**  
   - Each product has an ID, name, price, and category.

2. **Shopping Cart:**  
   - Supports adding products with quantities.  
   - Supports removing products or adjusting quantities.

3. **Discount Types:**  
   - **Percentage Discount:** Applies a percentage off on the total cart price.  
   - **Fixed Amount Discount:** Deducts a fixed amount from the total price.  
   - **Buy X Get Y Free:** For a given product, buy X units and get Y units free.  
   - **Category Discount:** Applies a percentage discount only to products in a specified category.

4. **Discount Application Rules:**  
   - Multiple discounts can be applied but must not reduce the total below zero.  
   - Discounts are applied in this order: item-level (Buy X Get Y), category-level, then cart-level (percentage and fixed amount).  
   - The system should handle invalid discount inputs gracefully.

5. **API Endpoints:**  
   - `POST /cart/add` — Add product with quantity  
   - `POST /cart/remove` — Remove product or decrease quantity  
   - `POST /cart/discount` — Apply a discount (type and parameters)  
   - `GET /cart/total` — Return current total and itemized breakdown  

6. **Response Format:**  
   - JSON with details of cart items, discounts applied, and final total.

---

### Technical Requirements:

- Use the new programming language and appropriate web backend libraries or frameworks.  
- Code should be modular, with classes or modules for Product, Cart, Discount, etc.  
- Validate all inputs and return meaningful error messages.  
- Include unit tests covering core logic (discount calculations, cart operations).  
- Document code and API endpoints clearly (e.g., README with usage instructions).  
- (Optional) Persist cart state in-memory or use a lightweight database/file.

---

### Validation Criteria:

- Correct implementation of all discount types and application rules  
- Clean, readable, and maintainable code structure  
- Proper error handling and input validation  
- Functional API endpoints with expected request and response formats  
- Passing unit tests demonstrating core functionality  
- Ability to extend or modify discounts without major refactoring  

---
