package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/fernandomalmeida/fatorial/pkg/fatorial"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	const appID = "com.fernandomalmeida.fatorial"
	app, err := gtk.ApplicationNew(appID, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatalf("Não pode criar a aplicação: %s", err)
	}

	app.Connect("activate", func() {
		NewFatorialWindow(app)
	})

	os.Exit(app.Run(os.Args))
}

func NewFatorialWindow(app *gtk.Application) error {
	builder, err := gtk.BuilderNew()
	if err != nil {
		return fmt.Errorf("Não pode criar a FatorialWindow: %s", err)
	}

	// builder.AddFromFile("./fatorial.glade")
	builder.AddFromString(`<?xml version="1.0" encoding="UTF-8"?>
		<!-- Generated with glade 3.22.1 -->
		<interface>
		  <requires lib="gtk+" version="3.20"/>
		  <object class="GtkWindow" id="fatorial">
			<property name="can_focus">False</property>
			<child>
			  <placeholder/>
			</child>
			<child>
			  <object class="GtkBox">
				<property name="visible">True</property>
				<property name="can_focus">False</property>
				<property name="orientation">vertical</property>
				<property name="spacing">10</property>
				<child>
				  <object class="GtkLabel">
					<property name="visible">True</property>
					<property name="can_focus">False</property>
					<property name="label" translatable="yes">Para calcular o fatorial, insira o número desejado (menor que 20) e clique no botão "Calcular!"</property>
					<property name="wrap">True</property>
					<property name="width_chars">40</property>
				  </object>
				  <packing>
					<property name="expand">False</property>
					<property name="fill">True</property>
					<property name="position">0</property>
				  </packing>
				</child>
				<child>
				  <object class="GtkEntry" id="input">
					<property name="visible">True</property>
					<property name="can_focus">True</property>
					<property name="input_purpose">digits</property>
				  </object>
				  <packing>
					<property name="expand">False</property>
					<property name="fill">True</property>
					<property name="position">1</property>
				  </packing>
				</child>
				<child>
				  <object class="GtkButton">
					<property name="label" translatable="yes">Calcular!</property>
					<property name="visible">True</property>
					<property name="can_focus">True</property>
					<property name="receives_default">True</property>
					<signal name="clicked" handler="btnCalcularClicked" swapped="no"/>
				  </object>
				  <packing>
					<property name="expand">False</property>
					<property name="fill">True</property>
					<property name="position">2</property>
				  </packing>
				</child>
				<child>
				  <object class="GtkTextView" id="output">
					<property name="visible">True</property>
					<property name="can_focus">True</property>
					<property name="editable">False</property>
					<property name="cursor_visible">False</property>
					<property name="accepts_tab">False</property>
				  </object>
				  <packing>
					<property name="expand">True</property>
					<property name="fill">True</property>
					<property name="position">3</property>
				  </packing>
				</child>
			  </object>
			</child>
		  </object>
		</interface>
		`)
	builder.ConnectSignals(map[string]interface{}{
		"btnCalcularClicked": func() {
			inputObj, err := builder.GetObject("input")
			input := inputObj.(*gtk.Entry)
			inputText, err := input.GetText()
			if err != nil {
				panic(err)
			}

			outputObj, err := builder.GetObject("output")
			if err != nil {
				panic(err)
			}

			output := outputObj.(*gtk.TextView)

			buffer, err := output.GetBuffer()
			if err != nil {
				panic(err)
			}

			inputInt, err := strconv.ParseUint(inputText, 10, 64)
			if err != nil {
				buffer.SetText(err.Error())
				return
			}

			n, err := fatorial.Fatorial(inputInt)
			if err != nil {
				buffer.SetText(err.Error())
				return
			}

			buffer.SetText(fmt.Sprintf("%d! = %d", inputInt, n))
		},
	})

	winObj, err := builder.GetObject("fatorial")

	win := winObj.(*gtk.Window)

	app.AddWindow(win)

	win.SetDefaultSize(80, 50)
	win.SetPosition(gtk.WIN_POS_CENTER)
	win.Connect("destroy", gtk.MainQuit)
	win.ShowAll()

	gtk.Main()

	return nil
}
