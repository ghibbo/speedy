# Speedy



## Intro

Creazione di un modulo Go riutilizzabile che rende la creazione di un'applicazione Web semplice, veloce e sicura.



## Caratteristiche principali:

- Implementazione di un **Object Relation Mapper** (ORM) indipendente dal database. 

- Creazione di un sistema di autenticazione utente che include: 

  - Un sistema di reimpostazione della password. 
  - Autenticazione basata sulla sessione (per applicazioni basate sul Web).

  - Autenticazione basata su token (per API e sistemi creati con front-end come React e Vue).

- Un sistema di migrazione del database completamente funzionale.

- Un sistema di template completo (utilizzando sia i template Go che i template Jet).

- Un sistema di caching completo che supporta Redis e Badger. 

- Facile gestione delle sessioni, con cookie, database (MySQL e Postgres). 

- Gestione Response HTML, XML, JSON e download di file. 

- Convalida dei Form. 

- Un sistema di mail completo che supporta server SMTP e API di terze parti tra cui MailGun, SparkPost e SendGrid.

- Un'applicazione a riga di comando che consente una facile generazione di e-mail, handlers, modelli di database. 



Infine, l'applicazione a riga di comando ci permetterà di creare un'applicazione web con: `speedy new <myproject>`	 



## Start 

### Creazione del progetto

VStudio --> GO-FRAME

2 SottoCartelle --> speedy, myapp

#### speedy:

`go mod init github.com/ghibbo/speedy`

#### myapp:

`go mod init myapp`

Appendo in go.mod: `replace github.com/ghibbo/speedy => ../speedy`

#### Testing del modulo speedy:

Creo diverse funzioni di Test (speedy.go)

Check delle funzioni in main.go (myapp)

### Accorgimento:

VSCode non avverte immediatamente le modifiche fatte nel modulo **speedy**.

Per mantenere sincronizzata l'applicazione ed il package in **myapp** posso: 

1. `go get -u github.com/ghibbo/speedy`
2. `go mod vendor`
3. Creare un Makefile: `make run`  

Bene!!!

## Sviluppo di Speedy

