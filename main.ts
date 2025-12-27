/**
 * @author Joshua Adeyemi
 * @version 1.0.0
 * @date 2025-12-26
 * @fileoverview Calculator using predefined functions.
 */

console.log("Welcome to my calculator program.");
console.log("Which operation would you like to perform today?");
console.log("a. add");
console.log("b. subtract");
console.log("c. multiply");
console.log("d. divide");
console.log("e. absolute value");
console.log("f. round");
console.log("g. raise to an exponent");
console.log("h. square root");

const choice: string = prompt("What operation do you want to choose:") || "";

let num1: number;
let num2: number;
let result: number;

if (choice === "a") {
  num1 = Number(prompt("Enter first number:"));
  num2 = Number(prompt("Enter second number:"));
  result = num1 + num2;
  console.log(`${num1} + ${num2} = ${result}`);
} else if (choice === "b") {
  num1 = Number(prompt("Enter first number:"));
  num2 = Number(prompt("Enter second number:"));
  result = num1 - num2;
  console.log(`${num1} - ${num2} = ${result}`);
} else if (choice === "c") {
  num1 = Number(prompt("Enter first number:"));
  num2 = Number(prompt("Enter second number:"));
  result = num1 * num2;
  console.log(`${num1} * ${num2} = ${result}`);
} else if (choice === "d") {
  num1 = Number(prompt("Enter dividend:"));
  num2 = Number(prompt("Enter divisor:"));
  result = num1 / num2;
  console.log(`${num1} / ${num2} = ${result}`);
} else if (choice === "e") {
  num1 = Number(prompt("Enter a number:"));
  result = Math.abs(num1);
  console.log(`The absolute value of ${num1} is ${result}`);
} else if (choice === "f") {
  num1 = Number(prompt("Enter a number:"));
  result = Math.round(num1);
  console.log(`${num1} rounded is ${result}`);
} else if (choice === "g") {
  num1 = Number(prompt("Enter base number:"));
  num2 = Number(prompt("Enter exponent:"));
  result = Math.pow(num1, num2);
  console.log(`${num1} raised to the power of ${num2} is ${result}`);
} else if (choice === "h") {
  num1 = Number(prompt("Enter a non-negative number:"));
  result = Math.sqrt(num1);
  console.log(`The square root of ${num1} is ${result}`);
} else {
  console.log("Invalid choice.");
}

console.log("\nDone.");
