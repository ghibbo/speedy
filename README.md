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



Infine, l'applicazione a riga di comando ci permetterà di creare un'applicazione web con il comando: `speedy new <myproject>`	 



## Start 
