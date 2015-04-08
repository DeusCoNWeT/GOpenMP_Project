/*
 =====================================================================================================
 Name        : import_processor.go
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
               
 Description : Module that handles import declarations from the original code (especially "runtime").
 =====================================================================================================
*/

package import_processor

import (
	"go/token"
	. "goprep"
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

// Function that process an import declaration.
func Imports_declare(tok Token, in chan Token, out chan string, sync chan interface{}) {
	var enc bool = false
	for tok.Token == token.IMPORT {
		passToken(tok, out, sync)
		tok = <-in
		switch tok.Token {
		case token.LPAREN:
			passToken(tok, out, sync)
			tok = <-in
			for tok.Token != token.RPAREN {
				switch tok.Token {
				case token.PERIOD:
					passToken(tok, out, sync)
					tok = <-in
					if tok.Str == "\"runtime\"" {
						enc = true
						passToken(tok, out, sync)
						tok = <-in
						passToken(tok, out, sync)
						tok = <-in
					} else {
						passToken(tok, out, sync)
						tok = <-in
						passToken(tok, out, sync)
						tok = <-in
					}
				case token.IDENT:
					if tok.Str == "_" {
						passToken(tok, out, sync)
						tok = <-in
						if tok.Str == "\"runtime\"" {
							enc = true
							passToken(tok, out, sync)
							tok = <-in
							passToken(tok, out, sync)
							tok = <-in
						} else {
							passToken(tok, out, sync)
							tok = <-in
							passToken(tok, out, sync)
							tok = <-in
						}
					} else {
						panic("Unrecognized token \"" + tok.Str + "\" inside import declaration.")
					}
				case token.STRING:
					if tok.Str == "\"runtime\"" {
						enc = true
						passToken(tok, out, sync)
						tok = <-in
						passToken(tok, out, sync)
						tok = <-in
					} else {
						passToken(tok, out, sync)
						tok = <-in
						passToken(tok, out, sync)
						tok = <-in
					}
				case token.COMMENT: // Ignored comments
					passToken(tok, out, sync)
					tok = <-in
				}
			}
		case token.PERIOD:
			passToken(tok, out, sync)
			tok = <-in
			if tok.Str == "\"runtime\"" {
				enc = true
				passToken(tok, out, sync)
				tok = <-in
				passToken(tok, out, sync)
				tok = <-in
			} else {
				passToken(tok, out, sync)
				tok = <-in
				passToken(tok, out, sync)
				tok = <-in
			}
		case token.IDENT:
			if tok.Str == "_" {
				passToken(tok, out, sync)
				tok = <-in
				if tok.Str == "\"runtime\"" {
					enc = true
					passToken(tok, out, sync)
					tok = <-in
					passToken(tok, out, sync)
					tok = <-in
				} else {
					passToken(tok, out, sync)
					tok = <-in
					passToken(tok, out, sync)
					tok = <-in
				}
			} else {
				panic("Unrecognized token \"" + tok.Str + "\" inside import declaration.")
			}
		case token.STRING:
			if tok.Str == "\"runtime\"" {
				enc = true
				passToken(tok, out, sync)
				tok = <-in
				passToken(tok, out, sync)
				tok = <-in
			} else {
				passToken(tok, out, sync)
				tok = <-in
				passToken(tok, out, sync)
				tok = <-in
			}
		case token.COMMENT: // Ignored comments
			passToken(tok, out, sync)
			tok = <-in
		}
	}
	if !enc { // Include "runtime" package
		passToken(tok, out, sync)
		tok = <-in
		out <- tok.Str + "\n" + "import \"runtime\"\n"
		sync <- nil
	} else {
		passToken(tok, out, sync)
	}
}
