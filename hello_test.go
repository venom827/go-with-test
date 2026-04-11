package main

import "testing"

func TestHello(t *testing.T){
	t.Run("in hindi", func(t *testing.T){
		got:=Hello("Karan","Hindi")
		want:="Namaste Karan\n"
		assertCorrectMessage(t,got,want)
	})
	t.Run("in french", func(t *testing.T){
		got:=Hello("Camille","French")
		want:="Bonjour Camille\n"
		assertCorrectMessage(t,got,want)       
	})
	t.Run("in spanish", func (t *testing.T){
		got:= Hello("Elodie", "Spanish")
		want:= "Hola Elodie\n"
		assertCorrectMessage(t,got,want)
	})
	t.Run("saying hello to people", func (t *testing.T){
		got:=Hello("Chris","")
		want:="Hello Chris\n"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying 'Hello, World\n' when empty string is passed", func (t *testing.T){
		got:=Hello("","")
		want:="Hello World\n"

		assertCorrectMessage(t, got, want)
	})
		
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got!=want{
			t.Errorf("got %q want %q", got, want)
	}
}