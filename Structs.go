//Email
package main

import "fmt"

type Email struct {
	Direccion, Mensaje string
}

func NewEmail(direccion, mensaje string) Email {
	return Email{
		Direccion: direccion,
		Mensaje:   mensaje,
	}
}

func (e Email) Enviar() string {
	return e.Direccion + e.Mensaje
}

type SMS struct {
	Numero, Mensaje string
}

func NewSMS(numero, mensaje string) SMS {
	return SMS{
		Numero:  numero,
		Mensaje: mensaje,
	}
}

func (s SMS) Enviar() string {

	return s.Numero + s.Mensaje
}

type Notificable interface {
	Enviar() string
}

func EnviarNotificaciones(ns []Notificable) {
	for _, notificable := range ns {
		fmt.Println(notificable.Enviar())
	}
}

func main() {
	email := NewEmail("camy@mail.com", "Hola, soy yo")
	sms := NewSMS("1123395768", "Hola, soy yo")

	notificaciones := []Notificable{email, sms}

	EnviarNotificaciones(notificaciones)
}
