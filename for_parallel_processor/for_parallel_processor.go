/*
 =========================================================================================================
 Name        : for_processor.go
 Author      : Enrique Madridejos Zamorano
 Version     :
 Copyright   : Licensed under the Apache License, Version 2.0 (the "License");
   			   you may not use this file except in compliance with the License.
               You may obtain a copy of the License at

               http://www.apache.org/licenses/LICENSE-2.0

               Unless required by applicable law or agreed to in writing, software
               distributed under the License is distributed on an "AS IS" BASIS,
               WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
               See the License for the specific language governing permissions and
               limitations under the License.

 Description : Module that handles loops inside a pragma parallel for.
               Loop valid structure:

               for  init-expr , var relop b , incr-expr

               Where,

               - “init-expr”: initialization of “var” variable (loop variable), by an integer expression.
               - “relop”: valid operators <, <=, >, >=.
               - “b”: integer expression.
               - “incr-expr”: Increase or decrease of “var”, in a integer number,
                  using a standard operator (++, --,  +=, -=), or by the form “var = var + incr”.
 ==========================================================================================================
*/

package for_parallel_processor

import (
	. "github.com/DeusCoNWeT/GOpenMP_Project/goprep"
	. "github.com/DeusCoNWeT/GOpenMP_Project/var_processor"
	"go/token"
)

// Private token work functions.

// Funtion that let a token pass.
func passToken(tok Token, out chan string, sync chan interface{}) {
	out <- tok.Str
	sync <- nil
}

// Funtion that eliminate a token.
func eliminateToken(out chan string, sync chan interface{}) {
	out <- ""
	sync <- nil
}

// Function that determines valid logical operators.
func logic_operator(tok Token) (bool, string) {
	var err bool
	var inc string
	switch tok.Token {
	case token.LSS, token.GTR:
		err = false
		inc = "0"
	case token.LEQ, token.GEQ:
		err = false
		inc = "1"
	default:
		err = true
	}
	return err, inc
}

// Function to get the type of a variable declared as reduction.
// Launch panic if variable not previously declared.
func search_typ(id string, varGlobalList []Variable, varLocalList []Variable) string {
	var typ string = "error"
	for i := range varGlobalList {
		if id == varGlobalList[i].Ident {
			typ = varGlobalList[i].Type
			break
		}
	}
	for i := range varLocalList {
		if id == varLocalList[i].Ident {
			typ = varLocalList[i].Type
			break
		}
	}
	if typ == "error" {
		panic("Variable \"" + id + "\" in clause not previously declared.")
	}
	return typ
}

// Function that removes a variable from a private varible list.
func delete_element(id string, privateList []string) []string {
	for i := range privateList {
		if id == privateList[i] {
			privateList[i] = privateList[len(privateList)-1]
			privateList = privateList[:len(privateList)-1]
			break
		}
	}
	return privateList
}

// Function that adds a variable to a private variable list.
func add_element(id string, privateList []string) []string {
	for i := range privateList {
		if id == privateList[i] {
			break
		} else {
			element := id + " int"
			privateList = append(privateList, element)
		}
	}
	return privateList
}

// Function that process a loop declaration include in a pragma parallel for.
func For_parallel_declare(tok Token, in chan Token, out chan string, sync chan interface{}, varGlobalList []Variable, varLocalList []Variable) (string, string, string, string, Token) {
	var num_iter, ini, fin, inc, steps, var_indice, aux, assign string
	var err bool
	if tok.Token != token.FOR {
		panic("Error: It must start with keyword \"for\".")
	}
	passToken(tok, out, sync)
	tok = <-in
	var_indice = tok.Str
	// Rewrite the loop
	out <- "_i"
	sync <- nil
	tok = <-in
	if tok.Token != token.DEFINE && tok.Token != token.ASSIGN {
		panic("Error: Loop variable must be defined implicitly.")
	} else {
		assign = tok.Str
	}
	out <- ":="
	sync <- nil
	tok = <-in
	if tok.Token != token.INT {
		panic("Error: Variable \"" + tok.Str + "\" must be defined as an integer.")
	}
	ini = tok.Str
	// Rewrite the loop
	out <- "0"
	sync <- nil
	tok = <-in
	if tok.Token != token.SEMICOLON {
		panic("Error: Wait a semicolon.")
	}
	passToken(tok, out, sync)
	tok = <-in
	aux = tok.Str
	if aux != var_indice {
		panic("Error: It must use the same variable in the for declarion.")
	}
	// Rewrite the loop
	out <- "_i"
	sync <- nil
	tok = <-in
	err, inc = logic_operator(tok)
	if err {
		panic("Invalid logical operator.")
	}
	// Rewrite the loop
	out <- "<"
	sync <- nil
	tok = <-in
	if tok.Token == token.INT {
		fin = tok.Str
	} else {
		if tok.Token == token.IDENT {
			typ := search_typ(tok.Str, varGlobalList, varLocalList)
			if typ != "int" {
				panic("Error: Variable \"" + tok.Str + "\" must be defined as an integer.")
			} else {
				fin = tok.Str
			}
		} else {
			panic("Error: Variable \"" + tok.Str + "\" must be defined as an integer.")
		}
	}
	// Rewrite the loop
	out <- "_numCPUs"
	sync <- nil
	tok = <-in
	if tok.Token != token.SEMICOLON {
		panic("Error: Wait a semicolon.")
	}
	passToken(tok, out, sync)
	tok = <-in
	aux = tok.Str
	if aux != var_indice {
		panic("Error: It must use the same variable in the for declarion.")
	}
	out <- "_i"
	sync <- nil
	tok = <-in
	switch tok.Token {
	case token.INC, token.DEC:
		steps = "1"
		// Rewrite the loop
		out <- "++"
		sync <- nil
		tok = <-in
	case token.ADD_ASSIGN, token.SUB_ASSIGN:
		// Rewrite the loop
		out <- "++"
		sync <- nil
		tok = <-in
		if tok.Token != token.INT {
			panic("Error: Variable \"" + tok.Str + "\" must be an integer.")
		}
		steps = tok.Str
		// Rewrite the loop
		eliminateToken(out, sync)
		tok = <-in
	}
	num_iter = "(" + fin + " + " + inc + ") / " + steps // String: "(fin + inc) / steps"
	return num_iter, ini, var_indice, assign, tok
}
