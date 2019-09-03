package main

import (
  "fmt"
)

type helloworld interface {
  init()
  hello() string
  world() string
  hw() string
}

// default implementation of methods
type defaultHelloworld struct {}
func (in defaultHelloworld) init() {}
func (in defaultHelloworld) hello() string { return ""}
func (in defaultHelloworld) world() string { return ""}
// this doesn't work as it uses hello() & world() implemetation of the default class
func (in defaultHelloworld) hw() string { return in.hello() + " " + in.world()}

// ------------------------ german object
type german struct{
  defaultHelloworld
  name string
}

func (in german) hello() string {
  return in.name + ". Gutten"
}

func (in german) world() string {
  return "Tag"
}

// ------------------------ french object
type french struct{
  defaultHelloworld
  name string
}

func (in french) hello() string {
  return in.name + ". Bonjour"
}

func (in french) world() string {
  return "tout les mondes"
}

// ------------------------ russian object
type russian struct{
  defaultHelloworld
  name string
  balala string
}

//pass pointer here, otherwise the function can't set the variable
func (in *russian) init() {
  if in.balala == "" {in.balala = "Балалайка"}
}

func (in russian) hello() string {
  return in.name + ". Привет"
}

func (in russian) world() string {
  return in.balala
}

//every struct that implements same three methods as helloworld inteface get second type = helloworld
func allinone(in helloworld) {
  in.init()
  theline := in.hello() + " " + in.world()
  fmt.Println(theline)
  fmt.Println(in.hw())
}

func main3() {
  r := russian{name: "Русский"}
  f := french{name: "Francious"}
  g := german{name: "Deutsch"}

  //pass pointer here so the functino can really set the variable
  allinone(&r)
  allinone(f)
  allinone(g)

}
